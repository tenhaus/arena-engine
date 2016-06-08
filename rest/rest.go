package rest

import (
  "fmt"
  "net/http"
  "io/ioutil"
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
  var signup SignupRequest

  b, _ := ioutil.ReadAll(r.Body)
  fmt.Println(string(b))
  if err := json.Unmarshal(b, &signup); err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "Error: %s", err.Error())
    return
  }

  fmt.Fprintf(w, "Signup |%v|", signup)
}
