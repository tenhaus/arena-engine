package accounts

import (
  "testing"
)

func TestCreateServiceAccount(t *testing.T) {
  var account ServiceAccount
  serviceAccountHandle = "createserviceaccounttest"

  createError := CreateServiceAccount(serviceAccountHandle, &account)

  if createError != nil {
    t.Errorf("Error creating the test account", createError)
    return
  }

  deleteError := DeleteServiceAccount(account.Email)

  if deleteError != nil {
    t.Errorf("Error deleting the test account", deleteError)
    return
  }
}
