package serviceaccounts

import (
  "testing"
)

func TestCreateDeleteServiceAccount(t *testing.T) {
  var account ServiceAccount
  serviceAccountHandle := "createserviceaccounttest"

  createError := Create(serviceAccountHandle, &account)

  if createError != nil {
    t.Errorf("Error creating the test account", createError)
    return
  }

  deleteError := Delete(account.Email)

  if deleteError != nil {
    t.Errorf("Error deleting the test account", deleteError)
    return
  }
}
