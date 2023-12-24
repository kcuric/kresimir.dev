package templater

import (
	"fmt"
	"html/template"
	"net/http"
)

func GenerateTemplate(w http.ResponseWriter, name string, data any) error {
	w.Header().Set("Content-Type", "text/html")
	locatedTemplate := template.Must(template.ParseFiles(fmt.Sprintf("%s%s", "./templates", name)))
	return locatedTemplate.Execute(w, data)
}
