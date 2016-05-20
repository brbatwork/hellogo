package main_test

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "net/http/httptest"
  "fmt"
  "net/http"
  main "imdbSearcher"
)

func Test_FetchActor_WithResults(t *testing.T) {
  a := assert.New(t)
  body := `{
    "page":1,
    "results":[
      {
        "id":287,
        "name":"Brad Pitt",
        "popularity":12,
        "profile_path":"/brad.jpg"
      }
    ],
    "total_pages":1,
    "total_results":1
  }`

  FakeServer(body, func() {
    actor, err := main.FetchActor("Brad Pitt")
    a.NoError(err)
    a.Equal("Brad Pitt", actor.Name)
  })
}

func Test_FetchActor_WithNoResults(t *testing.T) {
  a := assert.New(t)
  body := `{
    "page":1,
    "results":[],
    "total_pages":1,
    "total_results":0
  }`

  FakeServer(body, func() {
    _, err := main.FetchActor("Brad Pitt")
    a.Equal("There are no search results for Brad Pitt!", err.Error())
  })
}

func FakeServer(b string, f func()) {
  root := main.ApiRoot // keep a copy so we can put back
  ts := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
    fmt.Fprint(resp, b)
    }))
  defer ts.Close()
  main.ApiRoot = ts.URL
  f() // Call the function to test
  main.ApiRoot = root // reset our ApiRoot back to the real service
}
