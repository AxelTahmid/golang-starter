package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// exact order as in database
type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Verified  bool      `json:"verified"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

const (
	InsertUserQuery     = "INSERT INTO users (name, email, password) VALUES (@userName, @userEmail, @hashedPassword);"
	GetUserByEmailQuery = "SELECT * FROM users WHERE email = @userEmail;"
)

func (u *User) GetByEmail(ctx context.Context, email string) (User, error) {
	args := pgx.NamedArgs{
		"userEmail": email,
	}

	row, err := pg.pool.Query(ctx, GetUserByEmailQuery, args)
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

type InsertUser struct {
	Name     string `json:"name"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (u *User) InsertOne(ctx context.Context, user InsertUser) error {
	args := pgx.NamedArgs{
		"userName":       user.Name,
		"userEmail":      user.Email,
		"hashedPassword": user.Password,
	}

	_, err := pg.pool.Exec(ctx, InsertUserQuery, args)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return fmt.Errorf("user already exists: %s", user.Email)
		}
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}
