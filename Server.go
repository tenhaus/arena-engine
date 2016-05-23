// Ties it all together

package main

import "fmt"
import "time"

func main() {
  run()
}

func run() {
  timer := time.Tick(100 * time.Millisecond)

  for range timer {
    fmt.Println("Tick")
  }
}
