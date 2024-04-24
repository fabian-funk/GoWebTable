package main

import (
    "fmt"
    "net/http"
    "log"
)

func main() {
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
    })
  
    log.Println("Server started on: http://localhost:80")
    http.ListenAndServe(":80", nil)
}