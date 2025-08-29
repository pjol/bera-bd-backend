package ramps

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
		CREATE TABLE IF NOT EXISTS ramps(
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			url TEXT NOT NULL,
			email TEXT NOT NULL,
			approval boolean NOT NULL DEFAULT false
		);
	`)
	if err != nil {
		return fmt.Errorf("error creating ramps table: %s", err)
	}

	_, err = a.db.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS regions(
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL
		);
	`)
	if err != nil {
		return fmt.Errorf("error creating regions table: %s", err)
	}

	_, err = a.db.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS ramp_regions(
			id SERIAL PRIMARY KEY,
			ramp INTEGER NOT NULL REFERENCES ramps(id),
			region INTEGER NOT NULL REFERENCES regions(id)
		);
	`)
	if err != nil {
		return fmt.Errorf("error creating ramp_regions table: %s", err)
	}

	_, err = a.db.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS assets(
			id SERIAL PRIMARY KEY,
			ticker TEXT NOT NULL,
			address TEXT NOT NULL
		);
	`)
	if err != nil {
		return fmt.Errorf("error creating assets table: %s", err)
	}

	_, err = a.db.Exec(context.Background(), `
		CREATE TABLE IF NOT EXITST ramp_assets(
			id SERIAL PRIMARY KEY,
			ramp INTEGER NOT NULL REFERENCES ramps(id),
			asset INTEGER NOT NULL REFERENCES assets(id)
		);
	`)
	if err != nil {
		return fmt.Errorf("error creating ramp_assets table: %s", err)
	}

	_, err = a.db.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS payment_methods(
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL
		);
	`)
	if err != nil {
		return fmt.Errorf("error creating payment_methods table: %s", err)
	}

	_, err = a.db.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS ramp_payment_methods(
			id SERIAL PRIMARY KEY,
			ramp INTEGER NOT NULL REFERENCES ramps(id),
			payment_method INTEGER NOT NULL REFERENCES payment_methods(id)
		);
	`)
	if err != nil {
		return fmt.Errorf("error creating ramp_payment_methods table: %s", err)
	}

	return err
}
