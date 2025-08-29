package cfi

import (
	"context"
	"fmt"

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
		CREATE TABLE IF NOT EXISTS stages(
			id SERIAL PRIMARY KEY,
			stage TEXT NOT NULL
		);
	`)
	if err != nil {
		return fmt.Errorf("error creating stages table: %s", err)
	}

	_, err = a.db.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS cfi_apps(
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			website TEXT NOT NULL,
			email TEXT NOT NULL,
			stage INTEGER NOT NULL REFERENCES stages(id),
			org_name TEXT NOT NULL,
			vision_statement TEXT NOT NULL
		);
	`)
	if err != nil {
		return fmt.Errorf("error creating cfi_maps table: %s", err)
	}

	return nil
}
