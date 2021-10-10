package main

import (
	"database/sql"
	_ "embed"
	"time"
)

var (
	//go:embed schema.sql
	schemaSQL string

	//go:embed insert.sql
	insertSQL string
)

func createTables(db *sql.DB) error {
	_, err := db.Exec(schemaSQL)
	return err
}

func insertLog(db *sql.DB, time time.Time, message string) error {
	_, err := db.Exec(insertSQL, time, message)
	return err
}
