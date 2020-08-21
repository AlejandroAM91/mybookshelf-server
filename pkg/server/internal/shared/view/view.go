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
	assets     assets
	layoutData layoutData
	template   *template.Template
}

type assets struct {
	CSS []string
}

type layoutData struct {
	Title string
}

type vm struct {
	Assets  assets
	Layout  layoutData
	Content interface{}
}

// CreateView creates a page view
func CreateView(title string, content string) View {
	// Loads assets
	cssFiles := make([]string, 0)

	// Loads templates
	tmplFiles, _ := filepath.Glob(layoutGlob)
	tmplFiles = append(tmplFiles, content)
	tmpl, _ := template.ParseFiles(tmplFiles...)

	return View{
		assets: assets{
			CSS: cssFiles,
		},
		layoutData: layoutData{
			Title: title,
		},
		template: tmpl,
	}
}

// AddCSS adds CSS file to the view
func (view *View) AddCSS(path string) {
	view.assets.CSS = append(view.assets.CSS, path)
}

// Render renderizes the view
func (view View) Render(writer http.ResponseWriter, data interface{}) {
	vm := vm{Assets: view.assets, Layout: view.layoutData, Content: data}
	view.template.Execute(writer, vm)
}
