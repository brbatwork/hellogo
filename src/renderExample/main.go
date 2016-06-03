package main

import (
  "net/http"
  "fmt"
  "html/template"
  "strings"
  "github.com/unrolled/render"
)
type User struct {
  Name string `json:"name" xml:"name"`
  Email string `json:"email" xml:"email"`
}

func main() {
  r := render.New(render.Options{
    IndentJSON: true,
    IndentXML: true,
    PrefixXML: []byte("<?xml version='1.0' encoding='UTF-8'?>\n"),
    Directory: "/Users/Thunderbird/bbarbour/Dropbox/projects/HelloGo/templates",
    Extensions: []string{".html", ".tpl"},
    Funcs: []template.FuncMap{{"upper": strings.ToUpper}},
    // IsDevelopment: true // recompiles templates on every request
    Layout:     "application",
  })

  http.HandleFunc("/json", func(res http.ResponseWriter, req *http.Request) {
    user := userFromReq(req)
    r.JSON(res, 200, user)
  })

  http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
    user := userFromReq(req)
    r.HTML(res, 200, "index", user)
  })

  http.HandleFunc("/xml", func(res http.ResponseWriter, req *http.Request) {
    user := userFromReq(req)
    r.XML(res, 200, user)
  })

  http.ListenAndServe(":4000", nil)
}

func userFromReq(req *http.Request) *User {
  name := req.URL.Query().Get("name")
  user := &User{
    Name: name,
    Email: fmt.Sprintf("%s@example.com", name),
  }

  return user
}
