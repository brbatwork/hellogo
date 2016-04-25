package main

import (
    "fmt"
)

func printer(words []string) {
  for _, word := range words {
    fmt.Printf("%s", word)
  }
  fmt.Println()
}

func main() {
  words := make([]string, 4)
  words[0] = "the"
  words[1] = "quick"
  words[2] = "Brown"
  words[3] = "Fox"
  printer(words)

}
