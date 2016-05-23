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
  Key string
  ProjectId string
  RoutingTopic string
}

type MappedConfiguration map[string]EnvironmentConfiguration

func ConfigForEnvironment(environment string) EnvironmentConfiguration {
  file, readError := ioutil.ReadFile("./config.json")
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
