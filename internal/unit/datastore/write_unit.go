package datastore

import (
	"context"
	"foxomni/internal/unit/model"
	"foxomni/pkg/errs"
)

func (ds *Datastore) CreateUnit(ctx context.Context, unit *model.Unit) (int, error) {
	query := `
	INSERT INTO units
	(
		unit_name, symbol, description
	) OUTPUT INSERTED.id
	VALUES 
	(
		:unit_name, :symbol, :description
	)
	`

	rows, err := ds.sql.DB.NamedQueryContext(ctx, query, unit)
	if err != nil {
		return -1, errs.New(errs.Internal, err)
	}

	var idNewUnit int

	defer rows.Close()
	if rows.Next() {
		if err := rows.Scan(&idNewUnit); err != nil {
			return -1, errs.New(errs.Internal, err)
		}
	} else {
		return -1, errs.New(errs.Internal, rows.Err())
	}

	return idNewUnit, nil
}

func (ds *Datastore) UpdateUnit(ctx context.Context, unit *model.Unit) error {
	query := `
	UPDATE unit SET 
		unit_name = :unit_name, symbol = :symbol, description = :description
	WHERE id = :id
	`

	_, err := ds.sql.DB.NamedExecContext(ctx, query, unit)
	if err != nil {
		return errs.New(errs.Internal, err)
	}

	return nil
}

func (ds *Datastore) DeleteUnit(ctx context.Context, id int) error {
	query := `DELETE units WHERE id = @p1`

	_, err := ds.sql.DB.ExecContext(ctx, query, id)
	if err != nil {
		return errs.New(errs.Internal, err)
	}

	return nil
}
