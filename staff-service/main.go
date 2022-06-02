package main

import (
	"fmt"
	configs "staff-service/Configs"
	models "staff-service/Models"
	routers "staff-service/Routers"

	"github.com/jinzhu/gorm"
)

func main() {
	var err error
	configs.DB, err = gorm.Open("mysql", configs.DbURL(configs.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status: ", err)
	}

	defer configs.DB.Close()
	configs.DB.AutoMigrate(
		&models.Staff{},
	)

	r := routers.SetupRouter()

	r.Run(":8000")
}
