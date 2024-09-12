package datastore

import (
	"context"
	"database/sql"
	"errors"
	"foxomni/internal/category/model"
	"foxomni/pkg/errs"
)

func (ds *Datastore) GetCategoryByID(ctx context.Context, id int) (*model.Category, error) {
	query := `
	SELECT 
		c.id, c.category_name, c.description, c.parent_id, c.thumbnail_id,
		c.created_at, c.updated_at
	FROM 
		categories c
	WHERE c.id = @p1
	`

	var cat model.Category
	if err := ds.sql.DB.GetContext(ctx, &cat, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.New(errs.NotFound)
		}
		return nil, errs.New(errs.Internal, err)
	}

	return &cat, nil
}

func (ds *Datastore) GetCategories(ctx context.Context, offset, limit int) ([]*model.Category, error) {
	query := `
	SELECT 
		c.id, c.category_name, c.description, c.parent_id, c.thumbnail_id,
		c.created_at, c.updated_at
	FROM 
		categories c
	ORDER BY c.id DESC
	OFFSET @p1 ROWS
	FETCH NEXT @p2 ROWS ONLY
	`

	rows, err := ds.sql.DB.QueryxContext(ctx, query, offset, limit)
	if err != nil {
		return nil, errs.New(errs.Internal, err)
	}

	defer rows.Close()

	var catList []*model.Category
	for rows.Next() {
		var cat model.Category
		if err := rows.StructScan(&cat); err != nil {
			return nil, errs.New(errs.Internal, err)
		}

		catList = append(catList, &cat)
	}

	return catList, nil
}

func (ds *Datastore) GetCategoryCount(ctx context.Context) (int, error) {
	query := `SELECT COUNT(*) FROM categories`

	var count int
	if err := ds.sql.DB.GetContext(ctx, &count, query); err != nil {
		return -1, errs.New(errs.Internal, err)
	}

	return count, nil
}
