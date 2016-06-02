package accounts

import (
  "fmt"
  "testing"
  "github.com/tenhaus/botpit/config"
)

func TestCreateAndDeleteUserAccount(t *testing.T) {
  encodedId, createError := CreateUserAccount("NecroPorkBopper")

  if createError != nil || encodedId == "" {
    t.Errorf("Couldn't create the account", createError)
  }

  deleteError := DeleteUserAccount(encodedId)

  if deleteError != nil {
    t.Errorf("Couldn't delete the account", deleteError)
  }
}
