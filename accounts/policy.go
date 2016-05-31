package accounts

import "fmt"

type Policy struct {
  Version int
  Bindings PolicyBindings
  Etag string
}

type PolicyBindings []PolicyBinding

type PolicyBinding struct {
  Role string
  Members PolicyMembers
}

type PolicyMembers []string

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
    saMember := fmt.Sprintf("serviceAccount:%s", member)
    if saMember == existingMember {
      return true
    }
  }

  return false
}
