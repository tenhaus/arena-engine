package rest

import (
  "fmt"
  "net/http"
  "encoding/json"
  "google.golang.org/appengine"
  "google.golang.org/appengine/log"
)

type SignupRequest struct {
  Email string
  Handle string
  Password string
}

func init() {
  http.HandleFunc("/signup", Signup)
}

func Signup(w http.ResponseWriter, r *http.Request) {
  context := appengine.NewContext(r)

  decoder := json.NewDecoder(r.Body)
  var signup SignupRequest

  if err := decoder.Decode(&signup); err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprint(w, err.Error())
    return
  }

  log.Infof(context, "This should go to the log")
  fmt.Fprintf(w, "Signup |%v|", signup)
}
