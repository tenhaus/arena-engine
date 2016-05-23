// Handle communication between clients and the pit
package bus

import "google.golang.org/cloud/pubsub"
import "golang.org/x/net/context"
import "golang.org/x/oauth2/google"
import "google.golang.org/cloud"
import "google.golang.org/cloud/storage"
import "os"
import "fmt"
import "log"

func OpenPit(projectId string, routingTopic string, subscription string, routingChannel chan string) {
  context, err := cloudContext(projectId)
  if err != nil {
    fmt.Println("Error creating context", err)
    os.Exit(1)
  }

  // Create the topic for routing incoming game requests
  // and subscribe to it

  pubsub.CreateTopic(context, routingTopic)
  pubsub.CreateSub(context, subscription, routingTopic, 0, "")
  go subscribe(context, subscription, routingChannel);
}

func subscribe(context context.Context, subscription string, routingChannel chan string) {

  for {
    // infinite loop while we blockwait for messages
    msgs, err := pubsub.PullWait(context, subscription, 10)

    if err != nil {
      log.Fatalf("could not pull: %v", err)
    }

    for _, m := range msgs {
      routingChannel <- string(m.Data)
      pubsub.Ack(context, subscription, m.AckID)
    }
  }

}

func cloudContext(projectId string) (context.Context, error) {
  ctx := context.Background()
	httpClient, err := google.DefaultClient(ctx, storage.ScopeFullControl, pubsub.ScopePubSub)
	if err != nil {
		return nil, err
	}
	return cloud.WithContext(ctx, projectId, httpClient), nil
}
