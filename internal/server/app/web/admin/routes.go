package admin

import "github.com/gorilla/mux"

// InitRoutes creates and initialise admin routes
func InitRoutes(router *mux.Router) {
	router.Handle("/admin", CreateAdminController())
	router.Handle("/admin/book", CreateBookController())
}
