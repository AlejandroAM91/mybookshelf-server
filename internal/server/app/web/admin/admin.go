package admin

import (
	"net/http"

	"github.com/AlejandroAM91/mybookshelf-server/internal/server/app/web/shared/view"
	"github.com/AlejandroAM91/mybookshelf-server/internal/server/core/model"
	"github.com/AlejandroAM91/mybookshelf-server/internal/server/core/services"
)

// AdminController manages admin calls
type AdminController struct{}

const (
	title   = "Administration Home"
	content = "web/template/admin/home.html"
)

// CreateAdminController creates admin calls handler
func CreateAdminController() AdminController {
	return AdminController{}
}

func (ctrl AdminController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ctrl.get(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
	}
}

type adminVM struct {
	Books []model.Book
}

// Handler manages home page calls
func (ctrl AdminController) get(w http.ResponseWriter, r *http.Request) {
	bookList, err := services.RetrieveAllBooks()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	model := adminVM{Books: bookList}
	view := view.CreateView(title, content)
	view.AddCSS(assetsPath + "styles.css")
	view.Render(w, model)
}
