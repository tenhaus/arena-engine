package accounts

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "golang.org/x/oauth2/google"
  "github.com/tenhaus/botpit/config"
  "golang.org/x/net/context"
  "google.golang.org/cloud/pubsub"
)

type Policy struct {
  Version int
  Bindings []PolicyBinding
  Etag string
}

type PolicyBinding struct {
  Role string
  Members []string
}

func GetPolicyForTopic(topicName string, policy *Policy) error {
  cfg := config.GetConfig()
  context := context.Background()

  // Build the url and parameters
  urlTemplate := "https://pubsub.googleapis.com/v1/projects/%s/topics/%s:getIamPolicy"
  apiUrl := fmt.Sprintf(urlTemplate, cfg.ProjectId, topicName)

  request, _ := http.NewRequest("GET", apiUrl, nil)
  client, _ := google.DefaultClient(context, pubsub.ScopePubSub)
  resp, err := client.Do(request)

  if err != nil {
    return err
  }

  if resp.StatusCode == 200 {
    contents, _ := ioutil.ReadAll(resp.Body)
    unmarshalError := json.Unmarshal(contents, policy)
    fmt.Println(Policy(*policy))

    return unmarshalError
  }

  return nil
}
