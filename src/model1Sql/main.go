package main

import (
  "database/sql"
_  "github.com/lib/pq"
  "time"
  "log"
)

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
  var subject string

  rows, err := db.Query("select subject from todos")
  if err != nil {
    log.Fatal(err)
  }

  for rows.Next() {
    err = rows.Scan(&subject)
    if err != nil {
      log.Fatal(err)
    }
    log.Printf("Subject is %s\n", subject)
  }
}
