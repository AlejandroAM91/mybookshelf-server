package view

import (
	"net/http"
	"path/filepath"
	"text/template"
)

const (
	layoutGlob = "web/template/shared/*.html"
)

type view struct {
	Title string
	Data  interface{}
}

// RenderView renderize a page view
func RenderView(w http.ResponseWriter, title string, content string, data interface{}) {
	vm := view{
		Title: title,
		Data:  data,
	}

	files, _ := filepath.Glob(layoutGlob)
	files = append(files, content)
	view, _ := template.ParseFiles(files...)
	view.Execute(w, vm)
}
