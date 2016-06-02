package accounts

import (
  "os"
  "fmt"
  "strings"
  "testing"
  "github.com/tenhaus/botpit/config"
)

var cfg config.EnvironmentConfiguration
var testServiceAccountHandle string
var testServiceAccount ServiceAccount

func setup() error {
  cfg = config.GetConfig()
  testServiceAccountHandle = "testytesterson1134"
  return CreateServiceAccount(testServiceAccountHandle, &testServiceAccount)
}

// Delete the service account
// The permissions we add to topics are automatically removed when
// the account is deleted, so we're not cleaning them up in teardown
func teardown() error {
  return DeleteServiceAccount(testServiceAccount.UniqueId)
}



func TestMain(m *testing.M) {
  if err := setup(); err != nil {
    fmt.Errorf("Could not create service account in setup", err)
  }

  code := m.Run()

  if err := teardown(); err != nil {
    fmt.Errorf("Could not delete service account in teardown", err)
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
  binding := PolicyBinding{Role: "roles/pubsub.subscriber"}
  bindings := PolicyBindings{binding}
  policy := Policy{Bindings: bindings}

  role := SUBSCRIBE_ROLE
  AddAccountToPolicy(testServiceAccount.Email, role, &policy)

  member := policy.Bindings[0].Members[0]
  if !strings.Contains(member, testServiceAccount.Email) {
    t.Errorf("Failed to add a member")
  }
}

// Add the account to a policy that does not have the role defined
func TestAddAccountToPolicyWithoutExistingRole(t *testing.T) {
  bindings := PolicyBindings{}
  policy := Policy{Bindings: bindings}

  role := SUBSCRIBE_ROLE
  AddAccountToPolicy(testServiceAccount.Email, role, &policy)

  member := policy.Bindings[0].Members[0]
  if !strings.Contains(member, testServiceAccount.Email) {
    t.Errorf("Failed to add a member")
  }
}

// Grant permissions to subscribe to a topic
func TestGrantSubscribe(t *testing.T) {
  if err := GrantSubscribe(cfg.RoutingTopic, testServiceAccount.Email); err != nil {
    t.Error(err)
  }
}

// Grant permissions to publish to a topic
func TestGrantPublish(t *testing.T) {
  if err := GrantPublish(cfg.RoutingTopic, testServiceAccount.Email); err != nil {
    t.Error(err)
  }
}
