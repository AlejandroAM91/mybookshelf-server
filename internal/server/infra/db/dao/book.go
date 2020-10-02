package dao

import (
	"github.com/AlejandroAM91/mybookshelf-server/internal/server/core/model"
	"github.com/AlejandroAM91/mybookshelf-server/internal/server/infra/db"
)

// CreateBook adds new book into the database
func CreateBook(book model.Book) error {
	db := db.Open()
	defer db.Close()

	title := book.Title
	author := book.Author
	if _, err := db.Exec("insert into book (title, author) values ($1, $2)", title, author); err != nil {
		return err
	}
	return nil
}

// GetAllBooks retrieve all books from database
func RetrieveAllBooks() ([]model.Book, error) {
	db := db.Open()
	defer db.Close()

	rows, _ := db.Query("select title, author from book")
	bookList := make([]model.Book, 0)
	for rows.Next() {
		var title, author string
		if err := rows.Scan(&title, &author); err != nil {
			return nil, err
		}
		bookList = append(bookList, model.Book{Title: title, Author: author})
	}

	return bookList, nil
}
