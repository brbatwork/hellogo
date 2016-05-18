package main

import (
  "errors"
  "log"
)

var ErrBar = errors.New("Bar!!")

func main() {
  err := Boom()
  if err != nil {
    switch e := err.(type) {
    case *Foo:
      log.Printf("Got a foo error: %s\n", e.Error())
    default:
      log.Println(err)
    }
  }
}

func Boom() error {
  // return errors.New("Boom!")
  return &Foo{}
  return ErrBar
}

type Foo struct {

}

func (f *Foo) Error() string {
  return "Boom goes the dynamite!"
}
