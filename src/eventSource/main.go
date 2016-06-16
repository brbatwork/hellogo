package main

import (
  "github.com/gorilla/pat"
  "github.com/negroni"
  "net/http"
)

func MessagesHandler(res http.ResponseWriter, req *http.Request) {

}

func main() {
  p := pat.New()
  p.Post("/messages", MessagesHandler)

  n := negroni.Classic()
  n.UseHandler(p)
  http.ListenAndServe(":3000", n)
}
