package models

import (
	"github.com/dixi225/go-fiber-postgres/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBook() []Book {
	books := []Book{}
	db.Find(&books)
	return books
}

func GetBookById(Id int) *Book {
	var book Book
	db.Where("ID=?", Id).Find(&book)
	return &book
}

func DeleteBookById(Id int) *Book {
	var book Book
	db.Where("ID=?", Id).Delete(&book)
	return &book
}
