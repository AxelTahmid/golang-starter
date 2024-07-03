package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (s AuthService) GetUser(ctx context.Context, pool *pgxpool.Pool, email string) (User, error) {
	query := `SELECT * FROM users WHERE email = @userEmail;`

	args := pgx.NamedArgs{
		"userEmail": email,
	}

	row, err := pool.Query(ctx, query, args)
	if err != nil {
		return User{}, fmt.Errorf("unable to query row: %w", err)
	}

	user, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[User])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return User{}, fmt.Errorf("user not found: %s", email)
		}
		return User{}, fmt.Errorf("unable to scan row: %w", err)
	}

	return user, nil
}

func (s AuthService) InsertUser(ctx context.Context, pool *pgxpool.Pool, user RegisterRequest) error {
	query := `INSERT INTO users (name, email, password) VALUES (@userName, @userEmail, @hashedPassword);`

	args := pgx.NamedArgs{
		"userName":       user.Name,
		"userEmail":      user.Email,
		"hashedPassword": user.Password,
	}

	_, err := pool.Exec(ctx, query, args)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return fmt.Errorf("user already exists: %s", user.Email)
		}
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}
