package datastore

import (
	"context"
	"foxomni/internal/partner/model"
	"foxomni/pkg/errs"

	"github.com/jmoiron/sqlx"
)

func (ds *Datastore) CreatePartner(ctx context.Context, pn *model.PartnerRequest) (int, error) {
	query := `
	INSERT INTO partners 
	(
		partner_name, partner_code, partner_type, avatar_id, tax_id, email, phone_number,
		country, city, district, ward, street, house_number
	) OUTPUT INSERTED.id
	VALUES
	(
		:partner_name, :partner_code, :partner_type, :avatar_id, :tax_id, :email, :phone_number,
		:country, :city, :district, :ward, :street, :house_number
	)
	`

	var idNewPn int
	err := ds.sql.InTx(ctx, func(tx *sqlx.Tx) error {
		rows, err := tx.NamedQuery(query, pn)
		if err != nil {
			return errs.New(errs.Internal, err)
		}

		defer rows.Close()
		if rows.Next() {
			if err := rows.Scan(&idNewPn); err != nil {
				return errs.New(errs.Internal, err)
			}
		} else {
			return errs.New(errs.Internal, rows.Err())
		}

		return nil
	})

	if err != nil {
		return -1, err
	}

	return idNewPn, nil
}
