package models

import (
	"github.com/jinzhu/gorm"
	"github.com/kave08/bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name string `gorm:""json:"name"`
	Author string `gorm:"author"`
	Publication string `gorm:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

