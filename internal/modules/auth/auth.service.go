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

type Postgres struct {
	db *pgxpool.Pool
}

// database interactions

func (pg *Postgres) InsertUser(ctx context.Context) error {
	query := `INSERT INTO users (name, email, password) VALUES (@userName, @userEmail, @hashedPassword);`

	args := pgx.NamedArgs{
		"userName":       "Bobby",
		"userEmail":      "bobby@donchev.is",
		"hashedPassword": "$2a$10$1Q6",
	}

	_, err := pg.db.Exec(ctx, query, args)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			// log.Printf("user %s already exists", user.Name)
			return fmt.Errorf("user already exists: %w", err)
		}
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}
