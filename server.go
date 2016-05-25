package main

import "fmt"
import "time"
// import "github.com/tenhaus/botpit/bus"
import "github.com/tenhaus/botpit/auth"

func main() {
  // routingChannel := make(chan string)
  // bus.OpenPit(routingChannel)

  url, _ := auth.CreateServiceAccount("asdf")
  fmt.Println(url)

  // run(routingChannel)
}

func run(routingChannel chan string) {
  timer := time.Tick(100 * time.Millisecond)
  fmt.Println("Running");

  for range timer {
    msg := <-routingChannel
    fmt.Println(msg)

    // Check for new game requests
    // Check the status of existing games
    // Create or close games if needed

    // fmt.Println("0 Games 0 Players")
  }
}
