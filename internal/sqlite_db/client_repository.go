package storage

import (
	"context"
	"embed"
	"keep_time/internal/domain"
	"log"
)

const (
	clientsSqlFileName = "queries/clients.sql"
)

//go:embed queries
var queryFiles embed.FS

type ClientRepo struct{}

func (c *ClientRepo) GetByID(ctx context.Context, id int) (*domain.Client, error)

func (c *ClientRepo) GetByName(ctx context.Context, name string) (*domain.Client, error)

func (c *ClientRepo) Safe(ctx context.Context, client *domain.Client) error

func (c *ClientRepo) List(ctx context.Context) ([]*domain.Client, error)

func NewClientRepo() map[string]string {
	queries, err := readQueriesFromFile(queryFiles, clientsSqlFileName)
	if err != nil {
		log.Fatal(err)
	}
	return queries
}
