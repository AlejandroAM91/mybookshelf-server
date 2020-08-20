package view

import (
	"html/template"
	"net/http"
	"path/filepath"
)

const (
	layoutGlob = "web/template/shared/*.html"
)

// View page view data
type View struct {
	layoutData layoutData
	template   *template.Template
}

type layoutData struct {
	Title string
}

type vm struct {
	Layout  layoutData
	Content interface{}
}

// CreateView creates a page view
func CreateView(title string, content string) View {
	files, _ := filepath.Glob(layoutGlob)
	files = append(files, content)
	tmpl, _ := template.ParseFiles(files...)

	return View{
		layoutData: layoutData{
			Title: title,
		},
		template: tmpl,
	}
}

// Render renderizes a page view
func (view View) Render(writer http.ResponseWriter, data interface{}) {
	vm := vm{Layout: view.layoutData, Content: data}
	view.template.Execute(writer, vm)
}
