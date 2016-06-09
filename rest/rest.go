package rest

import (
  "fmt"
  "net/http"
  "github.com/tenhaus/botpit/signup"
)

func init() {
  http.HandleFunc("/signup", Signup)
}

func UseMe() {
}

func Signup(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  email := r.Form["email"]
  password := r.Form["password"]
  handle := r.Form["handle"]

  err := signup.Signup(handle, email, password)
}
