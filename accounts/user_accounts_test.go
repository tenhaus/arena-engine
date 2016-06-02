package accounts

import (
  "testing"
)

func TestCreateDeleteUserAccount(t *testing.T) {
  var account UserAccount

  if err := CreateUserAccount("TestyTesterson1134", &account); err != nil {
    t.Error(err)
  }

  if err := DeleteUserAccount(account.Key); err != nil {
    t.Error(err)
  }
}
