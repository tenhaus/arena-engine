package accounts

import "fmt"

type PolicyWrapper struct {
  Policy Policy `json:"policy"`
}

type Policy struct {
  Version int `json:"version"`
  Bindings PolicyBindings `json:"bindings"`
  Etag string `json:"etag"`
}

type PolicyBindings []PolicyBinding

type PolicyBinding struct {
  Role string `json:"role"`
  Members PolicyMembers `json:"members"`
}

type PolicyMembers []string
type PolicyMember string

func (bindings PolicyBindings) contains(role string) bool {
  for _, binding := range bindings {
    if(binding.Role == role) {
      return true
    }
  }

  return false
}

func (members PolicyMembers) contains(member string) bool {
  for _, existingMember := range members {
    saMember := getServiceAccountString(member)
    if saMember == existingMember {
      return true
    }
  }

  return false
}

func getServiceAccountString(accountId string) string {
  return fmt.Sprintf("serviceAccount:%s", accountId)
}
