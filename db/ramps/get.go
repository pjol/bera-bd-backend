package ramps

import (
	"context"
	"fmt"

	"github.com/pjol/bera-bd-backend/structs"
)

func (r *Db) GetRegions(ctx context.Context) ([]structs.Region, error) {
	rows, err := r.db.Query(ctx, `
		SELECT
			id,
			name
		FROM
			regions;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var regions []structs.Region
	for rows.Next() {
		var region structs.Region
		err := rows.Scan(&region.Id, &region.Name)
		if err != nil {
			return nil, err
		}

		regions = append(regions, region)
	}

	return regions, nil
}

func (r *Db) GetAssets(ctx context.Context) ([]structs.Asset, error) {
	rows, err := r.db.Query(ctx, `
		SELECT
			id,
			ticker,
			address
		FROM
			assets;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var assets []structs.Asset
	for rows.Next() {
		var asset structs.Asset
		err := rows.Scan(&asset.Id, &asset.Ticker, &asset.Address)
		if err != nil {
			return nil, err
		}

		assets = append(assets, asset)
	}

	return assets, nil
}

func (r *Db) GetPaymentMethods(ctx context.Context) ([]structs.PaymentMethod, error) {
	rows, err := r.db.Query(ctx, `
		SELECT
			id,
			name
		FROM
			payment_methods;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var paymentMethods []structs.PaymentMethod
	for rows.Next() {
		var paymentMethod structs.PaymentMethod
		err := rows.Scan(&paymentMethod.Id, &paymentMethod.Name)
		if err != nil {
			return nil, err
		}

		paymentMethods = append(paymentMethods, paymentMethod)
	}

	return paymentMethods, nil
}

func (r *Db) GetRamps(ctx context.Context, all bool, approval bool) ([]*structs.Ramp, error) {
	statementSuffix := ";"
	if !all {
		statementSuffix = fmt.Sprintf("WHERE approval = %t", approval)
	}

	statement := fmt.Sprintf(`
		SELECT
			id,
			name,
			url,
			email,
			approval
		FROM
			ramps
		%s
	`, statementSuffix)

	rows, err := r.db.Query(ctx, statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ramps []*structs.Ramp
	for rows.Next() {
		ramp := structs.Ramp{}

		err = rows.Scan(
			&ramp.Id,
			&ramp.Name,
			&ramp.Url,
			&ramp.Email,
			&ramp.Approval,
		)
		if err != nil {
			return nil, err
		}

		ramps = append(ramps, &ramp)
	}

	// regions
	for _, ramp := range ramps {
		regions, err := r.GetRampRegions(ctx, ramp.Id)
		if err != nil {
			r.logger.Logf("error getting regions for ramp %d: %s", ramp.Id, err)
			break
		}

		ramp.Regions = regions
	}

	// assets
	for _, ramp := range ramps {
		assets, err := r.GetRampAssets(ctx, ramp.Id)
		if err != nil {
			r.logger.Logf("error getting assets for ramp %d: %s", ramp.Id, err)
			break
		}

		ramp.Assets = assets
	}

	// payment methods
	for _, ramp := range ramps {
		methods, err := r.GetRampPaymentMethods(ctx, ramp.Id)
		if err != nil {
			r.logger.Logf("error getting assets for ramp %d: %s", ramp.Id, err)
			break
		}

		ramp.PaymentMethods = methods
	}

	return ramps, nil
}

func (r *Db) GetRampRegions(ctx context.Context, id int) ([]structs.Region, error) {
	rows, err := r.db.Query(ctx, `
		SELECT
			re.id,
			re.name
		FROM
			regions re
		JOIN ramp_regions j ON j.region = re.id
		JOIN ramps ra ON j.ramp = ra.id
		WHERE ra.id = $1;
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var regions []structs.Region
	for rows.Next() {
		var region structs.Region
		err := rows.Scan(&region.Id, &region.Name)
		if err != nil {
			return nil, err
		}

		regions = append(regions, region)
	}

	return regions, nil
}

func (r *Db) GetRampAssets(ctx context.Context, id int) ([]structs.Asset, error) {
	rows, err := r.db.Query(ctx, `
		SELECT
			a.id,
			a.ticker,
			a.address
		FROM
			assets a
		JOIN ramp_assets j ON j.asset = a.id
		JOIN ramps r ON j.ramp = r.id
		WHERE r.id = $1;
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var assets []structs.Asset
	for rows.Next() {
		asset := structs.Asset{}
		err := rows.Scan(&asset.Id, &asset.Ticker, &asset.Address)
		if err != nil {
			return nil, err
		}

		assets = append(assets, asset)
	}

	return assets, nil
}

func (r *Db) GetRampPaymentMethods(ctx context.Context, id int) ([]structs.PaymentMethod, error) {
	rows, err := r.db.Query(ctx, `
		SELECT
			id,
			name
		FROM
			payment_methods p
		JOIN ramp_payment_methods j ON j.asset = p.id
		JOIN ramps r ON j.ramp = r.id
		WHERE r.id = $1;
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var paymentMethods []structs.PaymentMethod
	for rows.Next() {
		var paymentMethod structs.PaymentMethod
		err := rows.Scan(&paymentMethod.Id, &paymentMethod.Name)
		if err != nil {
			return nil, err
		}

		paymentMethods = append(paymentMethods, paymentMethod)
	}

	return paymentMethods, nil
}
