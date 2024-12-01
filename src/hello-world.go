package main

import (
    "fmt"
    "net/http"
)

func main() {
    message := "MESSAGE - cicd - 17"

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, message)
    })

    fmt.Println("Starting server on port 8080.")
    http.ListenAndServe(":8080", nil)
}
