package config

import "os"
import "testing"

func TestGet(t *testing.T) {
  config := GetConfig()
  environment := os.Getenv("BOTPIT_ENV")

  if environment == "" {
    t.Errorf("No environment set at $BOTPIT_ENV")
  }

  if environment == "development" &&
     config.RoutingSubscription != "pitmaster-dev" {
      t.Errorf("Wrong config for dev environment", environment, config)
  }

  if environment == "production" &&
     config.RoutingSubscription != "pitmaster" {
      t.Errorf("Wrong config for prod environment", environment, config)
  }
}

func TestGetContext(t *testing.T) {
  _, err := GetContext()

  if err != nil {
    t.Errorf("Error getting context", err)
  }
}
