package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

type webPage struct {
  url string
  body []byte
  err error
}

func (w *webPage) get() {
  resp, err := http.Get(w.url)

  if err != nil {
    return
  }

  defer resp.Body.Close()

  w.body, err = ioutil.ReadAll(resp.Body)
  if err != nil {
    w.err = err
  }
}

func main() {
  w := &webPage{url: "http://www.google.com/"}
  w.get()
  fmt.Printf("URL: %s Error: %s Length of the body is %d\n", w.url, w.err, len(w.body))
}

