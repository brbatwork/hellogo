package main_test

import (
  "testing"
  "net/http"
  "net/http/httptest"
  "github.com/stretchr/testify/assert"
  main "webHelloWorld"
)

func TestHomePageHandler(t * testing.T) {
  res := httptest.NewRecorder()
  req, _ := http.NewRequest("GET", "/", nil)

  main.LocalPageHandler(res,req)

  if res.Code != 200 {
    t.Fatal("Expected status of 200 but got %d", res.Code)
  }

  assert.Equal(res.Code, 200)
  assert.Equal(res.Body.String(), "Hello, World!")

}
