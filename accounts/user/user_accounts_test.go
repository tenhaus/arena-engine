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
    t.Error("Short handle worked")
  }
}

// Make sure long names (> 30) produce an error
func TestRejectLongName(t *testing.T) {
  longHandle := "1234567890123456789012345678901234567890"
  if err := Create(longHandle, "", "", nil); err == nil {
    t.Error("Long handle worked")
  }
}

// Stupid test, but helped me understand how it works
func TestEncryption(t *testing.T) {
  password := "timisadork"
  bPass := []byte(password)
  hash := Encrypt(password)

  if err := bcrypt.CompareHashAndPassword(hash, bPass); err != nil {
    t.Error(err)
  }
}

// Handle Exists works ?
func TestHandleExists(t *testing.T) {
  handle := "test_handleexists"
  email  := "test_handleexists@test.com"

  var account UserAccount
  if err := Create(handle, email, "timisadork", &account); err != nil {
    t.Error(err)
  }

  if err := HandleExists(handle); err == nil {
    t.Error("Handle exists failed")
  }

  if err := Delete(account.Key); err != nil {
    t.Error(err)
  }
}

// Email Exists works ?
func TestEmailExists(t *testing.T) {
  handle := "test_handleexists"
  email  := "test_handleexists@test.com"

  var account UserAccount
  if err := Create(handle, email, "timisadork", &account); err != nil {
    t.Error(err)
  }

  if err := EmailExists(email); err == nil {
    t.Error("Handle exists failed")
  }

  if err := Delete(account.Key); err != nil {
    t.Error(err)
  }
}

// Make sure we get an error if someone tries to use
// a handle that already exists
func TestHandleAlreadyInUse(t *testing.T) {
  handle := "test_handleinuse"
  email  := "test_handleinuse@test.com"

  var account UserAccount
  if err := Create(handle, email, "timisadork", &account); err != nil {
    t.Error(err)
  }

  if err := Create(handle, email, "timisadork", &account); err == nil {
    t.Errorf("Created the same account twice")
  }
}

// Make sure we get an error if someone tries to use
// a email that already exists
func TestEmailAlreadyInUse(t *testing.T) {
  t.Error("Make me")
}
