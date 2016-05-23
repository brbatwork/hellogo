package main

import (
  "github.com/negroni"
  "io/ioutil"
  "net/http"
  "fmt"
  "log"
  "github.com/gorilla/pat"
  "github.com/russross/blackfriday"
  "github.com/microcosm-cc/bluemonday"
)
func main() {
  p := pat.New()
  p.Post("/markdown", func(res http.ResponseWriter, req *http.Request) {
    md, err := ioutil.ReadAll(req.Body)
    if err != nil {
      log.Println("There was an error reading the body")
      res.WriteHeader(http.StatusBadRequest)
      fmt.Fprint(res, err)
      return
    }

    unsafe := blackfriday.MarkdownCommon(md)
    html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
    log.Println("Sending down html in the response");
    fmt.Fprint(res, string(html))
    })
  n := negroni.Classic()
  n.UseHandler(p)
  n.Run(":7070")
}
