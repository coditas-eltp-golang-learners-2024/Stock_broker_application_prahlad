package models

// UserInfo represents the structure of user information.
type UserInfo struct {
	Name          string `json:"userName" validate:"required,alpha" example:"TestUser"`
	Email         string `json:"userEmail" validate:"required,email" example:"testUser@gmail.com"`
	PhoneNumber   string `json:"userPhoneNumber" validate:"required,len=10" example:"7878543610"`
	PanCardNumber string `json:"userPanCardNumber" validate:"required,alphanum,len=10" example:"abgjhi6789"`
	Password      string `json:"userPassword" validate:"required,alphanum,min=8" example:"sample11110"`
}

// TableName sets the table name for UserInfo explicitly.
func (UserInfo) TableName() string {
	return "userSignUpCredentials" // Specify the desired table name here
}

// DatabaseConfig represents the configuration values for database connection.
type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yamls:"username"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}
