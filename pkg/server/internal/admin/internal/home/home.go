package home

import (
	"net/http"

	"github.com/AlejandroAM91/mybookshelf-server/pkg/db/dao"
	"github.com/AlejandroAM91/mybookshelf-server/pkg/server/internal/admin/internal/shared"
	"github.com/AlejandroAM91/mybookshelf-server/pkg/server/internal/shared/view"
	"github.com/AlejandroAM91/mybookshelf-server/pkg/shared/dto"
)

const (
	title   = "Administration Home"
	content = "web/template/admin/home.html"
)

type model struct {
	Books []dto.Book
}

// Handler manages home page calls
func Handler(w http.ResponseWriter, r *http.Request) {
	model := model{Books: dao.GetAllBooks()}
	view := view.CreateView(title, content)
	view.AddCSS(shared.AssetsPath + "styles.css")
	view.Render(w, model)
}
