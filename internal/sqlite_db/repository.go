package storage

import (
	"embed"
	"log"
	"context"
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
	GetByName(ctx context.Context, name str)
	List(ctx context.Context)
	Update()
}

func NewClientRepo() map[string] string {
	queries, err := readQueriesFromFile(queryFiles, clientsSqlFileName)
	if err != nil {
		log.Fatal(err)
	}
	return queries
}
