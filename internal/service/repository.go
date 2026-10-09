package service

import (
	"context"
	"keep_time/internal/domain"
	"errors"
)

type ClientRepo interface {
	GetByID(ctx context.Context, id int) (*domain.Client, error)
	GetByName(ctx context.Context, name string) (*domain.Client, error)
	Safe(ctx context.Context, client *domain.Client) error
	List(ctx context.Context) ([]*domain.Client, error)
}

var ErrClientNotFound = errors.New("client not found")