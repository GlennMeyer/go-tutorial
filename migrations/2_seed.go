package main

import (
  "fmt"

  "github.com/go-pg/migrations"
)

func init() {
  migrations.MustRegisterTx(func(db migrations.DB) error {
    fmt.Println("seeding articles...")
    _, err := db.Exec(`INSERT INTO articles (title, content) VALUES
    ('Article 1', 'Article 1 body'),
    ('Article 2', 'Article 2 body');`)
    return err
  }, func(db migrations.DB) error {
    fmt.Println("truncating articles...")
    _, err := db.Exec(`TRUNCATE articles`)
    return err
  })
}