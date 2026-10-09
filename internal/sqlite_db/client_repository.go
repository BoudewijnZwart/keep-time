package storage

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"keep_time/internal/domain"
	"keep_time/internal/service"
	"log"
)

const (
	clientsSqlFileName = "queries/clients.sql"
)

//go:embed queries
var queryFiles embed.FS

type SqliteClientRepo struct {
	db      *sql.DB
	queries map[string]string
}

func (c *SqliteClientRepo) GetByID(ctx context.Context, id int) (*domain.Client, error) {
	queryName := "get_client_by_id"
	query, ok := c.queries[queryName]
	if !ok {
		return nil, fmt.Errorf("No query found for query name: %s", queryName)
	}

	client := &domain.Client{}
	err := c.db.QueryRowContext(ctx, query, id).Scan(
		&client.Id, &client.Name, &client.IsPayingMe,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%w: id %d", service.ErrClientNotFound, id)
		}
		return nil, fmt.Errorf("get client by id %d: %w", id, err)
	}

	return client, err
}

func (c *SqliteClientRepo) GetByName(ctx context.Context, name string) (*domain.Client, error)

func (c *SqliteClientRepo) Safe(ctx context.Context, client *domain.Client) error

func (c *SqliteClientRepo) List(ctx context.Context) ([]*domain.Client, error)

func NewClientRepo(db *sql.DB) *SqliteClientRepo {
	queries, err := readQueriesFromFile(queryFiles, clientsSqlFileName)
	if err != nil {
		log.Fatal(err)
	}
	return &SqliteClientRepo{db: db, queries: queries}
}
