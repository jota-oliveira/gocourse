package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jota-oliveira/gocourse/pkg/handlers"
)

var portNumber = ":8081"

func main() {
    http.HandleFunc("/", handlers.Home)
    http.HandleFunc("/about", handlers.About)

    err := http.ListenAndServe(portNumber, nil)

    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }

    fmt.Println("Server is running on port", portNumber)
}
