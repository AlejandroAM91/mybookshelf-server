package admin

import (
	"net/http"

	"github.com/AlejandroAM91/mybookshelf-server/pkg/server/internal/admin/internal/home"
)

// Init initialize admin routes
func Init() {
	http.HandleFunc("/admin", home.Controller)
}
