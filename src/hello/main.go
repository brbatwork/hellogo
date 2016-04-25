package main

import (
    "fmt"
    "os"
    "errors"
)

var ( errorEmptyString = errors.New("Unwilling to print an empty string") )

func printer(msg string) error {
  if msg == "" {
    return errorEmptyString
  }
  _, err := fmt.Println("%s\n", msg)
  return err
}

func main() {
  if err := printer(""); err != nil {
    if err == errorEmptyString {
      fmt.Printf("You tried to print an empty string\n")
    } else {
      fmt.Printf("printer failed: %s\n",err)
    }
    os.Exit(1)

  }

}
