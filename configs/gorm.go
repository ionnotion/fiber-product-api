package configs

import (
	"github.com/ionnotion/fiber-product-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func GormConnect() error {
	DB, err = gorm.Open(sqlite.Open("Product.db?_foreign_keys=on"), &gorm.Config{})

	if err !=nil {
		return err
	}

	err = DB.AutoMigrate(&models.Product{},&models.User{})

	if err != nil {
		return err
	}

	return nil
}
