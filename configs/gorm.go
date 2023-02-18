package configs

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GormConnect() {
	database, err := gorm.Open(sqlite.Open("Property.db"), &gorm.Config{})

	if err !=nil {
		panic("Connection to Database Failed!!")
	}

	// err = database.AutoMigrate(&PropertyPromotion{})

	if err != nil {
		return
	}

	DB = database
}
