package configs

import (
	"github.com/ionnotion/fiber-product-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GormConnect() {
	database, err := gorm.Open(sqlite.Open("Product.db"), &gorm.Config{})

	if err !=nil {
		panic("Connection to Database Failed!!")
	}

	err = database.AutoMigrate(&models.Product{},&models.User{})

	if err != nil {
		return
	}

	DB = database
}
