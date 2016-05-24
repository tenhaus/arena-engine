package auth

import "testing"

func TestSomething(t *testing.T) {
  uuid, err := CreateUser("Test User", "some-project")

  if err != nil {
    t.Errorf("Error creating user", err)
  }

  if uuid != "" {
    t.Errorf("uuid is not blank", uuid)
  }

}
