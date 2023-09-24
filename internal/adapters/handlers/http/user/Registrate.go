package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ppaprikaa/golibs/httpkit/json"
	"github.com/ppaprikaa/golibs/validator"
	"github.com/ppaprikaa/shorty/internal/core/domain/models/user"
	userservice "github.com/ppaprikaa/shorty/internal/services/user"
)

type registrationRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type registrationResponse struct {
	Message string `json:"message"`
}

func (h *Handler) Registrate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := h.readRegistrationRequestAndHandleError(w, r)
		if err != nil {
			return
		}

		err = h.validateRegistrationRequestAndHandleError(w, req)
		if err != nil {
			return
		}

		err = h.registerAndHandleError(w, req)
		if err != nil {
			return
		}

		err = h.sendRegistrationResponseAndHandleError(w)
		if err != nil {
			return
		}
	}
}

func (h *Handler) readRegistrationRequestAndHandleError(w http.ResponseWriter, r *http.Request) (*registrationRequest, error) {
	var req registrationRequest

	if err := json.Read(r, &req); err != nil {
		json.InternalServerError(w)
		return nil, err
	}

	return &req, nil
}

func (h *Handler) validateRegistrationRequestAndHandleError(w http.ResponseWriter, r *registrationRequest) error {
	v := validator.New()

	usernameLength := len([]rune(r.Username))
	passwordLength := len([]rune(r.Password))

	v.Check(
		user.UsernameMinLength <= usernameLength && user.UsernameMaxLength >= usernameLength,
		"username",
		fmt.Sprintf("username's length must be between %d and %d", user.UsernameMinLength, user.UsernameMaxLength),
	)
	v.Check(
		user.PasswordMinLength <= passwordLength && user.PasswordMaxLength >= passwordLength,
		"password",
		fmt.Sprintf("password's length must be between %d and %d", user.PasswordMinLength, user.PasswordMaxLength),
	)
	v.Check(validator.MatchesRX(r.Email, validator.EmailRX), "email", "not an email")

	if !v.Valid() {
		_ = json.UnprocessableEntity(w, v.Errors)
		return errors.New("invalid request")
	}

	return nil
}

func (h *Handler) registerAndHandleError(w http.ResponseWriter, r *registrationRequest) error {
	registerCtx, registerCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer registerCancel()
	if err := h.userService.Registrate(registerCtx, userservice.RegistrationDTO{
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
	}); err != nil {
		switch {
		case errors.Is(err, userservice.ErrTakenUsernameAndEmail):
			json.Conflict(w, registrationResponse{Message: fmt.Sprintf("username: %s and email: %s are already taken.", r.Username, r.Email)})
		case errors.Is(err, userservice.ErrTakenUsername):
			json.Conflict(w, registrationResponse{Message: fmt.Sprintf("username: %s is already taken.", r.Username)})
		case errors.Is(err, userservice.ErrTakenEmail):
			json.Conflict(w, registrationResponse{Message: fmt.Sprintf("email: %s is already taken.", r.Email)})
		case registerCtx.Err() != nil:
			json.GatewayTimeout(w, registrationResponse{Message: "Processing your request took too much time, please retry."})
		default:
			json.InternalServerError(w)
		}

		return err
	}

	return nil
}

func (h *Handler) sendRegistrationResponseAndHandleError(w http.ResponseWriter) error {
	if err := json.OK(w, registrationResponse{Message: "the user has been successfully registered"}); err != nil {
		json.InternalServerError(w)
		return err
	}

	return nil
}
