package http

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
  "golang.org/x/oauth2/google"
  "golang.org/x/net/context"
)

func Post(url string, postData []byte, scope string) ([]byte, error) {
  return doRequest("POST", url, postData, scope)
}

func Get(url string, scope string) ([]byte, error){
  return doRequest("GET", url, nil, scope)
}

func doRequest(method string, url string, data []byte,
  scope string) ([]byte, error) {

  context := context.Background()
  b := strings.NewReader(string(data))

  request, _ := http.NewRequest(method, url, b)
  client, _ := google.DefaultClient(context, scope)
  resp, err := client.Do(request)
  contents, _ := ioutil.ReadAll(resp.Body)

  // If there was a network error?
  if err != nil {
    return nil, err
  }

  // - Client follows redirects so only testing 200 for now
  // - Will add if needed
  if resp.StatusCode != 200 {
    return nil, fmt.Errorf(resp.Status)
  }

  return contents, nil
}
