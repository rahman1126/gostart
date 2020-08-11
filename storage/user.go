package storage

import "time"

type User struct {
	ID int
	Name  string
	Email string
	Phone string
	CreatedAt time.Time
}