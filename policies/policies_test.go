package policies

import (
  "os"
  "fmt"
  "strings"
  "testing"
  "github.com/tenhaus/botpit/config"
  "github.com/tenhaus/botpit/accounts/bot"
)

var cfg config.EnvironmentConfiguration
var serviceAccountHandle string
var serviceAccount bot.ServiceAccount

func setup() error {
  cfg = config.GetConfig()
  serviceAccountHandle = "testytesterson1134"
  return bot.Create(serviceAccountHandle, &serviceAccount)
}

// Delete the service account
// The permissions we add to topics are automatically removed when
// the account is deleted, so we're not cleaning them up in teardown
func teardown() error {
  return bot.Delete(serviceAccount.UniqueId)
}



func TestMain(m *testing.M) {
  if err := setup(); err != nil {
    fmt.Println("Could not create service account in setup", err)
    os.Exit(1)
  }

  code := m.Run()

  if err := teardown(); err != nil {
    fmt.Println("Could not delete service account in teardown", err)
    os.Exit(1)
  }

	os.Exit(code)
}



// Make sure we can fetch a policy from GC
func TestGetPolicy(t *testing.T) {
  var policy Policy
  err := GetPolicyForTopic(cfg.RoutingTopic, &policy)

  if err != nil {
    t.Errorf("Error fetching a policy", err)
  }
}

// Add the account to a policy that already has the role defined
func TestAddAccountToPolicyWithExistingRole(t *testing.T) {
  binding := PolicyBinding{Role: SUBSCRIBE_ROLE}
  bindings := PolicyBindings{binding}
  policy := Policy{Bindings: bindings}

  AddAccountToPolicy(serviceAccount.Email, SUBSCRIBE_ROLE, &policy)

  member := policy.Bindings[0].Members[0]
  if !strings.Contains(member, serviceAccount.Email) {
    t.Errorf("Failed to add a member")
  }
}

// Add the account to a policy that does not have the role defined
func TestAddAccountToPolicyWithoutExistingRole(t *testing.T) {
  bindings := PolicyBindings{}
  policy := Policy{Bindings: bindings}

  AddAccountToPolicy(serviceAccount.Email, SUBSCRIBE_ROLE, &policy)

  member := policy.Bindings[0].Members[0]
  if !strings.Contains(member, serviceAccount.Email) {
    t.Errorf("Failed to add a member")
  }
}

// Delete an account from a policy
func TestRemoveAccountFromPolicy(t *testing.T) {
  email := fmt.Sprintf("serviceAccount:%s", serviceAccount.Email)
  users:= []string{email}
  binding := PolicyBinding{Role: SUBSCRIBE_ROLE, Members: users}
  bindings := PolicyBindings{binding}
  policy := Policy{Bindings: bindings}

  RemoveAccountFromPolicy(serviceAccount.Email, SUBSCRIBE_ROLE, &policy)

  t.Errorf("Finish me")
}

// Grant permissions to subscribe to a topic
func TestGrantRevokeSubscribe(t *testing.T) {
  if err := GrantSubscribe(cfg.RoutingTopic, serviceAccount.Email); err != nil {
    t.Error(err)
  }

  if err := RevokeSubscribe(cfg.RoutingTopic, serviceAccount.Email); err != nil {
    t.Error(err)
  }
}

// Grant permissions to publish to a topic
func TestGrantRevokePublish(t *testing.T) {
  if err := GrantPublish(cfg.RoutingTopic, serviceAccount.Email); err != nil {
    t.Error(err)
  }

  if err := RevokePublish(cfg.RoutingTopic, serviceAccount.Email); err != nil {
    t.Error(err)
  }
}
