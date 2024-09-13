package datastore

import (
	"context"
	"database/sql"
	"errors"
	"foxomni/internal/partner/model"
	"foxomni/pkg/errs"
)

func (ds *Datastore) GetPartnerGroupByID(ctx context.Context, id int) (*model.PartnerGroup, error) {
	query := `
	SELECT
		id, group_code, group_name, description, group_type,
		created_at, updated_at
	FROM 
		partner_groups
	WHERE id = @p1
	`

	var pnGrp model.PartnerGroup
	if err := ds.sql.DB.GetContext(ctx, &pnGrp, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.New(errs.NotFound)
		}
		return nil, errs.New(errs.Internal, err)
	}

	return &pnGrp, nil
}

func (ds *Datastore) GetPartnerGroups(ctx context.Context, offset, limit int) ([]*model.PartnerGroup, error) {
	query := `
	SELECT
		id, group_code, group_name, description, group_type,
		created_at, updated_at
	FROM 
		partner_groups
	ORDER BY id
	OFFSET @p1 ROWS
	FETCH NEXT @p2 ROWS ONLY
	`

	rows, err := ds.sql.DB.QueryxContext(ctx, query, offset, limit)
	if err != nil {
		return nil, errs.New(errs.Internal, err)
	}

	defer rows.Close()

	var grpList []*model.PartnerGroup
	for rows.Next() {
		var grp model.PartnerGroup
		if err := rows.StructScan(&rows); err != nil {
			return nil, errs.New(errs.Internal, err)
		}

		grpList = append(grpList, &grp)
	}

	return grpList, nil
}

func (ds *Datastore) GetPartnerGroupCount(ctx context.Context) (int, error) {
	query := `SELECT COUNT(*) FROM partner_groups`

	var count int
	if err := ds.sql.DB.GetContext(ctx, &count, query); err != nil {
		return -1, errs.New(errs.Internal, err)
	}

	return count, nil
}
