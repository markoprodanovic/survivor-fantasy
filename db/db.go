package db

import (
	"github.com/gocraft/dbr"
	_ "github.com/mattn/go-sqlite3"
)

var Session *dbr.Session

func Initialize() (*dbr.Connection, error) {
	connStr := "file:survivor_fantasy.db?cache=shared&mode=rwc"
	con, err := dbr.Open("sqlite3", connStr, nil)
	if err != nil {
		return nil, err
	}

	con.SetMaxOpenConns(25)
	con.SetMaxIdleConns(5)

	return con, nil
}
