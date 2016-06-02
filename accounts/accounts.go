// Manages users and permissions

package accounts

import (
  "fmt"  
  "google.golang.org/cloud/pubsub"
  "google.golang.org/cloud/datastore"
  "golang.org/x/net/context"
  "github.com/tenhaus/botpit/config"
)

type Fighter struct {
  Handle string
  Email string
  Password string
}

func CreateTopic(context context.Context, uuid string) {
  pubsub.CreateTopic(context, uuid)
}

func CreateUserAccount(handle string) (string, error) {
  handleLength := len(handle)

  if handleLength < 6 || handleLength > 30 {
    return "", fmt.Errorf("Handle must be between 6 and 30 characters")
  }

  client, context := config.GetClientWithContext()

  k := datastore.NewKey(context, "Fighter", "", 0, nil)
  e := new(Fighter)
  e.Handle = handle

  key, putError := client.Put(context, k, e)

  if putError != nil {
    return "", putError
  }

  return key.Encode(), nil
}

func DeleteUserAccount(encodedId string) error {
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
