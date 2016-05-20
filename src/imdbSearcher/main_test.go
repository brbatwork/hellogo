package main_test

import (
  "bytes"
  "testing"
  "github.com/stretchr/testify/assert"
  main "imdbSearcher"
)

func Test_E2E(t *testing.T) {
  a := assert.New(t)
  r := bytes.NewBuffer([]byte("Brad Pitt\nJennifer Aniston\nn\n"))
  w := &bytes.Buffer{}

  main.Run(r, w)
  res := w.String()
  a.Contains(res, "You selected the following 2 actors")
  a.Contains(res, "Jennifer Aniston")
  a.Contains(res, "Brad Pitt")
  a.Contains(res, "Friends")

}
