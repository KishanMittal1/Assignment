package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "OK")
    })
    fmt.Println("Listening on :8000")
    http.ListenAndServe(":8000", nil)
}