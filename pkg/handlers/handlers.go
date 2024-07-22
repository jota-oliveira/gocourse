package handlers

import (
	"net/http"

	"github.com/jota-oliveira/gocourse/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplateTest(w, "home")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplateTest(w, "about")
}
