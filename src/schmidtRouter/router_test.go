package routes

import (
  "io/ioutil"
  "net/http"
  "net/http/httptest"
  "testing"
  "github.com/stretchr/testify/assert"
)

func Test_Router(t *testing.T) {
  a := assert.New(t)

  ts := httptest.NewServer(Router())
  defer ts.Close()

  routes := []string{"/a", "/b", "/c/1", "/c/2", "/d/1/2/3/4/5"}

  for _, r := range routes {
    res, err := http.Get(ts.URL + r)
    a.NoError(err)

    body, _ := ioutil.ReadAll(res.Body)
    a.Contains(string(body), "URL: " + r)
  }
}

func Test_Router_Params(t *testing.T) {
  a := assert.New(t)
  ts := httptest.NewServer(Router())
  defer ts.Close()
  routes := []string{"1", "2", "3"}

  for _, r := range routes {
    res, err := http.Get(ts.URL + "/c/" + r)
    a.NoError(err)

    body, _ := ioutil.ReadAll(res.Body)
    a.Contains(string(body), "ID: " + r)
  }
}

func Test_Router_Params_Path(t *testing.T) {
  a := assert.New(t)
  ts := httptest.NewServer(Router())
  defer ts.Close()
  routes := []string{"/1/2/3", "/4/5/6", "/7/8/9"}

  for _, r := range routes {
    res, err := http.Get(ts.URL + "/d" + r)
    a.NoError(err)

    body, _ := ioutil.ReadAll(res.Body)
    a.Contains(string(body), "PATH: " + r)
  }
}
