package main

import (
	"database/sql"
	_ "embed"
	"fmt"
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
	ts := time.Format("2006-01-02 15:04:05") // Format time for SQL
	sql := fmt.Sprintf(insertSQL, ts, message)
	_, err := db.Exec(sql)
	return err
}
