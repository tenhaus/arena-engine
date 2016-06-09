package user

import (
  "fmt"
  "google.golang.org/cloud/datastore"
  "github.com/tenhaus/botpit/config"
  "golang.org/x/crypto/bcrypt"
)

const USER_ACCOUNT_KIND = "Fighter"

type UserAccount struct {
  Key string
  Handle string
  Email string
  Password []byte
  RoutingTopic string
}

func Encrypt(password string) []byte {
  bPass := []byte(password)
  hashedPassword, _ := bcrypt.GenerateFromPassword(bPass,
    bcrypt.DefaultCost)

  return hashedPassword
}

func Create(handle string, email string,
  password string, account *UserAccount) error {
  // Test handle length
  handleLength := len(handle)

  if handleLength < 6 || handleLength > 30 {
    return fmt.Errorf("Handle must be between 6 and 30 characters")
  }

  hashedPass := Encrypt(password)
  client, context := config.GetClientWithContext()

  k := datastore.NewKey(context, USER_ACCOUNT_KIND, "", 0, nil)
  account.Handle = handle
  account.Email = email
  account.Password = hashedPass

  if key, err := client.Put(context, k, account); err != nil {
    return err
  } else {
    account.Key = key.Encode()
  }

  return nil
}

func Delete(encodedId string) error {
  cfg := config.GetConfig()
  context, _ := config.GetContext()
  client, clientErr := datastore.NewClient(context, cfg.ProjectId)

  if clientErr != nil {
    return clientErr
  }

  k, decodeError := datastore.DecodeKey(encodedId)

  if decodeError != nil {
    return decodeError
  }

  deleteError := client.Delete(context, k)
  return deleteError
}
