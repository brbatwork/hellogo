package main

import (
  "fmt"
  "log"
  "os"
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

  err := db.Update(func(tx *bolt.Tx) error {
    b, err := tx.CreateBucketIfNotExists(bucketName)
    if err != nil {
      return err
    }
    b.Put([]byte("bill"), []byte("ooops"))
    return errors.New("ooops") // Will auto rollback
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
}
