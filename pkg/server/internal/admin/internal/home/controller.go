package home

import (
	"net/http"

	"github.com/AlejandroAM91/mybookshelf-server/pkg/server/internal/shared/view"
)

const (
	title   = "Administration Home"
	content = "web/template/admin/home.html"
)

type vm struct {
}

// Controller creates home controller
func Controller(w http.ResponseWriter, req *http.Request) {
	view.RenderView(w, title, content, vm{})
}
