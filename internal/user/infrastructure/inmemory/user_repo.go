package inmemory

import (
	"context"

	"github.com/Mezrik/fencing-dp/internal/user"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type InMemoryUserRepository struct {
	db  *sqlx.DB
	ctx context.Context
}

func NewInMemoryUserRepository(ctx context.Context, db *sqlx.DB) user.UserRepository {
	return &InMemoryUserRepository{db: db}
}

func (r *InMemoryUserRepository) GetUser(id uuid.UUID) (*user.User, error) {
	return nil, nil
}
