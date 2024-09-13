package datastore

import (
	"context"
	"foxomni/internal/partner/model"
	"foxomni/pkg/errs"
)

func (ds *Datastore) CreatePartnerGroup(ctx context.Context, grp *model.PartnerGroup) (int, error) {
	query := `
	INSERT INTO partner_groups
	(
		group_code, group_name, group_type, description
	) OUTPUT INSERTED.id
	VALUES 
	(
		:group_code, :group_name, :group_type, :description
	)
	`

	rows, err := ds.sql.DB.NamedQueryContext(ctx, query, grp)
	if err != nil {
		return -1, errs.New(errs.Internal, err)
	}

	defer rows.Close()

	var idNewGrp int
	if rows.Next() {
		if err := rows.Scan(idNewGrp); err != nil {
			return -1, errs.New(errs.Internal, err)
		}
	} else {
		return -1, errs.New(errs.Internal, rows.Err())
	}

	return idNewGrp, nil
}

func (ds *Datastore) UpdatePartnerGroup(ctx context.Context, grp *model.PartnerGroup) error {
	return nil
}
