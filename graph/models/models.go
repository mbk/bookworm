package models

import (
	"github.com/jinzhu/gorm"
)

func FetchConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "user:user@/bookworm")
	if err != nil {
		panic(err)
	}
	return db
}

type Author struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	BookID    int    `json:"bookID"`
}

type Book struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Author   []int  `json:"author" gorm:"polymorphic:Owner;"`
}
