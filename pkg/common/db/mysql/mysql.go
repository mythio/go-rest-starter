package mysql

import (
	"database/sql"

	// Imported for mysql connection
	_ "github.com/go-sql-driver/mysql"
)

// NewConnection creates a new connection to mySQL DB
func NewConnection(connectionURL string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connectionURL)
	if err != nil {
		return nil, err
	}

	return db, db.Ping()
}
