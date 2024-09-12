package datastore

import (
	"context"
	"foxomni/internal/product/model"
	"foxomni/pkg/errs"
)

func (ds *Datastore) UpdateProduct(ctx context.Context, prd *model.ProductRequest) error {
	query := `
	UPDATE products SET 
		name = :name, description = :description, information = :information, feature = :feature,
		specifications = :specifications, sku = :sku, price = :price, stock = :stock
	WHERE id = :id
	`

	_, err := ds.sql.DB.NamedExecContext(ctx, query, &prd)
	if err != nil {
		return errs.New(errs.Internal, err)
	}

	return nil
}

func (ds *Datastore) DeleteProduct(ctx context.Context, id int) error {
	query := `
	DELETE products WHERE id = @p1
	`
	_, err := ds.sql.DB.ExecContext(ctx, query, id)
	if err != nil {
		return errs.New(errs.Internal, err)
	}

	return nil
}

func (ds *Datastore) ChangeThumbnail(ctx context.Context, prdID int, imgID string) error {
	query := `
	UPDATE products SET
		thumbnail_id = @p1
	WHERE id = @p2
	`

	_, err := ds.sql.DB.ExecContext(ctx, query, imgID, prdID)
	if err != nil {
		return errs.New(errs.Internal, err)
	}

	return nil
}

func (ds *Datastore) ChangeActive(ctx context.Context, prdID int, isActive bool) error {
	query := `
	UPDATE products SET
		is_active = @p1
	WHERE id = @p2
	`
	_, err := ds.sql.DB.ExecContext(ctx, query, isActive, prdID)
	if err != nil {
		return errs.New(errs.Internal, err)
	}

	return nil
}
