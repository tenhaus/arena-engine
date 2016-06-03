package policies

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "golang.org/x/oauth2/google"
  "github.com/tenhaus/botpit/config"
  "golang.org/x/net/context"
  "google.golang.org/cloud/pubsub"
)

const SUBSCRIBE_ROLE =  "roles/pubsub.subscriber"
const PUBLISH_ROLE =  "roles/pubsub.publisher"

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

func GrantPublish(topicName string, accountId string) error {
  return grantRole(topicName, accountId, PUBLISH_ROLE);
}

func RevokePublish(topicName string, accountId string) error {
  return nil
}

func GrantSubscribe(topicName string, accountId string) error {
  return grantRole(topicName, accountId, SUBSCRIBE_ROLE);
}

func RevokeSubscribe(topicName string, accountId string) error {
  return nil
}

func grantRole(topicName string, accountId string, role string) error {
  cfg := config.GetConfig()
  context := context.Background()

  // Get the policy
  var policy Policy
  err := GetPolicyForTopic(topicName, &policy)

  if err != nil {
    return err
  }

  // Add the account to the policy object
  AddAccountToPolicy(accountId, role, &policy)

  // Commit the policy
  urlTemplate := "https://pubsub.googleapis.com/v1/projects/%s/topics/%s:setIamPolicy"
  apiUrl := fmt.Sprintf(urlTemplate, cfg.ProjectId, topicName)

  policyWrapper := PolicyWrapper{Policy: policy}
  postData, _ := json.Marshal(policyWrapper)
  b := strings.NewReader(string(postData))

  request, _ := http.NewRequest("POST", apiUrl, b)
  client, _ := google.DefaultClient(context, pubsub.ScopePubSub)
  resp, err := client.Do(request)

  if err != nil {
    return err
  }

  // Yay
  if resp.StatusCode == 200 {
    return nil;
  }

  return fmt.Errorf("Couldn't do it. Handle this error.", err)
}

func AddAccountToPolicy(accountId string, role string, policy *Policy) {
  // If we don't already have the role, we have to add it
  AddRoleToPolicy(role, policy)

  // We know we have the role so add the account to it
  for i, binding := range policy.Bindings {
    if binding.Role == role && !binding.Members.contains(accountId) {
      saAccountId := getServiceAccountString(accountId)
      binding.Members = append(binding.Members, saAccountId)
    }

    policy.Bindings[i] = binding
  }
}

func AddRoleToPolicy(role string, policy *Policy) {
  if !policy.Bindings.contains(role) {
    binding := PolicyBinding{Role: role}
    policy.Bindings = append(policy.Bindings, binding)
  }
}