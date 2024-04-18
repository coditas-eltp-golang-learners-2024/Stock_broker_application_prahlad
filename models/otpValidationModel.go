package models

import (
	"time"
)

// ValidateOTPRequest represents the request body for validating OTP
type ValidateOTPRequest struct {
	Email string `json:"email" binding:"required"`
	OTP   int    `json:"otp" binding:"required"`
}

// OTP represents the database entity for OTP values
type OTP struct {
	Email          string    `gorm:"column:email"`
	Value          int       `gorm:"column:otp"`
	CreatedAt      time.Time `gorm:"column:createdAt"` // Include a field to track creation time
	EpochTimeStamp int64     `gorm:"column:epochTimeStamp"`
}

// TableName sets the table name for UserInfo explicitly.
func (OTP) TableName() string {
	return "users" // Specify the actual table name here
}
