package main

import (
  "fmt"
  "net/http"
  "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, World")
}

func main() {
  port := os.Getenv("PORT")
  http.HandleFunc("/", handler)
  http.ListenAndServe(":"+port, nil)
}
