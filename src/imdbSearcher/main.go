package main

import (
  "bufio"
  "io"
  "fmt"
  "sync"
  "os"
)

var ActorNames = []string{}

func Run(in stringReader, out io.Writer) {
  ActorNames = []string{}
  AskForNames(in)
  actors := []Actor{}
  m := sync.Mutex{}

  fmt.Fprintf(out, "You selected the following %d actors:\n", len(ActorNames))

  var wg sync.WaitGroup
  wg.Add(len(ActorNames))

  for i := 0; i < len(ActorNames); i++ {
    go func(wg *sync.WaitGroup, index int) {
      actor, err := FetchActor(ActorNames[index])
      if err != nil {
        fmt.Fprintln(out, err)
      } else {
        m.Lock()
        actors = append(actors, actor)
        m.Unlock()
        fmt.Fprintln(out, actor.Name)
      }

      wg.Done()
    }(&wg, i)
  }

  wg.Wait()

  if len(actors) < 2 {
    fmt.Fprintln(out, "\nLess than two actors found. nothing to filter out")
    os.Exit(1)
  }

  credits := FilterCredits(actors)
  if len(credits) > 0 {
    fmt.Fprintln(out, "\nThey have appeared in the following movies/TV shows together:")
    for _, c := range credits {
      fmt.Fprintln(out, c.NameOrTitle())
    }
  } else {
    fmt.Fprintln(out, "\nHave not appeared in anything together")
  }
}


func main() {
  Run(bufio.NewReader(os.Stdin), os.Stdout)
}
