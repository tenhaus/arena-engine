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

    return unmarshalError
  }

  return nil
}

func AllowServiceAccountToSubscribeToTopic(topicName string, accountId string) error {
  // cfg := config.GetConfig()
  // context := context.Background()

  // Get the policy
  var policy Policy
  err := GetPolicyForTopic(topicName, &policy)

  if err != nil {
    return err
  }

  // Add the account to the policy
  // err = AddAccountToPolicy(accountId, &policy)

  // if err != nil {
    // return err
  // }

  // fmt.Println(Policy(policy))
  // Commit the policy

  return fmt.Errorf("wtf")
}

func AddAccountToPolicy(accountId string, role string, policy *Policy) {
  // If we don't already have the role, we have to add it
  if !policy.Bindings.contains(role) {
    binding := PolicyBinding{Role: role}
    policy.Bindings = append(policy.Bindings, binding)
  }

  // We know we have the role so add the account to it
  for i, binding := range policy.Bindings {
    if binding.Role == role && !binding.Members.contains(accountId) {
      s := append(binding.Members, accountId)
      binding.Members = s
    }

    policy.Bindings[i] = binding
  }
}
