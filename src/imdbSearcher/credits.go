package main

import (
  "fmt"
  "net/http"
  "sync"
  "encoding/json"
)
type Credit struct {
  Id    int     `json:"id"`
  Title string  `json:"title"`
  Name  string  `json:"name"`
}

func (c Credit) NameOrTitle() string {
  if c.Name != "" {
    return c.Name
  }

  return c.Title
}

type CreditSearchResults struct {
  Cast []Credit `json:"cast"`
}

func FetchCredits(actor *Actor) error {
  u := fmt.Sprintf("%s/person/%d/combined_credits?api_key=%s", ApiRoot, actor.Id, ApiKey)
  results := CreditSearchResults{}

  res, err := http.Get(u)
  if err != nil {
    return err
  }

  err = json.NewDecoder(res.Body).Decode(&results)
  if err != nil {
    return err
  }

  actor.Credits = results.Cast
  return nil
}

func FilterCredits(actors []Actor) []Credit {
  credits := []Credit{}
  actorLength := len(actors)
  a := actors[0]
  clen := len(a.Credits)

  var wg sync.WaitGroup
  wg.Add(clen)

  m := sync.Mutex{}

  for i := 0; i < clen; i++ {
    go func(wg *sync.WaitGroup, index int) {
      c := a.Credits[index]
      count := 1

      for _, ab := range actors[1:] {
        for _, ac := range ab.Credits {
          if ac.Id == c.Id {
            count += 1
            break
          }
        }
      }

      if count == actorLength {
        m.Lock()
        credits = append(credits, c)
        m.Unlock()
      }
      wg.Done()
    }(&wg, i)
  }

  wg.Wait()
  return credits
}
