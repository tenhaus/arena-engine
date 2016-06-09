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

func HandleExists(handle string) error {
  client, context := config.GetClientWithContext()

  var results []UserAccount
  q := datastore.NewQuery("Fighter").Filter("Handle =", handle)
  _, err := client.GetAll(context, q, &results)

  if len(results) > 0 {
    return fmt.Errorf("Handle already exists")
  }

  return err
}

func EmailExists(email string) error {
  client, context := config.GetClientWithContext()

  var results []UserAccount
  q := datastore.NewQuery("Fighter").Filter("Email =", email)
  _, err := client.GetAll(context, q, &results)

  if len(results) > 0 {
    return fmt.Errorf("Email already exists")
  }

  return err
}

func Create(handle string, email string,
  password string, account *UserAccount) error {

  // Test handle length
  handleLength := len(handle)

  if handleLength < 6 || handleLength > 30 {
    return fmt.Errorf("Handle must be between 6 and 30 characters")
  }

  // Make sure handle doesn't already exists
  if err := HandleExists(handle); err != nil {
    return err
  }

  // Make sure email doesn't already exists
  if err := EmailExists(email); err != nil {
    return err
  }

  // Add the account
  hashedPass := Encrypt(password)
  client, context := config.GetClientWithContext()

  k := datastore.NewKey(context, USER_ACCOUNT_KIND, "", 0, nil)
  account.Handle = handle
  account.Email = email
  account.Password = hashedPass

  key, err := client.Put(context, k, account);

  if err == nil {
    account.Key = key.Encode()
  }

  return err
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
