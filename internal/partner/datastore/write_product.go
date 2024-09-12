package datastore

import (
	"context"
	"foxomni/internal/partner/model"
	"foxomni/pkg/errs"
)

func (ds *Datastore) UpdatePartner(ctx context.Context, pn *model.PartnerRequest) error {
	query := `
	UPDATE partners SET
		partner_name = :partner_name, partner_type = :partner_type, tax_id = :tax_id,
		email = :email, phone_number = :phone_number, country = :country, city = :city,
		district = :district, ward = :ward, street = :street, house_number = :house_number
	WHERE id = :id
	`

	_, err := ds.sql.DB.NamedExecContext(ctx, query, pn)
	if err != nil {
		return errs.New(errs.Internal, err)
	}

	return nil
}

func (ds *Datastore) DeletePartner(ctx context.Context, id int) error {
	query := ``

	_, err := ds.sql.DB.ExecContext(ctx, query, id)
	if err != nil {
		return errs.New(errs.Internal, err)
	}

	return nil
}
