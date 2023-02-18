package models

type User struct {
	Id       uint64 `gorm:"primaryKey; not null" json:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
