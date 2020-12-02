package mysql

import (
	"database/sql"

	// Imported for mysql connection
	_ "github.com/go-sql-driver/mysql"
)

// NewConnection creates a new connection to mySQL DB
func NewConnection(connectionURL string) (*sql.DB, error) {
	return sql.Open("mysql", connectionURL)
}
