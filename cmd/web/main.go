package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jota-oliveira/gocourse/pkg/config"
	"github.com/jota-oliveira/gocourse/pkg/handlers"
	"github.com/jota-oliveira/gocourse/pkg/render"
)

var portNumber = ":8081"

func main() {
    var app config.AppConfig

    tc, err := render.CreateTemplateCache()

    if err != nil {
        log.Fatal("Cannot create template cache")
    }

    app.TemplateCache = tc
    app.UseCache = false

    repo := handlers.NewRepo(&app)
    handlers.NewHandlers(repo)

    render.NewTemplates(&app)

    serve := &http.Server {
        Addr: portNumber,
        Handler: routes(&app),
    }

    err = serve.ListenAndServe()
    log.Fatal(err)

    fmt.Println("Server is running on port", portNumber)
}
