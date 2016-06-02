// Handle communication between clients and the pit
package bus

import (
  "os"
  "fmt"
  "log"
  "golang.org/x/net/context"
  "google.golang.org/cloud/pubsub"
  "github.com/tenhaus/botpit/config"
)

func OpenPit(routingChannel chan string) {
    cfg := config.GetConfig()
    context, err := config.GetContext()

    if err != nil {
      fmt.Println("Error creating context", err)
      os.Exit(1)
    }

    // Create the topic for routing incoming game requests
    // and subscribe to it
    pubsub.CreateTopic(context, cfg.RoutingTopic)
    pubsub.CreateSub(context, cfg.RoutingSubscription,
      cfg.RoutingTopic, 0, "")

    go subscribe(context, cfg.RoutingSubscription, routingChannel);
    routingChannel <- "Bus Running"
}

func subscribe(context context.Context, subscription string,
  routingChannel chan string) {
    // infinite loop while we blockwait for messages
    for {
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
