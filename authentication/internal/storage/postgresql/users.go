package storage

import (
	"context"
	"micro/authentication/internal/model"
	"micro/authentication/internal/storage/postgresql/adapter"
	"micro/authentication/internal/storage/postgresql/queries"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Users struct {
	db *pgxpool.Pool
}

func NewUsers(db *pgxpool.Pool) *Users {
	return new(Users{
		db: db,
	})
}

func (s *Users) queries() queries.Queries {
	return queries.New(s.db)
}

func (s *Users) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	u, err := s.queries().GetByEmail(ctx, email)
	if err != nil {
		return model.User{}, err
	}
	return adapter.DBToUser(u), nil
}
