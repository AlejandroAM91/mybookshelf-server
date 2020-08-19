package home

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type model struct {
	Title string
}

// Controller creates home controller
func Controller(w http.ResponseWriter, req *http.Request) {
	model := &model{Title: "Administration Home"}
	files, _ := filepath.Glob("web/template/shared/*.html")
	files = append(files, "web/template/admin/home.html")
	view, _ := template.ParseFiles(files...)
	view.Execute(w, model)
}
