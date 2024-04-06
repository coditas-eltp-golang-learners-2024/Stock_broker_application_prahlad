package models

type UserInfo struct {
	ID            int32  `json:"user_id"`
	Name          string `json:"user_name"`
	Email         string `json:"user_email"`
	PhoneNumber   string `json:"user_phoneNumber"` // Changed type to string
	PanCardNumber string `json:"user_panCardNumber"`
	Password      string `json:"user_password"`
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
