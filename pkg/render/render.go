package render

import (
	"net/http"
	"html/template"
	"log"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
			log.Fatal("Error parsing template: ", err)
			return
	}
}