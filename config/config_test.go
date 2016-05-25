package config

import "os"
import "testing"

func TestGet(t *testing.T) {
  config := Get()
  environment := os.Getenv("BOTPIT_ENV")

  if environment == "development" &&
     config.RoutingSubscription != "pitmaster-dev" {
      t.Errorf("Wrong config for dev environment", environment, config)
  }

  if environment == "production" &&
     config.RoutingSubscription != "pitmaster" {
      t.Errorf("Wrong config for prod environment", environment, config)
  }
}
