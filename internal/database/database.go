package database

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var Db *sql.DB

func InitDB() {
	// PostgreSQL connection string. Modify this as per your environment.
	connStr := "user=postgres password=postgres host=localhost port=5432 dbname=go_movies sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	Db = db
}

func CloseDB() error {
	return Db.Close()
}

func Migrate() {
	if err := Db.Ping(); err != nil {
		log.Fatal(err)
	}
	driver, err := postgres.WithInstance(Db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Adjust path to your migrations directory
	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/database/migrations",
		"postgres", // Database name
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
