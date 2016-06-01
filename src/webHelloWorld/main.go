package main

import (
  "net/http"
  "encoding/json"
  "fmt"
  "time"
)

type User struct {
  FirstName string
  LastName  string
  Email     string
  CreatedAt time.Time
}
type HomePageHandler struct{}

func (h *HomePageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  user := new(User)
  json.NewDecoder(r.Body).Decode(user)
  user.CreatedAt = time.Now()

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)

  data, _ := json.Marshal(user)
  w.Write(data)

}

func LocalPageHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Hello, World!")
}

func main() {
  // http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  //   fmt.Fprintln(w, "Hello World!")
  // })
  http.Handle("/", &HomePageHandler{})


  http.ListenAndServe(":3000", nil)
}
