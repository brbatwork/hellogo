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
  DueDate time.Time `db:"due_date"`
  IsComplete bool   `db:"is_complete"`
}

func main() {
  db, err := sqlx.Open("postgres", "postgres://postgres:postgres@localhost/playground?sslmode=disable")
  if err != nil {
    log.Fatal(err)
  }

  todos := []Todo{}
  db.Select(&todos, "select * from todos")

  for _, todo := range todos {
    log.Printf("Id %d: subject is %s\n", todo.Id, todo.Subject)
  }
}
