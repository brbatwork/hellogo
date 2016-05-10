package main

import (
    "fmt"
    "encoding/json"
    "net/http"
    "os"
    "sync"
    "log"
    "flag"
    "time"
    "poetry"
)

var cacheMutex sync.Mutex
var cache map[string]poetry.Poem

type config struct {
  Route string
  BindAddress string `json:"addr"`
  ValidPoems []string `json:"valid"`
}

var c config // Global so poemhandler gets it

type poemWithTitle struct {
  Title string
  Body poetry.Poem
  WordCount int
  TheCount int
}

func poemHandler(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  if len(r.Form["name"]) <= 0 {
    http.Error(w, "name is required", http.StatusBadRequest)
    return
  }

  poemName := r.Form["name"][0]
  p, ok := cache[poemName]

  if !ok {
    http.Error(w, fmt.Sprintf(`{"Error":"Poem %s not found"}`, poemName), http.StatusNotFound)
    return
  }

  log.Printf("User has requested poem %s", poemName)
  // sort.Sort(p[0])
  pwt := poemWithTitle{poemName, p, p.NumWords(), p.NumThe()}
  enc := json.NewEncoder(w)
  enc.Encode(pwt)
}

func main() {

  configFileName := flag.String("conf", "config", "Name of configuration file")
  flag.Parse()
  configFile, err := os.Open(*configFileName)

  if err != nil {
    log.Fatalf("Error can't find config file")
  }

  dec := json.NewDecoder(configFile)
  err = dec.Decode(&c)
  configFile.Close()

  if err != nil {
    log.Fatalf("Error decoding configFile %s", err)
  }

  cache = make(map[string]poetry.Poem)
  var wg sync.WaitGroup

  startTime := time.Now()

  // Load the in memory cache
  for _, name := range c.ValidPoems {
    //async load the poems
    wg.Add(1)
    go func(n string) {
      fmt.Printf("Loading poem %s into cache\n", n)
      cacheMutex.Lock()
      cache[n], err = poetry.LoadPoem(n)
      cacheMutex.Unlock()
      if err != nil {
        log.Fatalf("Error loading poem %s", n)
      }
      wg.Done()
    }(name)

  }

  wg.Wait() // Wait for all the poems to be loaded
  elapsedTime := time.Since(startTime)

  log.Printf("Statup complete in %s Route: %s, BindAddress: %s\n", elapsedTime, c.Route, c.BindAddress)
  http.HandleFunc(c.Route, poemHandler)
  http.ListenAndServe(c.BindAddress, nil)
}
