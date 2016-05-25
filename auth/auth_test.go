package auth

import "testing"

func TestCreateUserAccount(t *testing.T) {
  encodedId, err := CreateUserAccount("NecroPorkBopper")

  if(err != nil || encodedId == "") {
    t.Errorf("Couldn't create the account", err)
  }


}
