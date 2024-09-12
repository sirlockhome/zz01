package datastore

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"foxomni/internal/common"
	"foxomni/internal/product/model"
	"foxomni/pkg/errs"
)

func (ds *Datastore) GetProductByID(ctx context.Context, id int) (*model.Product, error) {
	query := `
	SELECT 
		p.id, p.name, p.description, p.sku, p.information, p.specifications, p.price, p.stock,
		p.unit_group_id, p.thumbnail_id, f.file_path as thumbnail_path,
		un.id AS unit_price_id, un.unit_name AS unit_price_name,
		p.created_at, p.updated_at
	FROM 
		products p
	LEFT JOIN uploaded_files f ON p.thumbnail_id = f.file_id
	LEFT JOIN unit_groups ug ON p.unit_group_id = ug.id
	LEFT JOIN units un ON ug.base_unit_id = un.id
	WHERE p.id = @p1
	`

	var prd model.Product

	if err := ds.sql.DB.GetContext(ctx, &prd, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.New(errs.NotFound, err, fmt.Sprintf(ErrProductIDNotFound, id))
		}
		return nil, errs.New(errs.Internal, err)
	}

	if prd.ThumbnailPath != nil {
		link := common.Domain + *prd.ThumbnailPath
		prd.ThumbnailLink = &link
	}

	imgList, err := ds.GetProductImages(ctx, id)
	if err != nil {
		return nil, err
	}

	attList, err := ds.GetProductAttachment(ctx, id)
	if err != nil {
		return nil, err
	}

	catList, err := ds.GetCategories(ctx, id)
	if err != nil {
		return nil, err
	}

	prd.ProductImages = imgList
	prd.ProductAttachments = attList
	prd.ProductCategories = catList

	return &prd, nil
}

func (ds *Datastore) GetProducts(ctx context.Context, offset, limit int) ([]*model.Product, error) {
	query := `
	SELECT 
		p.id, p.name, p.description, p.sku, p.information, p.specifications, p.price, p.stock,
		p.unit_group_id, p.thumbnail_id, f.file_path as thumbnail_path,
		un.id AS unit_price_id, un.unit_name AS unit_price_name,
		p.created_at, p.updated_at
	FROM 
		products p
	LEFT JOIN uploaded_files f ON p.thumbnail_id = f.file_id
	LEFT JOIN unit_groups ug ON p.unit_group_id = ug.id
	LEFT JOIN units un ON ug.base_unit_id = un.id
	ORDER BY p.id DESC
	OFFSET @p1 ROWS
	FETCH NEXT @p2 ROWS ONLY
	`

	rows, err := ds.sql.DB.QueryxContext(ctx, query, offset, limit)
	if err != nil {
		return nil, errs.New(errs.Internal, err)
	}

	defer rows.Close()

	var prdList []*model.Product
	for rows.Next() {
		var prd model.Product
		if err := rows.StructScan(&prd); err != nil {
			return nil, errs.New(errs.Internal, err)
		}

		if prd.ThumbnailPath != nil {
			link := common.Domain + *prd.ThumbnailPath
			prd.ThumbnailLink = &link
		}

		prdList = append(prdList, &prd)
	}

	return prdList, nil
}

func (ds *Datastore) GetProductCount(ctx context.Context) (int, error) {
	query := `
	SELECT COUNT(*) FROM products
	`

	var count int
	if err := ds.sql.DB.GetContext(ctx, &count, query); err != nil {
		return -1, errs.New(errs.Internal, err)
	}

	return count, nil
}
