package signup

import (
  "fmt"
  "github.com/tenhaus/botpit/useraccounts"
  "github.com/tenhaus/botpit/serviceaccounts"
  "github.com/tenhaus/botpit/bus"
  "github.com/tenhaus/botpit/policies"
)

// Create user account √
// Create service account √
// Get the key √
// Create the game routing topic √
// Set permissions √
// Return the key √
func Signup(handle string, email string, password string,
  userAccount *useraccounts.UserAccount,
  serviceAccount *serviceaccounts.ServiceAccount) error {

  // Create the user account
  if err := useraccounts.Create(handle, userAccount); err != nil {
    return fmt.Errorf("Error creating the user account %v", err)
  }

  // Create the service account
  if err := serviceaccounts.Create(handle, serviceAccount); err != nil {
    return fmt.Errorf("Error creating the service account %v", err)
  }

  // Get the key for the service account
  if err := serviceaccounts.CreateKey(serviceAccount); err != nil {
    return fmt.Errorf("Error creating the key %v", err)
  }

  // Create the game routing topic
  if routingTopic, err := bus.CreateRoutingTopicForHandle(handle); err != nil {
    return fmt.Errorf("Error creating routing topic %v", err)
  } else {
    userAccount.RoutingTopic = routingTopic
  }

  // Grant publish to routing topic
  if err := policies.GrantPublish(userAccount.RoutingTopic,
    serviceAccount.Email); err != nil {
    return fmt.Errorf("Error granting publish %v", err)
  }

  // Grant subscribe to routing topic
  if err := policies.GrantSubscribe(userAccount.RoutingTopic,
    serviceAccount.Email); err != nil {
    return fmt.Errorf("Error granting subscribe %v", err)
  }

  return nil
}

func KillUser(userAccount useraccounts.UserAccount,
  serviceAccount serviceaccounts.ServiceAccount) error {

  // Delete the user account
  if err := useraccounts.Delete(userAccount.Key); err != nil {
    return err
  }

  // Delete the service account
  if err := serviceaccounts.Delete(serviceAccount.Email); err != nil {
    return err
  }

  // Delete the topic
  if err := bus.DeleteTopic(userAccount.RoutingTopic); err != nil {
    return err
  }

  return nil
}
