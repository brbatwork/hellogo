package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/pat"
  "github.com/gorilla/mux"
  "io/ioutil"
)

func IndexUsers(res http.ResponseWriter, req *http.Request) {
  fmt.Fprint(res, "Users Index")
}

func AIndexUsers(res http.ResponseWriter, req *http.Request) {
  fmt.Fprint(res, "Users Index A")
}

func BIndexUsers(res http.ResponseWriter, req *http.Request) {
  fmt.Fprint(res, "Users Index B")
}

func ShowUser(res http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(res, "Users Show: %s", req.URL.Query().Get(":id"))
}

func MuxShowUser(res http.ResponseWriter, req *http.Request) {
  vars := mux.Vars(req)
  fmt.Fprintf(res, "Users Show: %s", vars["id"])
}
func CreateUser(res http.ResponseWriter, req *http.Request) {
  body, _ := ioutil.ReadAll(req.Body)
  fmt.Fprintf(res, "Users Create: %s", body)
}

func IndexPosts(res http.ResponseWriter, req *http.Request) {
  fmt.Fprint(res, "Posts Index")
}

// pat router more simpler
func NewPatMux() http.Handler {
  pat := pat.New()
  pat.Get("/posts", IndexPosts)
  pat.Get("/users/{id}", ShowUser)
  pat.Get("/users", IndexUsers)
  pat.Post("/users/", CreateUser)
  return pat
}

func NewMux() http.Handler {
  r := mux.NewRouter()
  r.HandleFunc("/posts", IndexPosts)
  r.HandleFunc("/users/{id:[0-9]+}", MuxShowUser)
  r.HandleFunc("/users", IndexUsers).Methods("GET")
  r.HandleFunc("/users", AIndexUsers).Methods("GET").Host("a.example.com")
  r.HandleFunc("/users", BIndexUsers).Methods("GET").Host("b.example.com")
  r.HandleFunc("/users", CreateUser).Methods("POST")
  return r
}

func main() {
  http.ListenAndServe(":3000", NewMux())
}
