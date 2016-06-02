package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

func IndexUsers(res http.ResponseWriter, req *http.Request) {
  fmt.Fprint(res, "Users Index")
}

func ShowUser(res http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(res, "Users Show: %s", req.URL.Path)
}

func CreateUser(res http.ResponseWriter, req *http.Request) {
  body, _ := ioutil.ReadAll(req.Body)
  fmt.Fprintf(res, "Users Create: %s", body)
}

func IndexPosts(res http.ResponseWriter, req *http.Request) {
  fmt.Fprint(res, "Posts Index")
}


func NewMux() http.Handler {
  mux := http.NewServeMux()
  users := func(res http.ResponseWriter, req *http.Request) {
    if req.URL.Path == "" {
      switch req.Method {
      case "GET":
        IndexUsers(res, req)
      case "POST":
        CreateUser(res, req)
      }
    } else {
      ShowUser(res, req)
    }
  }

  // Handle {id} at the end of the url    vvvvvvv is stripped
  mux.Handle("/users/", http.StripPrefix("/users/", http.HandlerFunc(users)))
  mux.HandleFunc("/posts", IndexPosts)
  return mux
}

func main() {
  http.ListenAndServe(":3000", NewMux())
}
