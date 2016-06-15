package engine

import (
  "os"
  "fmt"
)

func Start() {
  environment := os.Getenv("ARENA_ENV")

  fmt.Println("Start -", environment)
}
