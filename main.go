package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/dsykes16/pgfib/fibonacci"
	"github.com/dsykes16/pgfib/server"
)

func main() {
	config, err := LoadConfig("./")
	if err != nil {
		log.Fatalf("unable to load config: %s", err)
	}

	db, err := OpenDB(config)
	if err != nil {
		log.Fatalf("unable to connect to postgres: %s", err)
	}

	fib, err := fibonacci.New(db, config.SQLInitPath)
	if err != nil {
		log.Fatalf("unable to create fibonacci generator: %s", err)
	}

	log.Fatal(
		http.ListenAndServe(
			config.ServerAddress,
			server.New(fib),
		),
	)
}

func OpenDB(config Config) (db *sql.DB, err error) {
	db, err = sql.Open(
		"postgres",
		config.ConnectionString(),
	)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return
}
