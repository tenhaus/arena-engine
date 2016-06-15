package www

import (
  "os"
  "testing"
  "net/http"
  "net/url"
  "strings"
  "io/ioutil"
  "encoding/json"
  "net/http/httptest"
  "github.com/tenhaus/botpit/accounts/bot"
)

var email string
var handle string
var password string

func TestMain(m *testing.M) {
  email = "r2estsignup@test.com"
  password = "t2est"
  handle = "r2estsignup"

  code := m.Run()

  os.Exit(code)
}

func TestSignup(t *testing.T) {
  ts := httptest.NewServer(http.HandlerFunc(SignupHandler))
  defer ts.Close()

  form := url.Values{}
  form.Add("email", email)
  form.Add("password", password)
  form.Add("handle", handle)

  reader := strings.NewReader(form.Encode())
  res, err := http.Post(ts.URL, "application/x-www-form-urlencoded", reader)

  if err != nil {
    t.Error(err)
  }

  accountJson, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		t.Error(err)
	}

  if res.StatusCode != 200 {
    t.Errorf(res.Status)
    return
  }

  var account bot.ServiceAccount
  json.Unmarshal(accountJson, &account)

  t.Errorf("Name %v", account)
}
