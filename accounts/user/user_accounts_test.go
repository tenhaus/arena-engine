package user

import (
  "testing"
  "golang.org/x/crypto/bcrypt"
)

// Test create user, then cleanup
func TestCreateDeleteUserAccount(t *testing.T) {
  var account UserAccount

  if err := Create("TestyTesterson1134", "t@t.com", "pass", &account); err != nil {
    t.Error(err)
  }

  if err := Delete(account.Key); err != nil {
    t.Error(err)
  }
}

// Make sure short names (< 6) produce an error
func TestRejectShortName(t *testing.T) {
  if err := Create("short", "", "", nil); err == nil {
    t.Error("Short password worked")
  }
}

// Make sure long names (> 30) produce an error
func TestRejectLongName(t *testing.T) {
  longHandle := "1234567890123456789012345678901234567890"
  if err := Create(longHandle, "", "", nil); err == nil {
    t.Error("Short password worked")
  }
}

func TestEncryption(t *testing.T) {
  password := "timisadork"
  bPass := []byte(password)
  hash := Encrypt(password)

  if err := bcrypt.CompareHashAndPassword(hash, bPass); err != nil {
    t.Error(err)
  }
}
