package main

import (
  "github.com/tenhaus/botpit/www"
  "net/http"
)

func main() {
  www.Serve()
  http.ListenAndServe(":8000", nil)
}
