package main

import (
	"fmt"
	"user-service/Configs"
	"user-service/Models"
	"user-service/Routes"

	"github.com/jinzhu/gorm"
)

func main() {
	var err error
	Configs.DB, err = gorm.Open("mysql", Configs.DbURL(Configs.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status: ", err)
	}

	defer Configs.DB.Close()
	Configs.DB.AutoMigrate(
		&Models.User{},
	)

	r := Routes.SetupRouter()

	r.Run(":5001")
}
