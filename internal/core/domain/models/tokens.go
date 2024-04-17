package models

type TokenPair struct {
	Refresh RefreshToken
	Access  AccessToken
}

type RefreshToken struct {
	Value  string
	Claims RefreshTokenClaims
}

type AccessToken struct {
	Value  string
	Claims AccessTokenClaims
}

type RefreshTokenClaims struct {
	UserID   string
	Username string
}

type AccessTokenClaims struct {
	UserID    string
	Username  string
	SessionID string
}

type Token struct {
	ID        string
	Value     string
	SessionID string
}
