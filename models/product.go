package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Id       uint64 `gorm:"primaryKey; not null" json:"id"`
	ProductName string `json:"productName" binding:"required"`
	ProductAmount uint64 `json:"productAmount" binding:"required"`
	ProductPrice uint64 `json:"productPrice" binding:"required"`
	UserId int64 `json:"userId"`
	User User
}

type ProductForm struct {
	ProductName string `json:"productName" binding:"required"`
	ProductAmount uint64 `json:"productAmount" binding:"required"`
	ProductPrice uint64 `json:"productPrice" binding:"required"`
}
