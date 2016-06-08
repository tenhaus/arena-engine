package rest

import (
  "fmt"
  "net/http"
  "google.golang.org/appengine"
  "google.golang.org/appengine/log"
)

func init() {
  http.HandleFunc("/", Signup)
}

func Signup(w http.ResponseWriter, r *http.Request) {
  context := appengine.NewContext(r)
  log.Infof(context, "This should go to the log")
  fmt.Fprint(w, "Works")
}
