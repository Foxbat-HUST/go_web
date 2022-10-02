package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Age      uint16
	Email    string
	Password string
	Type     string
}
