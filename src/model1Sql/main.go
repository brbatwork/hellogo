package main

import (
  "github.com/jmoiron/sqlx"
_  "github.com/lib/pq"
  "time"
  "log"
)

type Todo struct {
  Id int
  Subject string
  DueDate time.Time
  IsComplete bool
}

func main() {
  db, err := sqlx.Open("postgres", "postgres://postgres:postgres@localhost/playground?sslmode=disable")
  if err != nil {
    log.Fatal(err)
  }

  todos := []Todo{}
  db.Select(&todos, "select id, subject, due_date AS DueDate, is_complete AS IsComplete from todos")

  for _, todo := range todos {
    log.Printf("Id %d: subject is %s\n", todo.Id, todo.Subject)
  }
}
