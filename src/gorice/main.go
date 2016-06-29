package main

import (
  "github.com/GeertJohan/go.rice"
  "log"
  "html/template"
  "net/http"
)

func main() {
  tmpls, err := rice.FindBox("templates")
  if err != nil {
    log.Fatal(err)
  }
  http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
    t, err := tmpls.String("index.html")

    if err != nil {
      log.Fatal(err)
    }

    tmpl, err := template.New("Index").Parse(t)
    if err != nil {
      log.Fatal(err)
    }
    tmpl.Execute(res, nil)
  })

  assets := rice.MustFindBox("assets")
  assetHandler := http.StripPrefix("/assets/", http.FileServer(assets.HTTPBox()))
  http.Handle("/assets/", assetHandler)
  http.ListenAndServe(":3000", nil)
}
