package main

import "fmt"
import "time"
// import "github.com/tenhaus/botpit/bus"
import "github.com/tenhaus/botpit/accounts"

func main() {
  // routingChannel := make(chan string)
  // bus.OpenPit(routingChannel)
  var account accounts.ServiceAccount
  testAccountName := "testeisadorkdfuhdddddbibcgrus"
  createError := accounts.CreateServiceAccount(testAccountName, &account)
  if createError == nil {
    fmt.Println(account)
    err := accounts.DeleteServiceAccount(account.UniqueId)
    fmt.Println(err)
  } else {
    fmt.Println(createError)
  }


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
