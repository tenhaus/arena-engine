package accounts

import "testing"

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

func TestCreateServiceAccount(t *testing.T) {
  var account ServiceAccount
  testAccountName := "testisadorkddsf"

  createError := CreateServiceAccount(testAccountName, &account)

  if createError != nil {
    t.Errorf("Error creating the test account", createError)
    return
  }

  deleteError := DeleteServiceAccount(account.UniqueId)

  if deleteError != nil {
    t.Errorf("Error deleting the test account", deleteError)
    return
  }
}
