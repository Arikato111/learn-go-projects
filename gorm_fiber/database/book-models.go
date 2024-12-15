package database

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Price       int    `json:"price"`
}

func (db *Db) CreateBook(book *Book) {
	result := db.Query.Create(book)
	if result.Error != nil {
		log.Fatalf("Error create book: %v", result.Error)
	}

	fmt.Println("create book successfull")
}

func (db *Db) GetBook(id uint) (Book, error) {
	var book Book
	result := db.Query.First(&book, id)
	if result.Error != nil {
		log.Fatalf("Error get book: %v", result.Error)
		return Book{}, result.Error
	}
	return book, nil
}

func (db *Db) GetManyBook() []Book {
	var books []Book
	result := db.Query.Find(&books)
	if result.Error != nil {
		log.Fatalf("Error get book: %v", result.Error)
	}
	return books
}

func (db *Db) Updatebook(book *Book) {
	result := db.Query.Model(&book).Updates(book)
	if result.Error != nil {
		log.Fatalf("Error cannot update book: %v", result.Error)
	}

	fmt.Println("Update book success")
}

func (db *Db) DeleteBook(id uint) {
	var book Book
	result := db.Query.Delete(&book, id)
	if result.Error != nil {
		log.Fatalf("Error cannot delete %d: %v", id, result.Error)
	}

	fmt.Println("Delete success")
}

func (db *Db) SearchBook(bookName string) (Book, error) {
	var book Book
	result := db.Query.Where("name = ?", bookName).First(&book)
	if result.Error != nil {
		return Book{}, result.Error
	}
	return book, nil
}

func (db *Db) SearchManyBook(bookName string) ([]Book, error) {
	var books []Book
	result := db.Query.Where("name LIKE ?", bookName).Order("price").Find(&books)
	if result.Error != nil {
		return []Book{}, result.Error
	}
	return books, nil
}
