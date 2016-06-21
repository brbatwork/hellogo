package controllers

import (
  "github.com/revel/revel"
  "revel_todo/app/models"
)

type Todos struct {
  *revel.Controller
}

func (c Todos) Index() revel.Result {
  ts := models.DB
  return c.Render(ts)
}

func (c Todos) Show(id int) revel.Result {
  t, err := models.Get(id)

  if err != nil {
    return c.RenderError(err)
  }

  return c.Render(t)
}
