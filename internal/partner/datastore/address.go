package datastore

import (
	"context"
	"database/sql"
	"errors"
	"foxomni/internal/partner/model"
	"foxomni/pkg/errs"
)

func (ds *Datastore) AddAddress(ctx context.Context, addr *model.Address) error {
	query := `
	INSERT INTO partner_addresses
	(
		partner_id, country, city, district, ward, street, house_number,
		address_line1, address_line2
	)
	VALUES
	(
		:partner_id, :country, :city, :district, :ward, :street, :house_number,
		:address_line1, :address_line2
	)
	`

	_, err := ds.sql.DB.NamedExecContext(ctx, query, addr)
	if err != nil {
		return errs.New(errs.Internal, err)
	}

	return nil
}

func (ds *Datastore) RemoveAddress(ctx context.Context, addrID int, pnID int) error {
	query := `
	DELETE partner_addresses WHERE id = @p1 AND partner_id = @p2
	`

	_, err := ds.sql.DB.ExecContext(ctx, query, addrID, pnID)
	if err != nil {
		return errs.New(errs.Internal, err)
	}

	return nil
}

func (ds *Datastore) GetAddresses(ctx context.Context, pnID int) ([]*model.Address, error) {
	query := `
	SELECT
		id, partner_id, country, city, district, ward, street,
		house_number, is_default_shipping_address, is_billing_address, address_line1, address_line2
	FROM 
		partner_addresses
	WHERE partner_id = @p1
	`

	rows, err := ds.sql.DB.QueryxContext(ctx, query, &pnID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, errs.New(errs.Internal, err)
	}

	defer rows.Close()

	var addrList []*model.Address
	for rows.Next() {
		var addr model.Address
		if err := rows.StructScan(&addr); err != nil {
			return nil, errs.New(errs.Internal, err)
		}

		addrList = append(addrList, &addr)
	}

	return addrList, nil
}
