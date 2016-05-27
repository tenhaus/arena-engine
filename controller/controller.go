package controller

import (
  "github.com/tenhaus/botpit/bus"
)

func Start(controlChannel chan string) {
  bus.OpenPit(controlChannel)
  controlChannel <- "Running"
}
