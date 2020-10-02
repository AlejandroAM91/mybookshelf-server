package services

import (
	"github.com/AlejandroAM91/mybookshelf-server/internal/server/core/model"
	"github.com/AlejandroAM91/mybookshelf-server/internal/server/infra/db/dao"
)

// CreateBook creates a new book
func CreateBook(book model.Book) error {
	return dao.CreateBook(book)
}

// RetrieveAllBooks retreave a list of databases
func RetrieveAllBooks() ([]model.Book, error) {
	bookList, err := dao.RetrieveAllBooks()
	return bookList, err
}
