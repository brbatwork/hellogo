package main

import (
  "github.com/fzzy/radix/redis"
  "github.com/fzzy/radix/extra/pubsub"
  "fmt"
  "time"
  "log"
)

func main() {
  client, err := redis.Dial("tcp", "localhost:6379")

  if err != nil {
    log.Fatal(err)
  }

  r := client.Cmd("SET", "foo", "bar")
  if r.Err != nil {
    log.Fatal(r.Err)
  }

  log.Printf("r.String(): %s\n", r.String())

  r = client.Cmd("GET", "foo")
  if r.Err != nil {
    log.Fatal(r.Err)
  }

  log.Printf("r.String(): %s\n", r.String())

  // queing commands and run at once
  client.Append("SET", "name", "Mark")
  client.Append("GET", "name")
  r = client.GetReply()
  log.Printf("r.String(): %s\n", r.String())
  r = client.GetReply()
  log.Printf("r.String(): %s\n", r.String())

  // Get back multiple keys
  client.Cmd("SET", "first_name", "Bill")
  client.Cmd("SET", "last_name", "Jones")
  r = client.Cmd("MGET", "first_name", "last_name")
  if r.Err != nil {
    log.Fatal(r.Err)
  }

  list, _ := r.List()
  for _, m := range list {
    log.Printf("m: %s\n", m)
  }

  // pub sub
  go func() {
    client, err := redis.Dial("tcp", "localhost:6379")
    if err != nil {
      log.Fatal(err)
    }
    i := 0
    for {
      i++
      client.Cmd("PUBLISH", "news.tech", fmt.Sprintf("This is tech story #%d", i))
      client.Cmd("PUBLISH", "news.sports", fmt.Sprintf("This is sports story #%d", i))

    }
  }()

  go func() {
    client, err := redis.Dial("tcp", "localhost:6379")
    if err != nil {
      log.Fatal(err)
    }
    sub := pubsub.NewSubClient(client)
    sr := sub.PSubscribe("news.*")
    if sr.Err != nil {
      log.Fatal(sr.Err)
    }

    for {
      r := sub.Receive()
      if r.Err != nil {
        log.Fatal(r.Err)
      }
      log.Printf("r.Message: %s\n", r.Message)
    }
  }()

  time.Sleep(1 * time.Second)

}
