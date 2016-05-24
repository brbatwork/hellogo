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

  tx := db.MustBegin() // Begin transaction
  var currId int
  currIdRow := db.QueryRow("select max(id) from todos")
  err = currIdRow.Scan(&currId)
  now := time.Now()

  t := Todo {
    Id: currId + 1,
    Subject: "Mow Lawn",
    DueDate: now,
    IsComplete: false,
  }

  tx.Exec("insert into todos (id, subject, due_date, is_complete) values ($1, $2, $3, $4)", t.Id, t.Subject, t.DueDate, t.IsComplete)
  tx.Commit() //End transaction

  todos := []Todo{}
  db.Select(&todos, "select * from todos")

  for _, todo := range todos {
    log.Printf("Id %d: subject is %s\n", todo.Id, todo.Subject)
  }
}
