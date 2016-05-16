package main

import (
  "io/ioutil"
  "net/http"
  "net/url"
  "github.com/DavidHuie/quartz/go/quartz"
  "encoding/json"
)

type Poster struct {}

type Args struct {
  Params map[string]string
}

type Response struct {
  Json map[string]interface{}
  Body string
  Values url.Values
  Args Args
}

func (p *Poster) Echo(args Args, response *Response) error {
  u := "http://quiet-waters-1228.herokuapp.com/echo.json"

  response.Args = args
  params := url.Values{}

  for name, value := range args.Params {
    params.Add(name, value)
  }

  response.Values = params
  resp, err := http.PostForm(u, params)
  if err != nil {
    return err
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return err
  }

  response.Body = string(body)
  json.Unmarshal(body, &response.Json)
  return nil
}

func main() {
  poster := &Poster{}
  quartz.RegisterName("my_poster", poster)
  quartz.Start()
}
