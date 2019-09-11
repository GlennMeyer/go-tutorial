package main

import (
	"github.com/go-pg/pg"
	"time"
)

func pgOptions() *pg.Options {
	return &pg.Options{
		User:     "postgres",
		Password: "",
		Database: "postgres",
		Addr: 		"db:5432",

		MaxRetries:      1,
		MinRetryBackoff: -1,

		DialTimeout:  30 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,

		PoolSize:           10,
		MaxConnAge:         10 * time.Second,
		PoolTimeout:        30 * time.Second,
		IdleTimeout:        10 * time.Second,
		IdleCheckFrequency: 100 * time.Millisecond,
	}
}

func pgConnect() *pg.DB {
	return pg.Connect(pgOptions())
}
