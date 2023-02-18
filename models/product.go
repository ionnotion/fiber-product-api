package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Id       uint64 `gorm:"primaryKey; not null" json:"id"`
	ProductName string `json:"productName" binding:"required"`
	ProductAmount uint64 `json:"productAmount" binding:"required"`
	ProductPrice uint64 `json:"productPrice" binding:"required"`
	OrderedAt string `json:"OrderedAt"`
	User User `gorm:"ForeignKey:Id"`
}

type ProductForm struct {
	ProductName string `json:"productName" binding:"required"`
	ProductAmount uint64 `json:"productAmount" binding:"required"`
	ProductPrice uint64 `json:"productPrice" binding:"required"`
}
