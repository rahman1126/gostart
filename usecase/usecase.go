package usecase

import (
	"context"
	"gostart/dto"
)

type ExampleUsecaseI interface {
	FindAll(ctx context.Context) ([]*dto.UserResponse, error)
	Find(ctx context.Context, id int) (*dto.UserResponse, error)
}
