package dao

import (
	"fmt"

	"github.com/AlejandroAM91/mybookshelf-server/pkg/db"
	"github.com/AlejandroAM91/mybookshelf-server/pkg/shared/dto"
)

// BookCreate adds new book into the database
func BookCreate(book dto.Book) {
	db := db.Open()
	defer db.Close()

	title := book.Title
	author := book.Author
	if _, err := db.Exec("insert into book (title, author) values ($1, $2)", title, author); err != nil {
		fmt.Println(err)
	}
}

// BookGetAll retrieve all books from database
func BookGetAll() []dto.Book {
	db := db.Open()
	defer db.Close()
	rows, _ := db.Query("select title, author from book")
	bookList := make([]dto.Book, 0)
	for rows.Next() {
		var title, author string
		if err := rows.Scan(&title, &author); err != nil {
			fmt.Println(err)
		}
		bookList = append(bookList, dto.Book{Title: title, Author: author})
	}

	return bookList
}
