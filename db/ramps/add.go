package ramps

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/pjol/bera-bd-backend/structs"
)

func (r *Db) AddRamp(ctx context.Context, ramp *structs.Ramp) (int, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("error beginnning addramp tx: %s", err)
	}
	defer tx.Rollback(ctx)

	row := tx.QueryRow(ctx, `
		INSERT INTO ramps(
			name,
			url,
			email,
		) VALUES (
			$1,
			$2,
			$3,
		)
		RETURNING id;
	`, ramp.Id, ramp.Url, ramp.Email)
	var id int
	err = row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error scanning row id: %s", err)
	}

	err = r.AddRampRegions(ctx, tx, id, ramp.Regions)
	if err != nil {
		return 0, fmt.Errorf("error adding ramp regions: %s", err)
	}

	err = r.AddRampAssets(ctx, tx, id, ramp.Assets)
	if err != nil {
		return 0, fmt.Errorf("error adding ramp assets: %s", err)
	}

	err = r.AddRampPaymentMethods(ctx, tx, id, ramp.PaymentMethods)
	if err != nil {
		return 0, fmt.Errorf("error adding ramp payment methods: %s", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return 0, fmt.Errorf("error committing tx: %s", err)
	}
	return id, nil
}

func (r *Db) AddRampRegions(ctx context.Context, tx pgx.Tx, rampId int, regions []structs.Region) error {
	for _, region := range regions {
		_, err := tx.Exec(ctx, `
			INSERT INTO ramp_regions(
				ramp,
				region
			) VALUES (
				$1,
				$2
			)
			ON CONFLICT (ramp, region)
			DO NOTHING;
		`, rampId, region.Id)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Db) AddRampAssets(ctx context.Context, tx pgx.Tx, rampId int, assets []structs.Asset) error {
	for _, asset := range assets {
		_, err := tx.Exec(ctx, `
			INSERT INTO ramp_assets(
				ramp,
				asset
			) VALUES (
				$1,
				$2
			)
			ON CONFLICT (ramp, asset)
			DO NOTHING;
		`, rampId, asset.Id)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Db) AddRampPaymentMethods(ctx context.Context, tx pgx.Tx, rampId int, paymentMethods []structs.PaymentMethod) error {
	for _, asset := range paymentMethods {
		_, err := tx.Exec(ctx, `
			INSERT INTO ramp_payment_methods(
				ramp,
				payment_method
			) VALUES (
				$1,
				$2
			)
			ON CONFLICT (ramp, payment_method)
			DO NOTHING;
		`, rampId, asset.Id)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Db) AddRegion(ctx context.Context, region structs.Region) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO regions(
			name
		) VALUES (
			$1
		)
		ON CONFLICT(name)
		DO NOTHING;
	`, region.Name)

	return err
}

func (r *Db) AddAsset(ctx context.Context, method structs.Asset) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO payment_methods(
			ticker,
			address
		) VALUES (
			$1,
			$2
		)
		ON CONFLICT(address)
		DO NOTHING;
	`, method.Ticker, method.Address)

	return err
}

func (r *Db) AddPaymentMethod(ctx context.Context, method structs.PaymentMethod) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO payment_methods(
			name
		) VALUES (
			$1
		)
		ON CONFLICT(name)
		DO NOTHING;
	`, method.Name)

	return err
}
