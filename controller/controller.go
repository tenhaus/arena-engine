package controller

func Start(controlChannel chan string) {
  controlChannel <- "Controller Running"
}
