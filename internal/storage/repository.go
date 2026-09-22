package storage

import (
	"embed"
	"log"
)

const (
	clientsSqlFileName = "queries/clients.sql"
	projectsSqlFileName = "queries/projects.sql"
	timeBlocksSqlFileName = "queries/time_blocks.sql"
)


//go:embed queries
var queryFiles embed.FS

type ClientRepo interface {
	GetByID(ctx context.Context, id int)
	List(ctx context.Context)
}

func NewClientRepo() map[string] string {
	queries, err := readQueriesFromFile(queryFiles, clientsSqlFileName)
	if err != nil {
		log.Fatal(err)
	}
	return queries
}
