package main

import (
  "fmt"
  "github.com/gopherjs/jquery"
)

var jQuery = jquery.NewJQuery

func main() {
  m := jQuery("#main").SetHtml("")
  for i := 0; i < 100; i++ {
    m.Append(fmt.Sprintf("<p>%d</p>",i))
  }
}
