package main

import (
  "log"
  "bytes"
  "strings"
  "os/exec"
)

func main() {
  path, err := exec.LookPath("ruby")
  if err != nil {
    log.Fatal(err)
  }

  log.Printf("Ruby can be found at: %s", path)
  cmd := exec.Command(path, "-e", `puts "What's your name?"; n = gets.chomp; puts "Hi #{n}!"`)
  cmd.Stdin = strings.NewReader("Bill")

  var out bytes.Buffer
  cmd.Stdout = &out

  err = cmd.Run()

  if err != nil {
    log.Fatal(err)
  }

  log.Println(out.String())
}
