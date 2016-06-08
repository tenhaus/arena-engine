package rest

import (
  "fmt"
  "net/http"
  "encoding/json"
)

type SignupRequest struct {
  Email string
  Handle string
  Password string
}

func init() {
  http.HandleFunc("/signup", Signup)
}

func UseMe() {

}

func Signup(w http.ResponseWriter, r *http.Request) {
  decoder := json.NewDecoder(r.Body)
  var signup SignupRequest

  if err := decoder.Decode(&signup); err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "Error: %s", err.Error())
    return
  }

  fmt.Fprintf(w, "Signup |%v|", signup)
}
