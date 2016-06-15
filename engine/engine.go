package engine

import (
  "os"
  "fmt"
)

func Start() {
  environment := os.Getenv("ARENA_ENV")
  config := os.Getenv("ARENA_CONFIG")

  fmt.Println("Start", environment, config)
}
