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
}

type MappedConfiguration map[string]EnvironmentConfiguration

func ConfigForEnvironment(environment string) EnvironmentConfiguration {
  file, e := ioutil.ReadFile("./config.json")
  if e != nil {
    fmt.Println("Couldn't read config %v\n", e)
    os.Exit(1);
  }

  var config MappedConfiguration
  err := json.Unmarshal(file, &config)

  if err != nil {
    fmt.Println("Couldn't parse config %v\n", e)
    os.Exit(1);
  }

  return config[environment]
}
