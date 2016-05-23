package main

import "fmt"
import "time"
import "github.com/tenhaus/botpit/config"


func main() {
  config := config.ConfigForEnvironment("development")
  fmt.Println(config.Key)

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
