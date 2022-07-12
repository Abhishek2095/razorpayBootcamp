package main

import (
	"fmt"
	"q2/Config"
	"q2/Models"
	"q2/Routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}

	defer func() {
		if err = Config.DB.Close(); err != nil {
			fmt.Println("Error while Trying to close DB : ", err)
		}
	}()

	Config.DB.AutoMigrate(&Models.Student{})
	Config.DB.AutoMigrate(&Models.Marks{})
	r := Routes.SetupRouter()
	//running
	if err = r.Run(); err != nil {
		fmt.Println("Error while Trying to start server : ", err)
	}
}
