package model

import (
	"server-client/src/setup"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name     string `gorm:json:"name"`
	Author   string `json:"author`
	Language string `json:language`
}

func init() {

	setup.Connection()
	db = setup.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBooksById(Id int) *Book {
	var book Book
	db.Where("ID=?", Id).Find(&book)
	return &book
}

func DeleteBookById(Id int) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}