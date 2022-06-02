package models

import (
	configs "staff-service/Configs"

	"github.com/google/uuid"
)

type Staff struct {
	Uid      uuid.UUID `json:"uid"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Phone    string    `json:"phone"`
	Password string    `json:"password"`
}

func GetStaffByUID(staff *Staff, staffUID uuid.UUID) error {
	if err := configs.DB.Where("uid = ?", staffUID).First(staff).Error; err != nil {
		return err
	}
	return nil
}

func GetStaffByUsername(staff *Staff, username string) error {
	if err := configs.DB.Where("username = ?", username).First(staff).Error; err != nil {
		return err
	}
	return nil
}
