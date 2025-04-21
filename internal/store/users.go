package store

import (
	"context"
	"database/sql"
)

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}

type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO
			users (
				id,
				username,
				password
			)
		VALUES (
			$1,
			$2,
			$3
		)
		RETURNING 
			id, 
			created_at
	`
	err := s.db.QueryRowContext(
		ctx,
		query,
		user.ID,
		user.Username,
		user.Password,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
