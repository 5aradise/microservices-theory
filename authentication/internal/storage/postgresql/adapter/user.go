package adapter

import (
	"micro/authentication/internal/model"
	"micro/authentication/internal/storage/postgresql/queries"

	"github.com/jackc/pgx/v5/pgtype"
)

func DBToUser(u queries.User) model.User {
	return model.User{
		ID:             int(u.ID),
		Email:          u.Email,
		FirstName:      u.FirstName.String,
		LastName:       u.LastName.String,
		HashedPassword: u.Password,
		Active:         u.Active,
		CreatedAt:      u.CreatedAt.Time,
		UpdatedAt:      u.UpdatedAt.Time,
	}
}

func UserToDB(u model.User) queries.User {
	return queries.User{
		ID:        int32(u.ID),
		Email:     u.Email,
		FirstName: pgtype.Text{String: u.FirstName, Valid: true},
		LastName:  pgtype.Text{String: u.LastName, Valid: true},
		Password:  u.HashedPassword,
		Active:    u.Active,
		CreatedAt: pgtype.Timestamptz{Time: u.CreatedAt, Valid: true},
		UpdatedAt: pgtype.Timestamptz{Time: u.UpdatedAt, Valid: true},
	}
}
