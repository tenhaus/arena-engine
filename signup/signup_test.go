package signup

import (
  "testing"
  "github.com/tenhaus/botpit/accounts/user"
  "github.com/tenhaus/botpit/accounts/bot"
)

// Signup and recieve key file, user account, service account
// Test pubsub send and receive
func TestSignup(t *testing.T) {
  var serviceAccount bot.ServiceAccount
  var userAccount user.UserAccount

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

func TestIsReservedHandle(t *testing.T) {
  if isReserved := IsReservedHandle("test_abcd"); !isReserved {
    t.Errorf("Able to sign up with reserved handle")
  }

  if isReserved := IsReservedHandle("abcd_nottest"); isReserved {
    t.Errorf("Non-reserved handle rejected")
  }
}

func TestIsReservedEmail(t *testing.T) {
  if isReserved := IsReservedEmail("test_abcd@testytesterson.com"); !isReserved {
    t.Errorf("Able to sign up with reserved email")
  }

  if isReserved := IsReservedEmail("abcd_nottest@google.com"); isReserved {
    t.Errorf("Non-reserved email rejected")
  }
}
