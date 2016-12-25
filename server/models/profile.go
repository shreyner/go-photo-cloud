package models

import "github.com/jinzhu/gorm"

type Profile struct {
	gorm.Model
	Name	string
}