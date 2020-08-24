package admin

import (
	"net/http"

	"github.com/AlejandroAM91/mybookshelf-server/pkg/server/internal/admin/internal/book"
	"github.com/AlejandroAM91/mybookshelf-server/pkg/server/internal/admin/internal/home"
	"github.com/AlejandroAM91/mybookshelf-server/pkg/server/internal/admin/internal/shared"
)

const (
	staticPath string = "web/static/admin"
)

// Init initialize admin routes
func Init() {
	assets := http.FileServer(http.Dir(staticPath))
	http.Handle(shared.AssetsPath, http.StripPrefix(shared.AssetsPath, assets))
	http.HandleFunc("/admin", home.Handler)
	http.HandleFunc("/admin/book", book.Handler)
}
