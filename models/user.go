package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Login	string `gorm:"unique"`
	Password string
	Email	string	`gorm:"unique"`
	Profile	Profile
	ProfileID uint
}