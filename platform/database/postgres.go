package database

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"api-fiber/pkg/utils"

	"github.com/jmoiron/sqlx"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib" // load pgx driver for PostgreSQL
)

// PostgreSQLConnection func for connection to PostgreSQL database.
func PostgreSQLConnection() (*sqlx.DB, error) {
	// Define database connection settings.
	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	// Build PostgreSQL connection URL.
	postgresConnURL, err := utils.ConnectionURLBuilder("postgres")
	if err != nil {
		return nil, err
	}

	// Define database connection for PostgreSQL.
	db, err := sqlx.Connect("pgx", postgresConnURL)
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	// Set database connection settings:
	// 	- SetMaxOpenConns: the default is 0 (unlimited)
	// 	- SetMaxIdleConns: defaultMaxIdleConns = 2
	// 	- SetConnMaxLifetime: 0, connections are reused forever
	db.SetMaxOpenConns(maxConn)
	db.SetMaxIdleConns(maxIdleConn)
	db.SetConnMaxLifetime(time.Duration(maxLifetimeConn))

	// Try to ping database.
	if err := db.Ping(); err != nil {
		defer db.Close() // close database connection
		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}

	return db, nil
}

func Migrate() {
	// Build PostgreSQL connection URL.
	postgresConnURL, err := utils.ConnectionURLBuilder("postgres")
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.New(
		"file://db/migrations",
		postgresConnURL)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}
