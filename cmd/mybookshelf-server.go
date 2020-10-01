package main

import (
	"github.com/AlejandroAM91/mybookshelf-server/internal/server/app/web"
	// "github.com/AlejandroAM91/mybookshelf-server/pkg/server"
	// _ "github.com/lib/pq"
)

func main() {
	// server.Start(":80")
	web := web.NewWeb()
	web.SyncStart(":8080")
}
