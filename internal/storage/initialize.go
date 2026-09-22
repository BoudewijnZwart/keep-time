package storage

import (
	"bufio"
	"embed"
	"io/fs"
	"log"
	"strings"
)

var queryStartDelimiter string = "-- name:"

//go:embed schema.sql
var Schema string

//go:embed queries
var queryFile embed.FS

func readQueriesFromFile(fsys fs.FS, fileName string) string {
	var query strings.Builder
	file, err := fsys.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if isQueryName(line) {
			if 

		}
		query.WriteString(line)
		query.WriteByte("\n")
	}
}

func isQueryName(s string) bool {
	return strings.Contains(s, queryStartDelimiter)
}
