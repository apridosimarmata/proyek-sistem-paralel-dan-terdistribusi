package Models

import (
	"time"

	"github.com/google/uuid"
)

type UserRequest struct {
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" gorm:"unique" binding:"required,email"`
	WhatsappNo string `json:"whatsappNo" gorm:"unique" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

type UserChangePasswordRequest struct {
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}

type UserChangeNameRequest struct {
	Email string `json:"email" binding:"required,email"`
	Name  string `json:"name" binding:"required"`
}

type User struct {
	Uuid     uuid.UUID `gorm:"primary_key;"`
	FcmToken string
	UserRequest
	CreatedAt       time.Time
	VerifiedAt      *time.Time
	CodeRequestedAt *int64
}

type UserResponse struct {
	Uuid       string `json:"uuid" gorm:"primaryKey"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	WhatsappNo string `json:"whatsappNo"`
}

type UserVerificationCodeRequest struct {
	WhatsappNo       string `json:"phone" binding:"required,numeric"`
	VerificationCode string `json:"verificationCode"`
}

type UserVerifyPasswordRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
