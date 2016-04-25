package main

import (
    "fmt"
    "os"
)

func printer(words []string) {
  for _, word := range words {
    fmt.Printf("%s", word)
  }
  fmt.Println()
}

func main() {
  f, err := os.Open("test.txt")

  if err != nil {
    fmt.Printf("%s\n", err)
    os.Exit(1)
  }

  defer f.Close()

  b := make([]byte, 100)

  n ,err := f.Read(b)

  fmt.Printf("%d: % x\n", n, b)

  // Convert to string
  stringOut := string(b)
  fmt.Println(stringOut)

}
