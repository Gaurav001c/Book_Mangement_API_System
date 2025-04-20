package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string
	PassWord string
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
