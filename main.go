package main

import (
	"fmt"
	"net/http"
)

// func Home(w http.ResponseWriter, r *http.Request) {

// }

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        n, err := fmt.Fprintf(w, "Hello, World!")
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Printf("Number of bytes written: %d\n", n)
    })

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println(err)
    }
}
