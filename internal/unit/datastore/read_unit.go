package datastore

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"foxomni/internal/unit/model"
	"foxomni/pkg/errs"
)

func (ds *Datastore) GetUnitByID(ctx context.Context, id int) (*model.Unit, error) {
	query := `
	SELECT 
		u.id, u.unit_name, u.symbol, u.description, u.is_active, u.created_at, u.updated_at
	FROM units u
	WHERE u.id = @p1
	`

	var unit model.Unit
	if err := ds.sql.DB.GetContext(ctx, &unit, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.New(errs.NotFound, fmt.Sprintf("unit with id `%d` not found", id))
		}
		return nil, errs.New(errs.Internal, err)
	}

	return &unit, nil
}

func (ds *Datastore) GetUnits(ctx context.Context, offset, limit int) ([]*model.Unit, error) {
	query := `
	SELECT 
		u.id, u.unit_name, u.symbol, u.description, u.is_active, u.created_at, u.updated_at
	FROM units u
	ORDER BY u.id DESC
	OFFSET @p1 ROWS
	FETCH NEXT @p2 ROWS ONLY
	`

	rows, err := ds.sql.DB.QueryxContext(ctx, query, offset, limit)
	if err != nil {
		return nil, errs.New(errs.Internal, err)
	}

	defer rows.Close()

	var unitList []*model.Unit
	for rows.Next() {
		var unit model.Unit
		if err := rows.StructScan(&unit); err != nil {
			return nil, errs.New(errs.Internal, err)
		}

		unitList = append(unitList, &unit)
	}

	return unitList, err
}

func (ds *Datastore) GetUnitCount(ctx context.Context) (int, error) {
	query := `
	SELECT COUNT(*) FROM units
	`
	var count int
	if err := ds.sql.DB.GetContext(ctx, &count, query); err != nil {
		return -1, errs.New(errs.Internal, err)
	}

	return count, nil
}
