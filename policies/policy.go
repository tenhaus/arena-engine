package policies

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


func (bindings PolicyBindings) getBindingWithRole(role string) int {
  for index, binding := range bindings {
    if(binding.Role == role) {
      return index
    }
  }

  return -1
}

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
    if member == existingMember {
      return true
    }
  }

  return false
}

func (members PolicyMembers) remove(member string) []string {
  position := -1

  for i, existingMember := range members {
    if member == existingMember {
      position = i
      break
    }
  }

  if position == -1 {
    return members
  }

  return append(members[:position], members[position+1:]...)
}

func getServiceAccountString(accountId string) string {
  return fmt.Sprintf("serviceAccount:%s", accountId)
}
