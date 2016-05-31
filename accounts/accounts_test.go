package accounts

import (
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

func TestGetPolicy(t *testing.T) {
  cfg := config.GetConfig()
  var policy Policy
  err := GetPolicyForTopic(cfg.RoutingTopic, &policy)

  if err != nil {
    t.Errorf("Error fetching a policy", err)
  }
}

func TestAddAccountToPolicy(t *testing.T) {
  cfg := config.GetConfig()
  var policy Policy
  GetPolicyForTopic(cfg.RoutingTopic, &policy)

  accountId := "test@test.com"
  role := "roles/pubsub.subscriber"
  AddAccountToPolicy(accountId, role, &policy)

  // if err != nil {
  //   t.Errorf("Could not add the account to the role", err)
  // }
}
