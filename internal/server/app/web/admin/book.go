package admin

import (
	"net/http"

	"github.com/AlejandroAM91/mybookshelf-server/internal/server/core/model"
	"github.com/AlejandroAM91/mybookshelf-server/internal/server/core/services"
)

// BookController manages book calls
type BookController struct{}

// CreateBookController creates book call handler
func CreateBookController() BookController {
	return BookController{}
}

func (ctrl BookController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		ctrl.post(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
	}
}

func (ctrl BookController) post(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	title := r.FormValue("title")
	author := r.FormValue("author")
	book := model.Book{Title: title, Author: author}
	services.CreateBook(book)

	http.Redirect(w, r, "/admin", http.StatusFound)
}
