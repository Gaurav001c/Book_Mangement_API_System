package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name  string `json:"name" gorm:"unique"`
	Books []Book `json:"books" gorm:"foreignKey:CategoryID"`
}
