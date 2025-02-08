package store

import (
	"context"
	"database/sql"
)

type UsersStore struct {
	db        *sql.DB
	User      string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
}

func (s *UsersStore) Create(ctx context.Context) error {
	return nil
}
