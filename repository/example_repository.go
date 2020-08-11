package repository

import (
	"context"
	"github.com/jinzhu/gorm"
	"gostart/storage"
)

type ExampleRepository struct {
	Conn *gorm.DB
}

func NewExampleRepository(Conn *gorm.DB) ExampleRepositoryI {
	return &ExampleRepository{Conn}
}

func (er ExampleRepository) FindAll(ctx context.Context) ([]*storage.User, error) {
	users := []*storage.User{}

	er.Conn.Find(&users)

	return users, nil
}

func (er ExampleRepository) Find(ctx context.Context, id int) (*storage.User, error) {
	user := &storage.User{}

	er.Conn.Where("id = ?", id).First(&user)

	return user, nil
}