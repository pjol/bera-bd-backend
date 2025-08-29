package auth

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pjol/bera-bd-backend/logger"
)

type Db struct {
	db     *pgxpool.Pool
	logger *logger.LogCloser
}

func NewDb(db *pgxpool.Pool, logger *logger.LogCloser) *Db {
	return &Db{db, logger}
}

func (a *Db) CreateTables() error {
	_, err := a.db.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS users(
			id SERIAL PRIMARY KEY,
			did TEXT NOT NULL,
			is_admin BOOLEAN NOT NULL DEFAULT FALSE
		);

		CREATE UNIQUE INDEX IF NOT EXISTS user_did ON users(did);
	`)
	return err
}

func (a *Db) IsAdmin(ctx context.Context, did string) (bool, error) {
	row := a.db.QueryRow(ctx, `
		SELECT
			is_admin
		FROM
			users
		WHERE
			did = $1;
	`, did)
	var isAdmin bool
	err := row.Scan(&isAdmin)
	if err == pgx.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return isAdmin, nil
}
