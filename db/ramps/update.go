package ramps

import "context"

func (r *Db) UpdateRampApproval(ctx context.Context, id int, approval bool) error {
	_, err := r.db.Exec(ctx, `
		UPDATE
			ramps
		SET
			approval = $1
		WHERE
			id = $2;
	`, approval, id)

	return err
}
