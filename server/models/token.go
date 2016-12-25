package models

import (
	"time"
	"github.com/jinzhu/gorm"
)

type Token struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	Token	string 	`gorm:"unique_index"`
	User	User
	UserID	uint
}

func (token *Token) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())
	return nil
}