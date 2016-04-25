package main

import (
    "fmt"
)

func emit(wordChannel chan string, done chan bool) {
  words := []string{"The", "quick", "brown", "fox"}
  i := 0

  for {
    select {
    case wordChannel <- words[i]:
      i += 1
      if i == len(words) {
        i = 0
      }
    case <- done:
      fmt.Println("Received done message")
      close(done)
      return
    }
  }
}

func main() {
  wordChannel := make(chan string)
  doneChannel := make(chan bool)
  
  // concurrently run emit
  go emit(wordChannel, doneChannel)

  for i := 0; i < 100; i++ {
    fmt.Printf("%s ", <- wordChannel)
  }

  doneChannel <- true
}
