package main

import (
  "net/http"
  "github.com/tenhaus/botpit/rest"
)

func main() {
  rest.UseMe()
  http.ListenAndServe(":8000", nil)
}
