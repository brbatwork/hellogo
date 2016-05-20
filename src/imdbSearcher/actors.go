package main

import (
  "os"
  "fmt"
  "net/url"
  "encoding/json"
  "net/http"
)

type Actor struct {
  Popularity  float64 `json:"popularity"`
  Id          int     `json:"id"`
  ProfilePath string  `json:"profile_path"`
  Name        string  `json:"name"`
  Credits     []Credit
}

type ActorSearchResults struct {
  Page         int     `json:"page"`
  Results      []Actor `json:results"`
  TotalPages   int     `json:"total_pages"`
  TotalResults int     `json:"total_results"`
}

var ApiRoot = "https://api.themoviedb.org/3"
var ApiKey = ""

func init() {
  ApiKey = os.Getenv("TMDB_KEY")
}

func FetchActor(name string) (Actor, error) {
  u := fmt.Sprintf("%s/search/person?api_key=%s&query=%s", ApiRoot, ApiKey, url.QueryEscape(name))
  results := ActorSearchResults{}
  res, err := http.Get(u)

  a := Actor{}
  if err != nil {
    return a, err
  }

  err = json.NewDecoder(res.Body).Decode(&results)
  if err != nil {
    return a, err
  }

  if results.TotalResults == 0 {
    return a, fmt.Errorf("There are no search results for %s!", name)
  }

  a = results.Results[0]
  err = FetchCredits(&a)

  return a, nil
}
