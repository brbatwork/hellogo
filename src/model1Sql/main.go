package main

import (
  "database/sql"
_  "github.com/lib/pq"
  "time"
  "log"
)

type Todo struct {
  id int
  subject string
  dueDate time.Time
  isComplete bool
}

func main() {
  db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost/playground?sslmode=disable")
  if err != nil {
    log.Fatal(err)
  }

  var max int
  idRow := db.QueryRow("select max(id) from todos")
  err = idRow.Scan(&max)
  if err != nil {
    log.Fatal(err)
  }

  now := time.Now()
  res, err := db.Exec("insert into todos (id, subject, due_date, is_complete) values ($1, $2, $3, $4)", max + 1, "Todo from Go", now, "false")

  if err  != nil {
    log.Fatal(err)
  }

  affected, _ := res.RowsAffected()
  log.Printf("Rows affected %d\n", affected)

  rows, err := db.Query("select * from todos")
  if err != nil {
    log.Fatal(err)
  }

  for rows.Next() {
    todo := Todo{}

    err = rows.Scan(&todo.id, &todo.subject, &todo.dueDate, &todo.isComplete)
    if err != nil {
      log.Fatal(err)
    }
    log.Printf("Id %d: subject is %s\n", todo.id, todo.subject)
  }
}
