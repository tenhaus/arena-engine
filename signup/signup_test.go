package signup

import (
	"github.com/tenhaus/botpit/accounts/bot"
	"github.com/tenhaus/botpit/accounts/user"
	"testing"
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

// Reserved handle name
func TestIsReservedHandle(t *testing.T) {
	if isReserved := IsReservedHandle("test_abcd"); !isReserved {
		t.Errorf("Able to sign up with reserved handle")
	}

	if isReserved := IsReservedHandle("abcd_nottest"); isReserved {
		t.Errorf("Non-reserved handle rejected")
	}
}

// Reserved email
func TestIsReservedEmail(t *testing.T) {
	reservedEmail := "test_abcd@testytesterson.com"
	correctEmail := "abcd_nottest@google.com"

	if isReserved := IsReservedEmail(reservedEmail); !isReserved {
		t.Errorf("Able to sign up with reserved email")
	}

	if isReserved := IsReservedEmail(correctEmail); isReserved {
		t.Errorf("Non-reserved email rejected")
	}
}
