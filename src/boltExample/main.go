package main

import (
  "fmt"
  "log"
  "os"
  "errors"
  "encoding/json"
  "github.com/boltdb/bolt"
)

var (
  db_name = []byte("mystorage")
  db *bolt.DB
  bucketName = []byte("myBucket")
)

func init() {
  var err error
  db, err = bolt.Open("bolt.db", 0644, nil)
  if err != nil {
    log.Fatal(err)
  }
}

type Person struct {
  Name string
  Email string
}

func main() {
  defer db.Close()
  defer os.Remove(db.Path())

  err := db.Update(func(tx *bolt.Tx) error {
    b, err := tx.CreateBucketIfNotExists(bucketName)
    if err != nil {
      return err
    }
    return b.Put([]byte("bill"), []byte("bill jones"))
  })

  if err != nil {
    log.Fatal(err)
  }

  err = db.Update(func(tx *bolt.Tx) error {
    b, err := tx.CreateBucketIfNotExists(bucketName)
    if err != nil {
      return err
    }
    b.Put([]byte("bill"), []byte("ooops"))
    return errors.New("ooops") // Will auto rollback
  })

  err = db.Update(func(tx *bolt.Tx) error {
    b, err := tx.CreateBucketIfNotExists(bucketName)
    if err != nil {
      return err
    }
    p := Person{"bill jones", "bill@example.com"}
    by, _ := json.Marshal(p)
    return b.Put([]byte("bill"), by)
  })

  err = db.View(func(tx *bolt.Tx) error {
    b := tx.Bucket(bucketName)
    m := b.Get([]byte("bill"))
    fmt.Printf("m: %s\n", m)
    return nil
  })

  if err != nil {
    log.Fatal(err)
  }

  err = db.View(func(tx *bolt.Tx) error {
    b := tx.Bucket(bucketName)
    by := b.Get([]byte("bill"))
    p := Person{}
    json.Unmarshal(by, &p)

    fmt.Printf("m: %s\n", p)
    return nil
  })

}
