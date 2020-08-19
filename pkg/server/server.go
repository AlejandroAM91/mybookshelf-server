package server

import (
	"net/http"

	"github.com/AlejandroAM91/mybookshelf-server/pkg/server/internal/admin"
)

// Start initializes and starts application server
func Start(endpoint string) {

	// Initialize routes
	admin.Init()

	// Start server
	http.ListenAndServe(endpoint, nil)
}
