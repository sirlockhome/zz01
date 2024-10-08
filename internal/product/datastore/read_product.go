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

	unitGrp, _ := ds.getUnitGroup(ctx, prd.ID)

	prd.ProductImages = imgList
	prd.ProductAttachments = attList
	prd.ProductCategories = catList
	prd.UnitGroup = unitGrp

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

		unitGrp, err := ds.getUnitGroup(ctx, prd.ID)
		if err != nil {
			return nil, err
		}
		prd.UnitGroup = unitGrp

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

func (ds *Datastore) getUnitGroup(ctx context.Context, prdID int) (*model.UnitGroup, error) {
	query := `
	SELECT 
		ug.id, ug.unit_group_name, ug.base_unit_id, u.unit_name AS base_unit_name, ug.symbol,
		ug.description 
	FROM unit_groups ug 
	LEFT JOIN products p ON p.unit_group_id = ug.id
	LEFT JOIN units u ON ug.base_unit_id = u.id
	WHERE p.id = @p1
	`

	var grp model.UnitGroup
	if err := ds.sql.DB.GetContext(ctx, &grp, query, prdID); err != nil {
		return nil, err
	}

	convList, err := ds.getUnitConversions(ctx, grp.ID)
	if err != nil {
		return nil, err
	}

	grp.UnitConversions = convList
	return &grp, nil
}

func (ds *Datastore) getUnitConversions(ctx context.Context, grpID int) ([]*model.UnitConversion, error) {
	query := `
	SELECT 
		ucv.to_unit_id, u.unit_name AS to_unit_name,
		ucv.base_qty, ucv.alt_qty
	FROM unit_conversions ucv
	LEFT JOIN unit_groups ug ON ucv.unit_group_id = ug.id
	LEFT JOIN units u ON u.id = ucv.to_unit_id
	WHERE ug.id = @p1
	`

	rows, err := ds.sql.DB.QueryxContext(ctx, query, grpID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, errs.New(errs.Internal, err)
	}

	defer rows.Close()

	var convList []*model.UnitConversion
	for rows.Next() {
		var conv model.UnitConversion
		if err := rows.StructScan(&conv); err != nil {
			return nil, errs.New(errs.Internal, err)
		}

		convList = append(convList, &conv)
	}

	return convList, nil
}
