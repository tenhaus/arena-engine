package bot

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "golang.org/x/net/context"
  "golang.org/x/oauth2/google"
  "github.com/tenhaus/arena-engine/config"
)

func Create(handle string, account *ServiceAccount) error {
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
    return fmt.Errorf("Name already exists", resp.Body)
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

func Delete(accountId string) error {
  cfg := config.GetConfig()
  context := context.Background()

  // Build the url and parameters
  urlTemplate := "https://iam.googleapis.com/v1/projects/%s/serviceAccounts/%s"
  apiUrl := fmt.Sprintf(urlTemplate, cfg.ProjectId, accountId)

  // Run the request
  request, _ := http.NewRequest("DELETE", apiUrl, nil)

  client, _ := google.DefaultClient(
    context, "https://www.googleapis.com/auth/cloud-platform")

  resp, err := client.Do(request)

  if err != nil {
    return err
  }

  // Yay
  if resp.StatusCode == 200 {
    return nil;
  }

  return err
}

func CreateKey(serviceAccount *ServiceAccount) error {
  cfg := config.GetConfig()
  context := context.Background()

  // Build the url and parameters
  urlTemplate := "https://iam.googleapis.com/v1/projects/%s/serviceAccounts/%s/keys"
  apiUrl := fmt.Sprintf(urlTemplate, cfg.ProjectId, serviceAccount.UniqueId)
  jsonParameters := fmt.Sprintf(`{"privateKeyType": "TYPE_GOOGLE_CREDENTIALS_FILE"}`)
  b := strings.NewReader(jsonParameters)

  // Run the request
  request, _ := http.NewRequest("POST", apiUrl, b)

  client, _ := google.DefaultClient(
    context, "https://www.googleapis.com/auth/cloud-platform")

  resp, err := client.Do(request)

  if err != nil {
    return err
  }

  contents, _ := ioutil.ReadAll(resp.Body)
  if(resp.StatusCode == 200) {
    jsonError := json.Unmarshal(contents, &serviceAccount.Key)
    return jsonError
  } else {
    return fmt.Errorf("Handle this %s", string(contents))
  }

  return nil
}
