package main

import (
  "net/http"
  "log"
  "github.com/negroni"
  "os/exec"
  "github.com/gorilla/pat"
  "github.com/gorilla/websocket"
)

var Connections = map[string]*websocket.Conn{}

func addConection(conn *websocket.Conn) string {
  uuid, err := getUUID()

  if err != nil {
    log.Fatal("Problem generating uuid", err)
  }

  Connections[uuid] = conn
  return uuid
}

func removeConnection(uuid string) {
  log.Printf("Removing connection %s", uuid)
  delete(Connections, uuid)
}

func sendMessage(m *Message) {
  for uuid, conn := range Connections {
    go func(m *Message, u string, c *websocket.Conn) {
      err := c.WriteJSON(m)
      if err != nil {
        log.Println(err)
        removeConnection(u)
      }
    }(m, uuid, conn)
  }
}

func getUUID()(string, error) {
  out, err := exec.Command("uuidgen").Output()
  if err != nil {
    return "", err
  }

  return string(out), nil
}
var upgrader = websocket.Upgrader{
  ReadBufferSize: 1024,
  WriteBufferSize:1024,
}

type Message struct {
  Type string       `json:"type"`
  Data interface{}  `json:"data"`
  UUID string       `json:"uuid"`
}

func WSHandler(res http.ResponseWriter, req *http.Request) {
  conn, err := upgrader.Upgrade(res, req, nil)
  if err != nil {
    log.Println(err)
    return
  }

  uuid := addConection(conn)

  err = conn.WriteJSON(Message{
    Type:"welcome",
    Data: "Welcome to Websockets!",
    UUID: uuid,
  })

  if err != nil {
    log.Println(err)
    removeConnection(uuid)
    return
  }

  for {
    m:= &Message{}
    err = conn.ReadJSON(m)
    if err != nil {
      log.Println(err)
      removeConnection(uuid)
      return
    }

    m.UUID = uuid
    sendMessage(m)
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
