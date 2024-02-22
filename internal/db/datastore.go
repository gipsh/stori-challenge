package db

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/mattn/go-sqlite3"
	migrate "github.com/rubenv/sql-migrate"
)

var once sync.Once
var connection *sql.DB

func Connection() *sql.DB {
	once.Do(func() {
		connection = initialize()
	})

	return connection
}

// Close closes the connection to the database
func Close() {
	connection.Close()
}

func initialize() *sql.DB {
	db, err := sql.Open(os.Getenv("DB_DRIVER"), fmt.Sprintf("file:%s", os.Getenv("SQLITE_FILE")))
	if err != nil {
		panic(err)
	}

	return db
}

func Migrate(db *sql.DB) error {
	migrations := &migrate.FileMigrationSource{
		Dir: "internal/db/migrations",
	}

	n, err := migrate.Exec(db, os.Getenv("DB_DRIVER"), migrations, migrate.Up)
	if err != nil {
		return err
	}
	fmt.Println("Applied ", n, " migrations")

	return err
}
