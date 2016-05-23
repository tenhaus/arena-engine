package config

import "fmt"

func ConfigForEnvironment(environment string) string {
  fmt.Println(environment)
  return environment
}
