package models

// UserInfo represents the structure of user information.
type UserInfo struct {
	Name          string `json:"userName" validate:"required,alpha"`
	Email         string `json:"userEmail" validate:"required,email"`
	PhoneNumber   string `json:"userPhoneNumber" validate:"required,gte=0000000000,lte=9999999999"`
	PanCardNumber string `json:"userPanCardNumber" validate:"required,alphanum,len=10"`
	Password      string `json:"userPassword" validate:"required,alphanum,min=8"`
}

// TableName sets the table name for UserInfo explicitly.
func (UserInfo) TableName() string {
	return "userSignUpCredentials" // Specify the desired table name here
}

// DatabaseConfig represents the configuration values for database connection.
type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}
