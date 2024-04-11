package models

type UserInfo struct {
	Name          string `json:"userName" validate:"required,alpha"`
	Email         string `json:"userEmail" validate:"required,email"`
	PhoneNumber   uint   `json:"userPhoneNumber" validate:"required,gte=0000000000,lte=9999999999"`
	PanCardNumber string `json:"userPanCardNumber" validate:"required,alphanum,len=10"`
	Password      string `json:"userPassword" validate:"required,alphanum,min=8"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

/*
Note: With these JSON alias tags specified, the JSON marshaling and unmarshaling process will use these tags to map the struct
fields to the corresponding JSON keys.
For example, when marshaling a UserInfo struct to JSON, the field ID will be represented as "user_id" in the JSON output.
*/
