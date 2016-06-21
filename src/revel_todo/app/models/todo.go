package models

import (
  "fmt"
  "time"
)
var DB = []*Todo{}

type Todo struct {
  ID int
  Subject string
  Description string
  CreatedAt time.Time
  UpdatedAt time.Time
}

func init() {
  t := &Todo{Subject: "Buy Milk!", Description: "We need more calcium!"}
  t.Save()
  t = &Todo{Subject: "Learn Go!"}
  t.Save()
}

func (t *Todo) Save() {
  t.CreatedAt = time.Now()
  t.UpdatedAt = t.CreatedAt
  DB = append(DB, t)
  t.ID = len(DB) - 1
}

func (t *Todo) Update() {
  t.UpdatedAt = time.Now()
  DB[t.ID] = t
}

func Get(id int) (*Todo, error) {
  if len(DB) >= id {
    return DB[id], nil
  }
  return &Todo{}, fmt.Errorf("Couldn't find todo with id = %d", id)
}
