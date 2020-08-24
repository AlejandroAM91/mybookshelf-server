package book

import (
	"net/http"

	"github.com/AlejandroAM91/mybookshelf-server/pkg/db/dao"
	"github.com/AlejandroAM91/mybookshelf-server/pkg/shared/dto"
)

// Handler manages add book page calls
func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		post(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}
}

func post(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	title := r.FormValue("title")
	author := r.FormValue("author")
	book := dto.Book{Title: title, Author: author}
	dao.CreateBook(book)

	http.Redirect(w, r, "/admin", http.StatusFound)
}
