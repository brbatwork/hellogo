package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func getPage(url string) (int, error) {
  resp, err := http.Get(url)

  if err != nil {
    return 0, err
  }

  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    return 0, err
  }

  return len(body), nil

}

func getter(url string, size chan string) {
  pageLength, err := getPage(url)

  if err == nil {
    size <- fmt.Sprintf("%s has length %d\n", url, pageLength)
  }
}

func main() {
  urls := []string {"http://www.yahoo.com", "http://www.google.com", "http://www.apple.com", "http://bbc.co.uk"}
  size := make(chan string)

  for _, url := range urls {
    go getter(url, size)
  }

  for i := 0; i < len(urls); i++ {
    fmt.Printf("%s\n", <- size)    
  }
}
