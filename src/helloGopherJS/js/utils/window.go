package utils

import (
  "github.com/gopherjs/gopherjs/js"
)

type window struct {
  JS js.Object
}

func (w window) Alert(msg string) {
  w.JS.Call("alert", msg)
}

// var Window = window{JS: js.Global.Get("window")}
