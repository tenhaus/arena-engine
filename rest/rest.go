package rest

import (
  "fmt"
  "http"
)

func init() {
  http.HandleFunc("/", signup)
}

func signup(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Works")
}
