package datastore

import (
	"context"
	"foxomni/internal/unit/model"
	"foxomni/pkg/errs"

	"github.com/jmoiron/sqlx"
)

func (ds *Datastore) CreateUnitGroup(ctx context.Context, grp *model.UnitGroup) (int, error) {
	query := `
	INSERT INTO unit_groups
	(
		unit_group_name, base_unit_id, symbol, description
	) OUTPUT INSERTED.id
	VALUES
	(
		:unit_group_name, :base_unit_id, :symbol, :description
	)
	`

	var idNewGrp int

	err := ds.sql.InTx(ctx, func(tx *sqlx.Tx) error {
		rows, err := tx.NamedQuery(query, grp)
		if err != nil {
			return errs.New(errs.Internal, err)
		}

		defer rows.Close()
		if rows.Next() {
			if err := rows.Scan(&idNewGrp); err != nil {
				return errs.New(errs.Internal, err)
			}
		} else {
			return errs.New(errs.Internal, rows.Err())
		}

		err = ds.addConversionWithTx(ctx, tx, grp.UnitConversions, idNewGrp)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return -1, err
	}

	return idNewGrp, nil
}

func (ds *Datastore) AddConversions(ctx context.Context, convList []*model.UnitConversion, grpID int) error {
	return ds.sql.InTx(ctx, func(tx *sqlx.Tx) error {
		return ds.addConversionWithTx(ctx, tx, convList, grpID)
	})
}

func (ds *Datastore) addConversionWithTx(ctx context.Context, tx *sqlx.Tx, convList []*model.UnitConversion, grpID int) error {
	query := `
	INSERT INTO unit_conversions
	(
		to_unit_id, base_qty, alt_qty, unit_group_id
	)
	VALUES
	(
		:to_unit_id, :base_qty, :alt_qty, :unit_group_id
	)
	`

	for _, conv := range convList {
		conv.UnitGroupID = grpID
		_, err := tx.NamedExecContext(ctx, query, conv)
		if err != nil {
			return errs.New(errs.Internal, err)
		}
	}

	return nil
}
