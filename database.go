package main

import (
	"fmt"
	"github.com/go-pg/pg"
	"time"
)

func initializeDatabase() {
	db := pg.Connect(&pg.Options{
    User:     "postgres",
    Password: "",
		Database: "postgres",
		Addr: 		"db:5432",
	})
	var n int
	_, err := db.QueryOne(pg.Scan(&n), "SELECT 1;")
	if err != nil {
		fmt.Println("Cannot initialize database: ", err)
		time.Sleep(1 * time.Second)
		initializeDatabase()
	} else {
		fmt.Println("Database initialized.", n)
	}
}