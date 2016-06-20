package main

import (
  "fmt"
  "net/http"
  "github.com/codegangsta/negroni"
  "github.com/gorilla/pat"
)

func main() {
  p := pat.New()
  p.Get("/", func(resp http.ResponseWriter, req *http.Request) {
    fmt.Fprint(resp, template)
  })

  n := negroni.Classic()
  n.UseHandler(p)
  n.Run(":4000")
}

var template = `
<html>
  <head>
    <title>GopherJS</title>
  </head>
  <body>
    <div id="main">
      GopherJS
    </div>
    <script src="./jquery-3.0.0.min.js"></script>
    <script src="./app.js"></script>
  </body>
</html>
`
