package datastore

import (
	"context"
	"foxomni/internal/product/model"
	"foxomni/pkg/errs"

	"github.com/jmoiron/sqlx"
)

func (ds *Datastore) CreateProduct(ctx context.Context, prd *model.ProductRequest) (int, error) {
	query := `
	INSERT INTO products
	(
		name, description, information, feature, specifications, sku, price, stock,
		thumbnail_id, unit_group_id
	)
	OUTPUT INSERTED.id
	VALUES
	(
		:name, :description, :information, :feature, :specifications, :sku, :price, :stock,
		:thumbnail_id, :unit_group_id
	)
	`

	var idNewPd int
	err := ds.sql.InTx(ctx, func(tx *sqlx.Tx) error {
		rows, err := tx.NamedQuery(query, prd)
		if err != nil {
			return errs.New(errs.Internal, err)
		}

		defer rows.Close()
		if rows.Next() {
			if err := rows.Scan(&idNewPd); err != nil {
				return errs.New(errs.Internal, err)
			}
		} else {
			return errs.New(errs.Internal, rows.Err())
		}

		if err := addImagesWithTx(ctx, tx, prd.ProductImages, idNewPd); err != nil {
			return err
		}

		if err := addAttchmentsWithTx(ctx, tx, prd.ProductAttachments, idNewPd); err != nil {
			return err
		}

		if err := addCategoriesWithTx(ctx, tx, prd.ProductCategories, idNewPd); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return -1, err
	}

	return idNewPd, nil
}
