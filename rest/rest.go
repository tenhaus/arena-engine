package rest

import (
  "net/http"
  "github.com/tenhaus/botpit/signup"
  "github.com/tenhaus/botpit/accounts/user"
  "github.com/tenhaus/botpit/accounts/bot"
)

func init() {
  http.HandleFunc("/signup", Signup)
}

func UseMe() {
}

func Signup(w http.ResponseWriter, r *http.Request) {
  var userAccount user.UserAccount
  var serviceAccount bot.ServiceAccount

  r.ParseForm()
  email := r.FormValue("email")
  password := r.FormValue("password")
  handle := r.FormValue("handle")

  err := signup.Signup(handle, email, password,
    &userAccount, &serviceAccount)

  if err != nil {
    
  }
}
