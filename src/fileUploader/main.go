package main

import (
  "net/http"
  "os"
  "io"
  "time"
  "math/rand"
  "fmt"
)

func Upload(res http.ResponseWriter, req *http.Request) {
  file, handler, err := req.FormFile("my_file")
  if err != nil {
    fmt.Fprint(res, err)
    return
  }

  dirname := uploadDirectoryName()
  os.MkdirAll(dirname, 0777)
  filename := fmt.Sprintf("%s/%s", dirname, handler.Filename)
  outfile, err := os.Create(filename)
  if err != nil {
    fmt.Fprint(res, err)
    return
  }

  defer outfile.Close()
  _, err = io.Copy(outfile, file)
  if err != nil {
    fmt.Fprint(res, err)
    return
  }

  fmt.Fprintln(res, filename)
}

var uploadDirectoryName = func() string {
  rand.Seed(time.Now().Unix())
  dirname := fmt.Sprintf("/tmp/go_file_uploads/%d", rand.Int())
  return dirname
}

func main() {
  http.HandleFunc("/upload", Upload)
  http.Handle("/", http.FileServer(http.Dir("public")))
  http.ListenAndServe(":4000", nil)
}
