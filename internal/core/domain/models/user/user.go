package user

type Model struct {
	ID string
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
