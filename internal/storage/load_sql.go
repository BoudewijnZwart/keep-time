package storage

import (
	"bufio"
	"io/fs"
	"strings"
	_ "embed"
)

var queryStartDelimiter string = "-- name:"

//go:embed schema.sql
var Schema string

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
