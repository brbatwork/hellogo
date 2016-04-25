package main

import (
    "fmt"
)

func emit(c chan string) {
  words := []string{"The", "quick", "brown", "fox"}

  for _, word := range words {
    c <- word
  }

  close(c)
}

func makeId(c chan int) {
  var id int
  id = 0

  for {
    c <- id
    id += 1
  }
}

func main() {
  wordChannel := make(chan string)
  idChannel := make(chan int)
  
  // concurrently run emit
  go emit(wordChannel)
  go makeId(idChannel)

  // word := <- wordChannel

  // fmt.Printf("%s ", word)
  // fmt.Println()

  fmt.Printf("%d\n", <- idChannel)
  fmt.Printf("%d\n", <- idChannel)
  fmt.Printf("%d\n", <- idChannel)
}
