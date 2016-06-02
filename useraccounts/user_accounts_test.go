package useraccounts

import (
  "testing"
)

func TestCreateDeleteUserAccount(t *testing.T) {
  var account UserAccount

  if err := Create("TestyTesterson1134", &account); err != nil {
    t.Error(err)
  }

  if err := Delete(account.Key); err != nil {
    t.Error(err)
  }
}
