package models

import (
	"github.com/ionnotion/fiber-product-api/helpers"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       uint64 `gorm:"primaryKey; not null;" json:"id"`
	Username string `gorm:"unique" json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) BeforeSave(*gorm.DB) error {

	hash, err := helpers.HashPassword(u.Password)
	
	if err != nil {
		return err
	}

	u.Password = string(hash)

	return nil
}