package db

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq"
)

func NewPostgresDB(connectionString string) *sql.DB {
    db, err := sql.Open("postgres", connectionString)
    if err != nil {
        log.Fatalf("could not connect to the database: %v", err)
    }
    if err := db.Ping(); err != nil {
        log.Fatalf("could not ping the database: %v", err)
    }
    return db
}
