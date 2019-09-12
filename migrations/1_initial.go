package main

import (
  "fmt"

  "github.com/go-pg/migrations"
)

func init() {
  migrations.MustRegisterTx(func(db migrations.DB) error {
    fmt.Println("creating table articles...")
    _, err := db.Exec(`CREATE TABLE articles(
      id serial PRIMARY KEY,
      title varchar(255) NOT NULL,
      content text NOT NULL
    )`)
    return err
  }, func(db migrations.DB) error {
    fmt.Println("dropping table articles...")
    _, err := db.Exec(`DROP TABLE articles`)
    return err
  })
}