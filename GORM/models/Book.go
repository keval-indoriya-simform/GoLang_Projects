package models

import (
	"GORM/connection"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model

	Title      string
	Author     string
	CallNumber int `gorm:"unique_index"`
	PersonID   int
}

var (
	db = connection.GetConnection()
)

func GetAllBooks(Books *[]Book) {
	db.Find(&Books)
}

func GetBookByID(Books *Book, id string) {
	db.Find(&Books, id)
}

func AddBook(book *Book) {
	db.Create(&book)
}

func DeleteBookByID(Books *Book, id string) {
	db.Delete(&Books, id)
}

func UpdateBookByID(Books *Book) {
	db.Select("ID", "Title", "Author").Save(&Books)
}
