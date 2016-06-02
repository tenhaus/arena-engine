package accounts

import (
  "testing"
  "strings"
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

func TestAddAccountToPolicyWithExistingRole(t *testing.T) {
  binding := PolicyBinding{Role: "roles/pubsub.subscriber"}
  bindings := PolicyBindings{binding}
  policy := Policy{Bindings: bindings}

  accountId := "test@test.com"
  role := "roles/pubsub.subscriber"
  AddAccountToPolicy(accountId, role, &policy)

  member := policy.Bindings[0].Members[0]
  if !strings.Contains(member, accountId) {
    t.Errorf("Failed to add a member")
  }
}

func TestAddAccountToPolicyWithoutExistingRole(t *testing.T) {
  bindings := PolicyBindings{}
  policy := Policy{Bindings: bindings}

  accountId := "test@test.com"
  role := "roles/pubsub.subscriber"
  AddAccountToPolicy(accountId, role, &policy)

  member := policy.Bindings[0].Members[0]
  if !strings.Contains(member, accountId) {
    t.Errorf("Failed to add a member")
  }
}

func TestGrantSubscribe(t *testing.T) {
  cfg := config.GetConfig()

  // Create a service account
  var serviceAccount ServiceAccount;
  handle := "testytesterson1134"
  accountId := "testytesterson1134@botpit-1134.iam.gserviceaccount.com"

  err := CreateServiceAccount(handle, &serviceAccount)

  if err != nil {
    t.Error(err)
  }

  // Grant subscribe to the service account
  err = GrantSubscribe(cfg.RoutingTopic, accountId)

  if err != nil {
    t.Error(err)
  }

  // Delete the service account
  err = DeleteServiceAccount(accountId)

  if err != nil {
    t.Error(err)
  }

}
