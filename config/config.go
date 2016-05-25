package config

import "fmt"
import "os"
import "encoding/json"
import "io/ioutil"

type Configuration struct {
  Development EnvironmentConfiguration
  Production EnvironmentConfiguration
}

type EnvironmentConfiguration struct {
  ProjectId string
  RoutingTopic string
  RoutingSubscription string
}

type MappedConfiguration map[string]EnvironmentConfiguration

func Get() EnvironmentConfiguration {
  environment := os.Getenv("BOTPIT_ENV")
  config := configForEnvironment(environment)
  return config
}

func configForEnvironment(environment string) EnvironmentConfiguration {
  configPath := os.Getenv("BOTPIT_CONFIG")
  file, readError := ioutil.ReadFile(configPath)

  if readError != nil {
    fmt.Println("Couldn't read config %v\n", readError)
    os.Exit(1);
  }

  var config MappedConfiguration
  unmarshalError := json.Unmarshal(file, &config)

  if unmarshalError != nil {
    fmt.Println("Couldn't parse config %v\n", unmarshalError)
    os.Exit(1);
  }

  return config[environment]
}
