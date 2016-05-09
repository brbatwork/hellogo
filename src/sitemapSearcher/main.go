package main

import (
    "fmt"
    "net/http"
)

func main() {
  // searchUrl := "ford-earns-usd38-bil-in-q1"
  resp, err := http.Get("http://blog.ihs.com/sitemap.xml")

  if err != nil {
    fmt.Printf("%s\n", resp.Body)
  }


}
