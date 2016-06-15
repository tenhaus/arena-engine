package www

import (
  "os"
  "net/http"
  "encoding/json"
  "github.com/tenhaus/arena-engine/signup"
  "github.com/tenhaus/arena-engine/accounts/user"
  "github.com/tenhaus/arena-engine/accounts/bot"
  // "github.com/gorilla/mux"
)

func init() {
  environment := os.Getenv("BOTPIT_ENV")

  if environment != "development" {
    Serve()
  }
}

func Serve() {
  // router := mux.NewRouter().StrictSlash(true)
  http.HandleFunc("/signup", SignupHandler)
  // http.ListenAndServe(":8000", router)
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
    var userAccount user.UserAccount
    var serviceAccount bot.ServiceAccount

    r.ParseForm()
    email := r.FormValue("email")
    password := r.FormValue("password")
    handle := r.FormValue("handle")

    err := signup.Signup(handle, email, password,
      &userAccount, &serviceAccount)

    // Write the error and return
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    js, _ := json.Marshal(serviceAccount.Key)

    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
}
