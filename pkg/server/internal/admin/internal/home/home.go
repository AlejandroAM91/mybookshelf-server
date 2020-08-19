package home

import (
	"net/http"

	"github.com/AlejandroAM91/mybookshelf-server/pkg/server/internal/shared/view"
	"github.com/AlejandroAM91/mybookshelf-server/pkg/shared/dto"
)

const (
	title   = "Administration Home"
	content = "web/template/admin/home.html"
)

type vm struct {
	Books []dto.Book
}

// Handler manages home page calls
func Handler(w http.ResponseWriter, req *http.Request) {
	books := make([]dto.Book, 0)
	books = append(books, dto.Book{Title: "Title 1", Author: "Author 1"})
	view.RenderView(w, title, content, vm{Books: books})
}
