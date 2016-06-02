package policies

import (
  "os"
  "fmt"
  "strings"
  "testing"
  "github.com/tenhaus/botpit/config"
  "github.com/tenhaus/botpit/serviceaccounts"
)

var cfg config.EnvironmentConfiguration
var serviceAccountHandle string
var serviceAccount serviceaccounts.ServiceAccount

func setup() error {
  cfg = config.GetConfig()
  serviceAccountHandle = "testytesterson1134"
  return serviceaccounts.Create(serviceAccountHandle, &serviceAccount)
}

// Delete the service account
// The permissions we add to topics are automatically removed when
// the account is deleted, so we're not cleaning them up in teardown
func teardown() error {
  return serviceaccounts.Delete(serviceAccount.UniqueId)
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
  binding := PolicyBinding{Role: "roles/pubsub.subscriber"}
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

// Grant permissions to subscribe to a topic
func TestGrantSubscribe(t *testing.T) {
  if err := GrantSubscribe(cfg.RoutingTopic, serviceAccount.Email); err != nil {
    t.Error(err)
  }
}

// Grant permissions to publish to a topic
func TestGrantPublish(t *testing.T) {
  if err := GrantPublish(cfg.RoutingTopic, serviceAccount.Email); err != nil {
    t.Error(err)
  }
}
