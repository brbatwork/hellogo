package main

import (
  "expvar"
  "fmt"
  "net/http"
)

func main() {
  views := expvar.NewInt("views")
  pages := expvar.NewMap("pages")

  http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
    u := req.URL.String()
    pages.Add(u, 1)
    views.Add(1)
    tp := pages.Get(u)
    fmt.Fprintf(res, "Total views for this page are %s", tp)
  })
  http.ListenAndServe(":7070", nil)
}
