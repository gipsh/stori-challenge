package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/mattn/go-sqlite3"
	migrate "github.com/rubenv/sql-migrate"
)

var once sync.Once
var connection *sql.DB

func Connection() (*sql.DB, error) {
	var err error
	once.Do(func() {
		connection, err = initialize()
		if err != nil {
			return
		}
	})
	return connection, err
}

// Close closes the connection to the database
func Close() {
	connection.Close()
}

func initialize() (*sql.DB, error) {
	return sql.Open(os.Getenv("DB_DRIVER"), fmt.Sprintf("file:%s", os.Getenv("SQLITE_FILE")))
}

func Migrate(db *sql.DB) error {
	migrations := &migrate.FileMigrationSource{
		Dir: "internal/db/migrations",
	}

	n, err := migrate.Exec(db, os.Getenv("DB_DRIVER"), migrations, migrate.Up)
	if err != nil {
		return err
	}
	log.Println("Applied ", n, " migrations")

	return err
}
