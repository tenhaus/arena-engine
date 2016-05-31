package main

import (
  "fmt"
  "time"
  "github.com/tenhaus/botpit/bus"
  "github.com/tenhaus/botpit/controller"
  "github.com/tenhaus/botpit/accounts"
  "github.com/tenhaus/botpit/config"
)

func main() {
  controlChannel := make(chan string)
  cfg := config.GetConfig()
  go bus.OpenPit(controlChannel)
  go controller.Start(controlChannel)
  // "dev-player@botpit-1134.iam.gserviceaccount.com"
  var policy accounts.Policy
  err := accounts.GetPolicyForTopic(cfg.RoutingTopic, &policy)

  if err != nil {
    fmt.Println(err)
  }

  timer := time.Tick(100 * time.Millisecond)

  // main loop
  // Check for new game requests
  // Check the status of existing games
  // Create or close games if needed
  for range timer {
    msg := <-controlChannel
    fmt.Println(msg)
  }
}
