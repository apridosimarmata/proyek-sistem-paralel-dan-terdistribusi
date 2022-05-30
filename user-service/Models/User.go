package Models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	Configs "user-service/Configs"

	"github.com/google/uuid"
)

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New().String())
	return nil
}

func Create(user *User) (err error) {
	if err = Configs.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetByWhatsappNo(user *User, whatsappNo string) (err error) {
	if err = Configs.DB.Where("whatsapp_no = ?", whatsappNo).Find(user).Error; err != nil {
		return err
	}
	return nil
}

func GetByEmail(user *User, email string) (err error) {
	if err = Configs.DB.Where("email = ?", email).Find(user).Error; err != nil {
		return err
	}
	return nil
}

func Update(user *User) (err error) {
	Configs.DB.Save(user)
	return nil
}
