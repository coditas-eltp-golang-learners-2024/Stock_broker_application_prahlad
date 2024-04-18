package models

// UserInfo represents the structure of user information.
type UserInfo struct {
	ID            uint   `gorm:"primaryKey"`
	Name          string `json:"name" validate:"required,alpha" example:"TestUser" gorm:"column:name"`
	Email         string `json:"email" validate:"required,email" example:"testUser@gmail.com" gorm:"column:email;unique"`
	PhoneNumber   string `json:"phoneNumber" validate:"required,len=10" example:"7878543610" gorm:"column:phoneNumber"`
	PanCardNumber string `json:"pancardNumber" validate:"required,alphanum,len=10" example:"abgjhi6789" gorm:"column:pancard"`
	Password      string `json:"password" validate:"required,alphanum,min=8" example:"sample11110" gorm:"column:password"`
}

// TableName sets the table name for UserInfo explicitly.
func (UserInfo) TableName() string {
	return "users" // Specify the desired table name here
}
