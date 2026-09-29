package storage

import (
	"bufio"
	"database/sql"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	_ "embed"

	_ "modernc.org/sqlite"
)

var queryStartDelimiter string = "-- name:"

//go:embed schema.sql
var Schema string

// OpenDB opens a connection to the sqlite database at path, creating the
// file (and any missing parent directories) if it does not already exist.
func OpenDB(path string) (*sql.DB, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, fmt.Errorf("creating directory for database: %w", err)
	}

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("opening database: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("connecting to database: %w", err)
	}

	return db, nil
}

// ApplySchema creates the tables defined in Schema if they do not already
// exist.
func ApplySchema(db *sql.DB) error {
	if _, err := db.Exec(Schema); err != nil {
		return fmt.Errorf("applying schema: %w", err)
	}
	return nil
}

// InitDB opens the sqlite database at path, creating the file if needed,
// and ensures the schema is applied. This is the entry point CLI commands
// should use to get a ready-to-use database connection.
func InitDB(path string) (*sql.DB, error) {
	db, err := OpenDB(path)
	if err != nil {
		return nil, err
	}

	if err := ApplySchema(db); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

// Read queries from a sql file
func readQueriesFromFile(fsys fs.FS, fileName string) (map[string] string, error) {
	var queryBuilder strings.Builder
	var queryName string
	queries := make(map[string] string)

	file, err := fsys.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if isQueryName(line) {
			if queryName != "" {
				queries[queryName] = queryBuilder.String()
				queryBuilder.Reset()
			}
			queryName = getQueryName(line)
			continue
		}
		queryBuilder.WriteString(line)
		queryBuilder.WriteByte('\n')
	}
	if queryName != "" {
		queries[queryName] = queryBuilder.String()
	}
	return queries, scanner.Err()
}

// Check if a line is a query name line
func isQueryName(s string) bool {
	return strings.Contains(s, queryStartDelimiter)
}

// Get the name from a query line in a sql file
func getQueryName(s string) string {
	s = strings.Replace(s, queryStartDelimiter,"",1)
	return strings.TrimSpace(s)
}
