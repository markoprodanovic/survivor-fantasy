package db

import (
	"github.com/gocraft/dbr"
	_ "github.com/lib/pq"
)

var Session *dbr.Session

func Initialize() (*dbr.Connection, error) {
	connStr := "user=marko dbname=survivor_fantasy password=survivor host=localhost port=5434 sslmode=disable"
	con, err := dbr.Open("postgres", connStr, nil)
	if err != nil {
		return nil, err
	}

	con.SetMaxOpenConns(25)
	con.SetMaxIdleConns(5)

	return con, nil
}
