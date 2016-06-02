package serviceaccounts

import (
  "testing"
)

// Make sure create and delete work
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

// Make sure short names produce an error
func TestRejectShortName(t *testing.T) {
  t.Error()
}

// Make sure long names produce an error
func TestRejectLongName(t *testing.T) {
  t.Error()
}
