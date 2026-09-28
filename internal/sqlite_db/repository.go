package storage

import (
	"context"
	"embed"
	"keep_time/internal/domain"
	"log"
)

const (
	clientsSqlFileName    = "queries/clients.sql"
	projectsSqlFileName   = "queries/projects.sql"
	timeBlocksSqlFileName = "queries/time_blocks.sql"
)

//go:embed queries
var queryFiles embed.FS

type ClientRepo interface {
	GetByID(ctx context.Context, id int) (*domain.Client, error)
	GetByName(ctx context.Context, name string) (*domain.Client, error)
	Safe(ctx context.Context, client *domain.Client) error
	List(ctx context.Context) ([]*domain.Client, error)
}
 
func NewClientRepo() map[string]string {
	queries, err := readQueriesFromFile(queryFiles, clientsSqlFileName)
	if err != nil {
		log.Fatal(err)
	}
	return queries
}
