package models

type SignInCredentials struct {
	Name     string `json:"userName" validate:"required,alpha"`
	Password string `json:"userPassword" validate:"required,alphanum,min=8"`
}
