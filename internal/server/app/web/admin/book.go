package admin

import "net/http"

// BookController manages book calls
type BookController struct{}

// CreateBookController creates book call handler
func CreateBookController() BookController {
	return BookController{}
}

func (ctrl BookController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// case http.MethodGet:
	// 	ctrl.get(w, r)
	// case http.MethodPost:
	// 	ctrl.post(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
	}
}

// Handler manages add book page calls
// func Handler(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case "POST":
// 		post(w, r)
// 	default:
// 		w.WriteHeader(http.StatusNotImplemented)
// 		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
// 	}
// }

// func post(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()

// 	title := r.FormValue("title")
// 	author := r.FormValue("author")
// 	book := dto.Book{Title: title, Author: author}
// 	dao.CreateBook(book)

// 	http.Redirect(w, r, "/admin", http.StatusFound)
// }
