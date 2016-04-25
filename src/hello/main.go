package main

import (
    "fmt"
    "os"
)

func printer(msg string) error {
  if msg == "" {
    return fmt.Errorf("Unwilling to print an empty string")
  }
  _, err := fmt.Println("%s\n", msg)
  return err
}

func main() {
  if err := printer(""); err != nil {
    fmt.Printf("printer failed: %s\n",err)
    os.Exit(1)
  }

}
