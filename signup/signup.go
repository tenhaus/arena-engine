package signup

import (
	"fmt"
	"github.com/tenhaus/arena-engine/accounts/bot"
	"github.com/tenhaus/arena-engine/accounts/user"
	"github.com/tenhaus/arena-engine/bus"
	"github.com/tenhaus/arena-engine/policies"
	"strings"
)

// Create user account √
// Create service account √
// Get the key √
// Create the game routing topic √
// Set permissions √
func Signup(handle string, email string, password string,
	userAccount *user.UserAccount,
	serviceAccount *bot.ServiceAccount) error {

	// Create the user account
	if err := user.Create(handle, email,
		password, userAccount); err != nil {
		return fmt.Errorf("Error creating the user account %v", err)
	}

	// Create the service account
	if err := bot.Create(handle, serviceAccount); err != nil {
		return fmt.Errorf("Error creating the service account %v", err)
	}

	// Get the key for the service account
	if err := bot.CreateKey(serviceAccount); err != nil {
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

func KillUser(userAccount user.UserAccount,
	serviceAccount bot.ServiceAccount) error {

	// Delete the user account
	if err := user.Delete(userAccount.Key); err != nil {
		return err
	}

	// Delete the service account
	if err := bot.Delete(serviceAccount.Email); err != nil {
		return err
	}

	// Delete the user's routing topic
	if err := bus.DeleteTopic(userAccount.RoutingTopic); err != nil {
		return err
	}

	return nil
}

func IsReservedHandle(handle string) bool {
	return strings.HasPrefix(handle, "test_")
}

func IsReservedEmail(email string) bool {
	index := strings.Index(email, "testytesterson")
	hasPrefix := strings.HasPrefix(email, "test_")

	return index != -1 && hasPrefix
}
