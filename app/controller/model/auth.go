package model

type Auth struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}

type JwtToken struct {
	Token string `json:"token"`
	Type  string `json:"type"`
}
