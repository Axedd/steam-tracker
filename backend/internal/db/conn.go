package db

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

// Connect opens a *sql.DB pool from the given Postgres URL.
// It verifies the connection and configures pooling.
func Connect(databaseURL string) (*sql.DB, error) {
	dbConn, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	// Tune these to match your Postgres limits and load:
	dbConn.SetMaxOpenConns(25)
	dbConn.SetMaxIdleConns(25)
	dbConn.SetConnMaxLifetime(5 * time.Minute)
	// Verify connection on startup:
	if err := dbConn.Ping(); err != nil {
		dbConn.Close()
		return nil, err
	}
	return dbConn, nil
}
