package models

type Product struct {
	Id       uint64 `gorm:"primaryKey; not null" json:"id"`
	ProductName string `json:"productName" binding:"required"`
	ProductAmount uint64 `json:"productAmount" binding:"required"`
	ProductPrice uint64 `json:"productPrice" binding:"required"`
}

type ProductForm struct {
	ProductName string `json:"productName" binding:"required"`
	ProductAmount uint64 `json:"productAmount" binding:"required"`
	ProductPrice uint64 `json:"productPrice" binding:"required"`
}
