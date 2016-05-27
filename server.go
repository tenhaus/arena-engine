package main

import (
  "fmt"
  "time"
  "github.com/tenhaus/botpit/bus"
  "github.com/tenhaus/botpit/controller"
)

func main() {
  controlChannel := make(chan string)

  go bus.OpenPit(controlChannel)
  go controller.Start(controlChannel)

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
