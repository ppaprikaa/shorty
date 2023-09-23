package server

import (
	"net/http"

	"github.com/ppaprikaa/golibs/httpkit/json"
)

func (h *Handler) Healthcheck() http.HandlerFunc {
	type Res struct {
		Status string `json:"status"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if err := json.OK(w, Res{Status: "OK"}); err != nil {
			json.InternalServerError(w)
			return
		}
	}
}
