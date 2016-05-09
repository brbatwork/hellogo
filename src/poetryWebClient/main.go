package main

import (
    "fmt"
    "encoding/json"
    "net/http"
    "os"
    "sort"
    "poetry"
)

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
  var found bool

  for _, v := range  c.ValidPoems {
    if v == poemName {
      found = true
      break
    }
  }

  if !found {
    http.Error(w, fmt.Sprintf(`{"Error":"Program not found"}`), http.StatusNotFound)
    return
  }

  p, err := poetry.LoadPoem(poemName)

  if err != nil {
    http.Error(w, fmt.Sprintf(`{"Error":"Program not found"}`), http.StatusNotFound)
  } else {
    // fmt.Fprintf(w, "%v\n", p) // as a string
    sort.Sort(p[0])
    pwt := poemWithTitle{poemName, p, p.NumWords(), p.NumThe()}
    enc := json.NewEncoder(w)
    enc.Encode(pwt)
  }
}

func main() {
  configFile, err := os.Open("config")

  if err != nil {
    fmt.Println("Error can't find config file")
    os.Exit(1)
  }

  dec := json.NewDecoder(configFile)
  err = dec.Decode(&c)
  configFile.Close()

  if err != nil {
    fmt.Printf("Error decoding configFile %s", err)
    os.Exit(1)
  }

  fmt.Printf("Route: %s, BindAddress: %s\n", c.Route, c.BindAddress)
  http.HandleFunc(c.Route, poemHandler)
  http.ListenAndServe(c.BindAddress, nil)
}
