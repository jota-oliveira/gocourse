package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
			fmt.Println("Error parsing template: ", err)
			return
	}
}

var templateCache = make(map[string]*template.Template)

func RenderTemplateTest(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check if we already have the template in the cache
	// it will return the value and a boolean, we care about the boolean now
	_, inMap := templateCache[t]

	if !inMap {
		log.Println("Creating template cache for", t)
		err = createTemplateCache(t)

		if err != nil {
			log.Println("Error creating template cache", err)
		}
	} else {
		log.Println("Using cached template")
	}

	tmpl = templateCache[t]

	err = tmpl.Execute(w, nil)
	
	if err != nil {
		log.Println("Error creating template cache", err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s.page.tmpl", t),
		"./templates/base.layout.tmpl",
	}

	tmpl, err := template.ParseFiles(templates...)

	if err != nil {
		return err
	}

	templateCache[t] = tmpl

	return nil
}