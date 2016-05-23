package main

import "fmt"
import "time"
import "github.com/tenhaus/botpit/config"
import "github.com/tenhaus/botpit/bus"

func main() {
  config := config.ConfigForEnvironment("development")
  bus.Authenticate(config.Key)
  
  run()
}

func run() {
  timer := time.Tick(100 * time.Millisecond)

  for range timer {
    // Check for new game requests
    // Check the status of existing games
    // Create or close games if needed

    fmt.Println("0 Games 0 Players")
  }
}
