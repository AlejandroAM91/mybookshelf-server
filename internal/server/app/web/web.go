package web

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/AlejandroAM91/mybookshelf-server/internal/server/app/web/admin"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// The Web manages application web server.
type Web struct {
	serverClosed chan struct{}
	handler      http.Handler
	server       *http.Server
}

// NewWeb creates and initializes aplication web server.
func NewWeb() *Web {
	// Initialize handler
	router := initRoutes()
	handler := initHandler(router)

	return &Web{
		serverClosed: make(chan struct{}),
		handler:      handler,
	}
}

// // Start starts http server
func (web *Web) Start(endpoint string) {
	// Stops previous server
	web.Stop()

	// Create http server
	web.server = &http.Server{
		Addr:    endpoint,
		Handler: web.handler,
	}
	go func() {
		if err := web.server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
		close(web.serverClosed)
	}()
}

// Stop stops http server
func (web Web) Stop() {
	if web.server != nil {
		web.server.Shutdown(context.Background())
	}
}

// Sync wait until server is stopped
func (web Web) Sync() {
	<-web.serverClosed
}

// SyncStart starts http server syncronous
func (web Web) SyncStart(endpoint string) {
	web.Start(endpoint)
	web.Sync()
}

func initHandler(router *mux.Router) http.Handler {
	return handlers.LoggingHandler(os.Stdout, router)
}

func initRoutes() *mux.Router {
	router := mux.NewRouter()
	admin.InitRoutes(router)
	return router
}
