// Package db provides a database connection pool.
package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Conn *sql.DB
)

func _ping() error {
	return Conn.Ping()
}
