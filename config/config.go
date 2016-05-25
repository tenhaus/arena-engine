package config

import "os"
import "fmt"
import "io/ioutil"
import "encoding/json"
import "google.golang.org/cloud"
import "golang.org/x/net/context"
import "golang.org/x/oauth2/google"
import "google.golang.org/cloud/pubsub"
import "google.golang.org/cloud/storage"

type Configuration struct {
  Development EnvironmentConfiguration
  Production EnvironmentConfiguration
}

type EnvironmentConfiguration struct {
  ProjectId string
  RoutingTopic string
  RoutingSubscription string
  UserBucket string
}

type MappedConfiguration map[string]EnvironmentConfiguration

func GetConfig() EnvironmentConfiguration {
  environment := os.Getenv("BOTPIT_ENV")
  config := configForEnvironment(environment)
  return config
}

func GetContext() (context.Context, error) {
  config := GetConfig()
  ctx := context.Background()
	httpClient, err := google.DefaultClient(
    ctx, storage.ScopeFullControl, pubsub.ScopePubSub)

	if err != nil {
		return nil, err
	}

	return cloud.WithContext(ctx, config.ProjectId, httpClient), nil
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
