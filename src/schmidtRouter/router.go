package routes

import (
  "github.com/julienschmidt/httprouter"
  "net/http"
  "github.com/gorilla/pat"
  "fmt"
)

func Router() *httprouter.Router {
  h := func(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
    fmt.Fprintf(res, "URL: %s\n", req.URL)
    fmt.Fprintf(res, "ID: %s\n", p.ByName("id"))
    fmt.Fprintf(res, "PATH: %s\n", p.ByName("path"))

  }
  r := httprouter.New()
  r.GET("/a", h)
  r.GET("/b", h)
  r.GET("/c/:id", h)
  r.GET("/d/*path", h)
  return r
}

func Pat() *pat.Router {
  h := func(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(res, "URL: %s\n", req.URL)
    fmt.Fprintf(res, "ID: %s\n", req.URL.Query().Get(":id"))
    fmt.Fprintf(res, "PATH: %s\n", req.URL.Query().Get(":path"))

  }
  r := pat.New()
  r.Get("/a", h)
  r.Get("/b", h)
  r.Get("/c/:id", h)
  r.Get("/d/*path", h)
  return r
}
