// Manages users and permissions

// So I think we want to:
//
// Create a uuid for a user (probably pull it from a db)
// Create a service account with the uuid
// Create a topic with the uuid
// Somehow retrieve some token the user can use to authenticate
// Send the token back


package accounts

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "golang.org/x/oauth2/google"
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

type ServiceAccount struct {
  Name string
  ProjectId string
  UniqueId string
  Email string
  DisplayName string
  Etag string
  Oauth2ClientId string
}

func CreateServiceAccount(handle string, account *ServiceAccount) error {
  handleLength := len(handle)

  if handleLength < 6 || handleLength > 30 {
    return fmt.Errorf("Handle must be between 6 and 30 characters")
  }

  cfg := config.GetConfig()
  context := context.Background()

  // Build the url and parameters
  apiUrl := fmt.Sprintf("https://iam.googleapis.com/v1/projects/%s/serviceAccounts", cfg.ProjectId)
  jsonParameters := fmt.Sprintf(`{"accountId": "%s", "serviceAccount": {"displayName": "%s"}}`, handle, handle)
  b := strings.NewReader(jsonParameters)
  
  // Run the request
  request, _ := http.NewRequest("POST", apiUrl, b)

  client, _ := google.DefaultClient(
    context, "https://www.googleapis.com/auth/cloud-platform")

  resp, err := client.Do(request)

  if err != nil {
    return err
  }

  if(resp.StatusCode == 409) {
    return fmt.Errorf("Name already exists")
  }

  if(resp.StatusCode == 200) {
    contents, _ := ioutil.ReadAll(resp.Body)
    jsonError := json.Unmarshal(contents, &account)
    return jsonError
  }


  // 200 OK
  // 409 w/ status ALREADY_EXISTS

  return fmt.Errorf("Unknown error while creating the account. Sorry.")
}

func DeleteServiceAccount(encodedId string) error {
  cfg := config.GetConfig()
  context := context.Background()


  // Build the url and parameters
  urlTemplate := "https://iam.googleapis.com/v1/projects/%s/serviceAccounts/%s"
  apiUrl := fmt.Sprintf(urlTemplate, cfg.ProjectId, encodedId)

  // Run the request
  request, _ := http.NewRequest("DELETE", apiUrl, nil)

  client, _ := google.DefaultClient(
    context, "https://www.googleapis.com/auth/cloud-platform")

  resp, err := client.Do(request)

  if err != nil {
    return err
  }

  fmt.Println(resp.StatusCode)
  contents, _ := ioutil.ReadAll(resp.Body)
  fmt.Println(string(contents))
  return nil
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
