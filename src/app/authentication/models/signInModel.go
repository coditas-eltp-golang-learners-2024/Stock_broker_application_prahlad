package models

type SignInCredentials struct {
	UserName string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}
