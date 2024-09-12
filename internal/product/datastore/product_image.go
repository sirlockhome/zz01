package datastore

import (
	"context"
	"foxomni/internal/common"
	"foxomni/internal/product/model"
	"foxomni/pkg/errs"

	"github.com/jmoiron/sqlx"
)

func (ds *Datastore) GetProductImages(ctx context.Context, prdID int) ([]*model.ProductImage, error) {
	query := `
	SELECT 
		i.image_id, f.file_path AS image_path, f.file_name AS image_name 
	FROM 
		product_images i
	LEFT JOIN uploaded_files f ON i.image_id = f.file_id
	WHERE i.product_id = @p1
	`

	rows, err := ds.sql.DB.QueryxContext(ctx, query, prdID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var imgList []*model.ProductImage
	for rows.Next() {
		var img model.ProductImage
		if err := rows.StructScan(&img); err != nil {
			return nil, err
		}

		if img.ImagePath != nil {
			link := common.Domain + *img.ImagePath
			img.ImageLink = &link
		}

		imgList = append(imgList, &img)
	}

	return imgList, nil
}

func (ds *Datastore) AddImages(ctx context.Context, imgList []*model.ProductImage, prdID int) error {
	return ds.sql.InTx(ctx, func(tx *sqlx.Tx) error {
		return addImagesWithTx(ctx, tx, imgList, prdID)
	})
}

func (ds *Datastore) RemoveImage(ctx context.Context, prdID int, imgID int) error {
	query := `
	DELETE product_images WHERE product_id = @p1 AND image_id = @p2
	`

	_, err := ds.sql.DB.ExecContext(ctx, query, prdID, imgID)
	if err != nil {
		return errs.New(errs.Internal, err)
	}

	return nil
}

func addImagesWithTx(ctx context.Context, tx *sqlx.Tx, imgList []*model.ProductImage, prdID int) error {
	query := `
	INSERT INTO product_images
	(
		image_id, product_id
	)
	VALUES
	(
		:image_id, :product_id
	)
	`

	for _, img := range imgList {
		img.ProductID = prdID
		_, err := tx.NamedExecContext(ctx, query, &img)
		if err != nil {
			return errs.New(errs.Internal, err)
		}
	}

	return nil
}
