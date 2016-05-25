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
  "fmt"
  "google.golang.org/cloud/pubsub"
  "google.golang.org/cloud/datastore"
  "golang.org/x/net/context"
  "github.com/tenhaus/botpit/config"
)

type Fighter struct {
  Handle string
  Name string
  ID string
}

func CreateUserAccount(handle string) (string, error) {
  cfg := config.GetConfig()
  context, _ := config.GetContext()
  client, err := datastore.NewClient(context, cfg.ProjectId)

  if err != nil {
    return "", err
  }

  q := datastore.NewQuery("Fighter")

  for t:= client.Run(context, q);; {
    var f Fighter
    key, err := t.Next(&f)

    if err == datastore.Done {
      break
    }

    if err != nil {
      // whatever
    }

    fmt.Println(key, f.Name)
  }

  return "eh", nil
}

func CreateServiceAccount(uuid string) {
}

func CreateTopic(context context.Context, uuid string) {
  pubsub.CreateTopic(context, uuid)
}
