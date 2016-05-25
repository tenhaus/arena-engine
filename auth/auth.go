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
  "strings"
  "net/http"
  "golang.org/x/oauth2/google"
  "net/url"
  "io/ioutil"
  "google.golang.org/cloud/pubsub"
  "google.golang.org/cloud/datastore"
  "golang.org/x/net/context"
  "github.com/tenhaus/botpit/config"
)

type Fighter struct {
  Handle string
}

func CreateServiceAccount(encodedId string) (string, error) {
  cfg := config.GetConfig()
  context := context.Background()

  apiUrl := fmt.Sprintf("https://iam.googleapis.com/v1/projects/%s/serviceAccounts", cfg.ProjectId)
  form := url.Values{}
  form.Set("accountId", encodedId)

  request, _ := http.NewRequest("POST",
    apiUrl, strings.NewReader(form.Encode()))

  client, err := google.DefaultClient(
    context, "https://www.googleapis.com/auth/cloud-platform")

  if err != nil {
    return "", err
  }

  resp, _ := client.Do(request)
  contents, err := ioutil.ReadAll(resp.Body)

  fmt.Println(resp.Status)
  fmt.Println(string(contents))
  fmt.Println(err)

  return "", nil
}


func CreateTopic(context context.Context, uuid string) {
  pubsub.CreateTopic(context, uuid)
}

func CreateUserAccount(handle string) (string, error) {
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
