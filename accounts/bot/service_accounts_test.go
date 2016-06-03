package bot

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

// Create a key for the account
func TestCreateKey(t *testing.T) {
  var serviceAccount ServiceAccount
  serviceAccountHandle := "createserviceaccounttest"

  if err := Create(serviceAccountHandle, &serviceAccount); err != nil {
    t.Error(err)
  }

  if err:= CreateKey(&serviceAccount); err != nil {
    t.Error(err)
  }

  if err := Delete(serviceAccount.Email); err != nil {
    t.Error(err)
  }
}
