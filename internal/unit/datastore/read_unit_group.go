package datastore

import (
	"context"
	"database/sql"
	"errors"
	"foxomni/internal/unit/model"
	"foxomni/pkg/errs"
)

func (ds *Datastore) GetUnitGroupByID(ctx context.Context, id int) (*model.UnitGroup, error) {
	query := `
	SELECT 
		ug.id, ug.unit_group_name, ug.base_unit_id, u.unit_name AS base_unit_name, ug.symbol,
		ug.description, ug.created_at, ug.updated_at 
	FROM unit_groups ug 
	LEFT JOIN units u ON ug.base_unit_id = u.id
	WHERE ug.id = @p1
	`

	var unitGroup model.UnitGroup
	if err := ds.sql.DB.GetContext(ctx, &unitGroup, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.New(errs.NotFound)
		}
		return nil, errs.New(errs.Internal, err)
	}

	convList, err := ds.GetUnitConversion(ctx, id)
	if err != nil {
		return nil, errs.New(errs.Internal, err)
	}

	unitGroup.UnitConversions = convList

	return &unitGroup, nil
}

func (ds *Datastore) GetUnitGroups(ctx context.Context, offset, limit int) ([]*model.UnitGroup, error) {
	query := `
	SELECT 
		ug.id, ug.unit_group_name, ug.base_unit_id, u.unit_name AS base_unit_name, ug.symbol,
		ug.description, ug.created_at, ug.updated_at 
	FROM unit_groups ug 
	LEFT JOIN units u ON ug.base_unit_id = u.id
	ORDER BY ug.id DESC
	OFFSET @p1 ROWS
	FETCH NEXT @p2 ROWS ONLY
	`

	rows, err := ds.sql.DB.QueryxContext(ctx, query, offset, limit)
	if err != nil {
		return nil, errs.New(errs.Internal, err)
	}

	defer rows.Close()

	var listGrp []*model.UnitGroup
	for rows.Next() {
		var grp model.UnitGroup
		if err := rows.StructScan(&grp); err != nil {
			return nil, errs.New(errs.Internal, err)
		}

		listGrp = append(listGrp, &grp)
	}

	return listGrp, nil
}

func (ds *Datastore) GetUnitGroupCount(ctx context.Context) (int, error) {
	query := `SELECT COUNT(*) FROM unit_groups`

	var count int
	if err := ds.sql.DB.GetContext(ctx, &count, query); err != nil {
		return -1, errs.New(errs.Internal, err)
	}

	return count, nil
}

func (ds *Datastore) GetUnitConversion(ctx context.Context, grpID int) ([]*model.UnitConversion, error) {
	query := `
	SELECT 
		ucv.id, ucv.to_unit_id, u.unit_name AS to_unit_name,
		ucv.base_qty, ucv.alt_qty
	FROM unit_conversions ucv
	LEFT JOIN unit_groups ug ON ucv.unit_group_id = ug.id
	LEFT JOIN units u ON u.id = ucv.to_unit_id
	WHERE ug.id = @p1
	`

	rows, err := ds.sql.DB.QueryxContext(ctx, query, grpID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, errs.New(errs.Internal, err)
	}

	defer rows.Close()

	var convList []*model.UnitConversion
	for rows.Next() {
		var conv model.UnitConversion
		if err := rows.StructScan(&conv); err != nil {
			return nil, errs.New(errs.Internal, err)
		}

		convList = append(convList, &conv)
	}

	return convList, nil
}
