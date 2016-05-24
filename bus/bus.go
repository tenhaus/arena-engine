// Handle communication between clients and the pit
package bus

import "os"
import "fmt"
import "log"
import "golang.org/x/net/context"
import "google.golang.org/cloud/pubsub"
import "github.com/tenhaus/botpit/cloud"

func OpenPit(projectId string, routingTopic string, subscription string, routingChannel chan string) {
  context, err := cloud.CloudContext(projectId)

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
