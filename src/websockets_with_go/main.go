package main

import (
  "net/http"
  "log"
  "github.com/negroni"
  "github.com/gorilla/pat"
  "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
  ReadBufferSize: 1024,
  WriteBufferSize:1024,
}

type Message struct {
  Type string       `json:"type"`
  Data interface{}  `json:"data"`
}

func WSHandler(res http.ResponseWriter, req *http.Request) {
  conn, err := upgrader.Upgrade(res, req, nil)
  if err != nil {
    log.Println(err)
    return
  }

  err = conn.WriteJSON(Message{Type:"welcome", Data: "Welcome to Websockets!"})
  if err != nil {
    log.Println(err)
    return
  }

  for {
    m:= &Message{}
    err = conn.ReadJSON(m)
    if err != nil {
      log.Println(err)
      return
    }

    err = conn.WriteJSON(m)
    if err != nil {
      log.Println(err)
      return
    }
  }
}

func main() {
  p := pat.New()
  p.Get("/ws", WSHandler)

  n := negroni.New()
  n.Use(negroni.NewStatic(http.Dir("/Users/Thunderbird/bbarbour/Dropbox/projects/HelloGo/src/websockets_with_go/public")))
  n.UseHandler(p)
  http.ListenAndServe(":3000", n)
}
