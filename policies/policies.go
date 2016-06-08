package policies

import (
  "fmt"
  "encoding/json"
  "google.golang.org/cloud/pubsub"
  "github.com/tenhaus/botpit/http"
  "github.com/tenhaus/botpit/config"
)

const SUBSCRIBE_ROLE =  "roles/pubsub.subscriber"
const PUBLISH_ROLE =  "roles/pubsub.publisher"

func GetPolicyForTopic(topicName string, policy *Policy) error {
  cfg := config.GetConfig()

  urlTemplate := "https://pubsub.googleapis.com/v1/projects/%s/topics/%s:getIamPolicy"
  apiUrl := fmt.Sprintf(urlTemplate, cfg.ProjectId, topicName)
  resp, err := http.Get(apiUrl, pubsub.ScopePubSub)

  if err != nil {
    return err
  }

  return json.Unmarshal(resp, policy)
}

func GrantPublish(topicName string, accountId string) error {
  return grantRole(topicName, accountId, PUBLISH_ROLE);
}

func RevokePublish(topicName string, accountId string) error {
  return revokeRole(topicName, accountId, PUBLISH_ROLE);
}

func GrantSubscribe(topicName string, accountId string) error {
  return grantRole(topicName, accountId, SUBSCRIBE_ROLE);
}

func RevokeSubscribe(topicName string, accountId string) error {
  return revokeRole(topicName, accountId, SUBSCRIBE_ROLE);
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

func RemoveAccountFromPolicy(accountId string, role string, policy *Policy) {
  saAccountId := getServiceAccountString(accountId)

  // If the policy doesn't have this role our work is already done
  if !policy.Bindings.contains(role) {
    return
  }

  bindingIndex := policy.Bindings.getBindingWithRole(role)
  binding := policy.Bindings[bindingIndex]

  // If the account isn't in the members list our work is done
  if !binding.Members.contains(saAccountId) {
    fmt.Println("Not Found", binding.Members, saAccountId)
    return
  }

  binding.Members = binding.Members.remove(saAccountId)
  policy.Bindings[bindingIndex] = binding
}

func AddRoleToPolicy(role string, policy *Policy) {
  if !policy.Bindings.contains(role) {
    binding := PolicyBinding{Role: role}
    policy.Bindings = append(policy.Bindings, binding)
  }
}

func revokeRole(topicName string, accountId string, role string) error {
  // Get the policy
  var policy Policy
  if err := GetPolicyForTopic(topicName, &policy); err != nil {
    return err
  }

  RemoveAccountFromPolicy(accountId, role, &policy)

  // Commit the policy
  urlTemplate := "https://pubsub.googleapis.com/v1/projects/%s/topics/%s:setIamPolicy"
  apiUrl := fmt.Sprintf(urlTemplate, cfg.ProjectId, topicName)

  policyWrapper := PolicyWrapper{Policy: policy}
  postData, _ := json.Marshal(policyWrapper)

  _, err := http.Post(apiUrl, postData, pubsub.ScopePubSub);
  return err
}

func grantRole(topicName string, accountId string, role string) error {

  // Get the policy
  var policy Policy
  if err := GetPolicyForTopic(topicName, &policy); err != nil {
    return err
  }

  // Add the account to the policy object
  AddAccountToPolicy(accountId, role, &policy)

  // Commit the policy
  urlTemplate := "https://pubsub.googleapis.com/v1/projects/%s/topics/%s:setIamPolicy"
  apiUrl := fmt.Sprintf(urlTemplate, cfg.ProjectId, topicName)

  policyWrapper := PolicyWrapper{Policy: policy}
  postData, _ := json.Marshal(policyWrapper)

  _, err := http.Post(apiUrl, postData, pubsub.ScopePubSub);
  return err
}
