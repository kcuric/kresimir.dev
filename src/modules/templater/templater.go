package templater

import (
	"html/template"
	"net/http"
)

var templates *template.Template

func ParseTemplates() {
	templates = template.Must(template.ParseGlob("./templates/*.html"))
	templates = template.Must(templates.ParseGlob("./templates/components/*.html"))
}

func GenerateTemplate(w http.ResponseWriter, name string, data any) error {
	w.Header().Set("Content-Type", "text/html")
	return templates.ExecuteTemplate(w, name, data)
}
