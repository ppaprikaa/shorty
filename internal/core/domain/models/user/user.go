package user

const (
	UsernameMinLength, UsernameMaxLength = 5, 24
	PasswordMinLength, PasswordMaxLength = 8, 24
)

type Model struct {
	ID       string
	Username string
	Email    string
	Auth
	Profile
}

type Auth struct {
	Password        string
	IsActivate      bool
	ActivationToken string
}

type Profile struct {
	FirstName string
	LastName  string
	ImageURL  string
}
