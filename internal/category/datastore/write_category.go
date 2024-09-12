package datastore

import (
	"context"
	"foxomni/internal/category/model"
	"foxomni/pkg/errs"
)

func (ds *Datastore) CreateCategory(ctx context.Context, cat *model.Category) (int, error) {
	query := `
	INSERT INTO categories
	(
		category_name, description, parent_id, thumbnail_id
	) OUTPUT INSERTED.id
	VALUES
	(
		:category_name, :description, :parent_id, :thumbnail_id
	)
	`

	var idNewCat int
	rows, err := ds.sql.DB.NamedQueryContext(ctx, query, cat)
	if err != nil {
		return -1, errs.New(errs.Internal, err)
	}

	defer rows.Close()
	if rows.Next() {
		if err := rows.Scan(&idNewCat); err != nil {
			return -1, errs.New(errs.Internal, err)
		}
	} else {
		return -1, errs.New(errs.Internal, rows.Err())
	}

	return idNewCat, nil
}

func (ds *Datastore) UpdateCategory(ctx context.Context, cat *model.Category) error {
	query := `
	UPDATE categories SET 
		category_name = :category_name, description = :description,
		parent_id = :parent_id, thumbnail_id = :thumbnail_id
	WHERE id = :id
	`

	_, err := ds.sql.DB.ExecContext(ctx, query, cat)
	if err != nil {
		return errs.New(errs.Internal, err)
	}

	return nil
}
