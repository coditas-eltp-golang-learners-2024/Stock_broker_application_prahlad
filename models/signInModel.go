package models

// SignInCredentials represents the structure of user sign-in credentials.
type SignInCredentials struct {
	Name     string `gorm:"column:name;primaryKey" json:"name" validate:"required,alpha" example:"Username"`
	Password string `gorm:"column:password" json:"password" validate:"required,alphanum,min=8" example:"userPassword"`
	Email    string `gorm:"column:email;primaryKey" json:"email" validate:"required,email" example:"testUser@gmail.com"`
}

// TableName sets the table name for SignInCredentials explicitly.
func (SignInCredentials) TableName() string {
	return "users" // Specify the desired table name here
}
