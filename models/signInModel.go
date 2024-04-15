package models

// SignInCredentials represents the structure of user sign-in credentials.
type SignInCredentials struct {
	Name     string `gorm:"column:name;primaryKey" json:"userName" validate:"required,alpha" example:"Username"`
	Password string `gorm:"column:password" json:"userPassword" validate:"required,alphanum,min=8" example:"userPassword"`
}

// TableName sets the table name for SignInCredentials explicitly.
func (SignInCredentials) TableName() string {
	return "userSignUpCredentials" // Specify the desired table name here
}
