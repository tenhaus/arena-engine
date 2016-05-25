package config

import (
  "os"
  "fmt"
  "io/ioutil"
  "encoding/json"
  "google.golang.org/cloud"
  "golang.org/x/net/context"
  "golang.org/x/oauth2/google"
  "google.golang.org/cloud/pubsub"
  "google.golang.org/cloud/storage"
  "google.golang.org/cloud/datastore"
)

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

func GetClientWithContext() (*datastore.Client, context.Context) {
  client, _ := GetClient()
  context, _ := GetContext()

  return client, context
}

func GetClient() (*datastore.Client, error) {
  cfg := GetConfig()
  context, _ := GetContext()
  client, clientErr := datastore.NewClient(context, cfg.ProjectId)

  if clientErr != nil {
    fmt.Println(clientErr)
    os.Exit(1)
  }

  return client, nil
}

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

func GetIAMContext() (context.Context, error) {
  config := GetConfig()
  ctx := context.Background()
	httpClient, err := google.DefaultClient(
    ctx, "https://www.googleapis.com/auth/iam", "https://www.googleapis.com/auth/cloud-platform")

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
