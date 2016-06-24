package routes

import (
  "github.com/julienschmidt/httprouter"
  "net/http"
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
