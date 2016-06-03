package useraccounts

import (
  "fmt"
  "google.golang.org/cloud/datastore"
  "github.com/tenhaus/botpit/config"
)

const USER_ACCOUNT_KIND = "Fighter"

type UserAccount struct {
  Key string
  Handle string
  Email string
  Password string
  RoutingTopic string
}

func Create(handle string, account *UserAccount) error {
  handleLength := len(handle)

  if handleLength < 6 || handleLength > 30 {
    return fmt.Errorf("Handle must be between 6 and 30 characters")
  }

  client, context := config.GetClientWithContext()

  k := datastore.NewKey(context, USER_ACCOUNT_KIND, "", 0, nil)
  account.Handle = handle

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
