package controllers

import (
  "github.com/revel/revel"
  "fmt"
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

func (c Todos) New() revel.Result {
  return c.Render(&models.Todo{})
}

func (c Todos) Edit(id int) revel.Result {
  t, err := models.Get(id)
  if err != nil {
    return c.RenderError(err)
  }
  return c.Render(t)
}

func (c Todos) Update(id int, t *models.Todo) revel.Result {
  c.Validation.Required(t.Subject)
  if c.Validation.HasErrors() {
    c.RenderArgs["t"] = t
    return c.RenderTemplate("Todos/Edit.html")
  }

  et, err := models.Get(id)
  if err != nil {
    return c.RenderError(err)
  }

  et.Subject = t.Subject
  et.Description = t.Description
  et.Update()
  c.Flash.Success("Your todo was successfully updated")
  return c.Redirect(fmt.Sprintf("/todos/%d", t.ID))
}
