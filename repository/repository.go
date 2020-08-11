package repository

import (
	"context"
	"gostart/storage"
)

type ExampleRepositoryI interface {
	FindAll(ctx context.Context) ([]*storage.User, error)
	Find(ctx context.Context, id int) (*storage.User, error)
}
