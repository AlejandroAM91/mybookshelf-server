package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/AlejandroAM91/mybookshelf-server/pkg/secrets"
)

const (
	defaultHost = "localhost"
	defaultPort = "5432"
	dbName      = "mybookshelf"
)

// DB contains database connection
type DB struct {
	db *sql.DB
}

// Open configures and creates a database connection.
func Open() DB {
	user, err := secrets.GetSecret("MYBOOKSHELF_DB_USER")
	if err != nil {
		panic(err)
	}

	pass, err := secrets.GetSecret("MYBOOKSHELF_DB_PASSWORD")
	if err != nil {
		panic(err)
	}

	host, exists := os.LookupEnv("MYBOOKSHELF_HOST")
	if !exists {
		host = defaultHost
	}

	port, exists := os.LookupEnv("MYBOOKSHELF_PORT")
	if !exists {
		port = defaultPort
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	return DB{db}
}

// Close destroy a database connection.
func (conn *DB) Close() {
	conn.db.Close()
}

// Exec perform request without output data
func (conn *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return conn.db.Exec(query, args...)
}

// Query perform request with output data
func (conn *DB) Query(query string) (*sql.Rows, error) {
	return conn.db.Query(query)
}
