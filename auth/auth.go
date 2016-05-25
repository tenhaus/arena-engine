// Manages users and permissions

// So I think we want to:
//
// Create a uuid for a user (probably pull it from a db)
// Create a service account with the uuid
// Create a topic with the uuid
// Somehow retrieve some token the user can use to authenticate
// Send the token back


package auth

import (
  "google.golang.org/cloud/pubsub"
  "google.golang.org/cloud/datastore"
  "golang.org/x/net/context"
  "github.com/tenhaus/botpit/config"
)

type Fighter struct {
  Handle string
}

func CreateUserAccount(handle string) (string, error) {
  cfg := config.GetConfig()
  context, _ := config.GetContext()
  client, clientErr := datastore.NewClient(context, cfg.ProjectId)

  if clientErr != nil {
    return "", clientErr
  }

  k := datastore.NewKey(context, "Fighter", "", 0, nil)
  e := new(Fighter)
  e.Handle = handle

  key, putError := client.Put(context, k, e)

  if putError != nil {
    return "", putError
  }

  return key.Encode(), nil
}

func CreateServiceAccount(uuid string) {
}

func CreateTopic(context context.Context, uuid string) {
  pubsub.CreateTopic(context, uuid)
}
