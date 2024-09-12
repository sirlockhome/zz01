package datastore

import (
	"context"
	"foxomni/pkg/errs"
)

func (ds *Datastore) ChangeStatusOrder(ctx context.Context, id int, status string) error {
	query := `
	UPDATE orders SET status = @p1 WHERE id = @p2
	`
	_, err := ds.sql.DB.ExecContext(ctx, query, status, id)
	if err != nil {
		return errs.New(errs.Internal, err)
	}

	return nil
}
