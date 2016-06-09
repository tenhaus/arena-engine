package signup

import (
  "testing"
  "github.com/tenhaus/botpit/useraccounts"
  "github.com/tenhaus/botpit/serviceaccounts"
)

// Signup and recieve key file, user account, service account
// Test pubsub send and receive
func TestSignup(t *testing.T) {
  var serviceAccount serviceaccounts.ServiceAccount
  var userAccount useraccounts.UserAccount

  err := Signup("testsignup", "testsignup@signup.com", "timisadork",
  &userAccount, &serviceAccount)

  if err != nil {
    t.Error(err)
  }

  err = KillUser(userAccount, serviceAccount)

  if err != nil {
    t.Error(err)
  }
}

func TestSignupDuplicateNameError(t *testing.T) {
  t.Errorf("Make me")
}

func TestPasswordEncryption(t *testing.T) {
  t.Errorf("Make me")
}
