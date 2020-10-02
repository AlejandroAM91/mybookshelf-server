package services

import (
	"github.com/AlejandroAM91/mybookshelf-server/internal/server/core/model"
)

// CreateBook creates a new book
func CreateBook(book model.Book) error {
	return nil
}

// RetrieveAllBooks retreave a list of databases
func RetrieveAllBooks() ([]model.Book, error) {
	return []model.Book{}, nil
}
