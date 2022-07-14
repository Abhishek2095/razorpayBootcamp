package main

import (
	"ecommerce/Config"
	"ecommerce/Customer"
	"ecommerce/Order"
	"ecommerce/Product"
	"ecommerce/Routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		panic(err)
	}

	defer func() {
		err := Config.DB.Close()
		if err != nil {
			panic(err)
		}
	}()

	Config.DB.AutoMigrate(&Customer.Customer{}, &Order.Order{}, &Product.Product{})
	r := Routes.SetupRouter()

	//running
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
