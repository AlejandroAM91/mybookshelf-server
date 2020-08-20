package home

import (
	"net/http"

	"github.com/AlejandroAM91/mybookshelf-server/pkg/db/dao"
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
func Handler(writer http.ResponseWriter, req *http.Request) {
	model := model{Books: dao.BookGetAll()}
	view := view.CreateView(title, content)
	view.Render(writer, model)
}
