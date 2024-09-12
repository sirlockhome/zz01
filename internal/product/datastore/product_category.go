package datastore

import (
	"context"
	"foxomni/internal/product/model"
	"foxomni/pkg/errs"

	"github.com/jmoiron/sqlx"
)

func (ds *Datastore) GetCategories(ctx context.Context, prdID int) ([]*model.ProductCategory, error) {
	query := `
	SELECT 
		pc.category_id, ct.category_name, ct.thumbnail_id, uf.file_path as thumbnail_path
	FROM 
		product_categories pc 
	LEFT JOIN categories ct ON pc.category_id = ct.id
	LEFT JOIN uploaded_files uf ON uf.file_id = ct.thumbnail_id 
	WHERE pc.product_id = @p1
	`

	rows, err := ds.sql.DB.QueryxContext(ctx, query, prdID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var catList []*model.ProductCategory
	for rows.Next() {
		var cat model.ProductCategory
		if err := rows.StructScan(&cat); err != nil {
			return nil, err
		}

		catList = append(catList, &cat)
	}

	return catList, nil
}

func (ds *Datastore) AddCategories(ctx context.Context, catList []*model.ProductCategory, prdID int) error {
	return ds.sql.InTx(ctx, func(tx *sqlx.Tx) error {
		return addCategoriesWithTx(ctx, tx, catList, prdID)
	})
}

func (ds *Datastore) RemoveCategory(ctx context.Context, prdID int, catID int) error {
	query := `
	DELETE product_categories WHERE product_id = @p1 AND category_id = @p2
	`

	_, err := ds.sql.DB.ExecContext(ctx, query, prdID, catID)
	if err != nil {
		return errs.New(errs.Internal, err)
	}

	return nil
}

func addCategoriesWithTx(ctx context.Context, tx *sqlx.Tx, catList []*model.ProductCategory, prdID int) error {
	query := `
	INSERT INTO product_categories
	(
		category_id, product_id
	)
	VALUES
	(
		:category_id, :product_id
	)
	`

	for _, cat := range catList {
		cat.ProductID = prdID
		_, err := tx.NamedExecContext(ctx, query, &cat)
		if err != nil {
			return errs.New(errs.Internal, err)
		}
	}

	return nil
}
