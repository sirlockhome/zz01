package datastore

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"foxomni/internal/partner/model"
	"foxomni/pkg/errs"
)

func (ds *Datastore) GetPartnerByID(ctx context.Context, id int) (*model.Partner, error) {
	query := `
	SELECT 
		p.id, p.partner_name, p.partner_code, p.avatar_id, p.partner_type, p.tax_id, p.email, p.phone_number,
		p.country, p.city, p.district, p.ward, p.street, p.house_number
	FROM 
		partners p
	WHERE id = @p1
	`

	var partner model.Partner
	if err := ds.sql.DB.GetContext(ctx, &partner, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.New(errs.NotFound, fmt.Sprintf(ErrPartnerIDNotFound, id))
		}
		return nil, errs.New(errs.Internal, err)
	}

	return &partner, nil
}

func (ds *Datastore) GetPartners(ctx context.Context, offset, limit int) ([]*model.Partner, error) {
	query := `
	SELECT 
		p.id, p.partner_name, p.partner_code, p.avatar_id, p.partner_type, p.tax_id, p.email, p.phone_number,
		p.country, p.city, p.district, p.ward, p.street, p.house_number
	FROM 
		partners p
	ORDER BY p.id DESC
	OFFSET @p1 ROWS
	FETCH NEXT @p2 ROWS ONLY
	`

	rows, err := ds.sql.DB.QueryxContext(ctx, query, offset, limit)
	if err != nil {
		return nil, errs.New(errs.Internal, err)
	}

	defer rows.Close()

	var pnList []*model.Partner
	for rows.Next() {
		var pn model.Partner
		if err := rows.StructScan(&pn); err != nil {
			return nil, errs.New(errs.Internal, err)
		}
		pnList = append(pnList, &pn)
	}

	return pnList, nil
}

func (ds *Datastore) GetPartnerCount(ctx context.Context) (int, error) {
	query := `
	SELECT COUNT(*) FROM partners
	`

	var count int
	if err := ds.sql.DB.GetContext(ctx, &count, query); err != nil {
		return -1, errs.New(errs.Internal, err)
	}

	return count, nil
}
