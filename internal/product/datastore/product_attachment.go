package datastore

import (
	"context"
	"foxomni/internal/product/model"
	"foxomni/pkg/errs"

	"github.com/jmoiron/sqlx"
)

func (ds *Datastore) GetProductAttachment(ctx context.Context, prdID int) ([]*model.ProductAttachment, error) {
	query := `
	SELECT
		a.file_id, a.description, f.file_name, f.file_path
	FROM 
		product_attachments a
	LEFT JOIN uploaded_files f ON a.file_id = f.file_id
	WHERE product_id = @p1
	`

	rows, err := ds.sql.DB.QueryxContext(ctx, query, prdID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var attList []*model.ProductAttachment
	for rows.Next() {
		var att model.ProductAttachment
		if err := rows.StructScan(&att); err != nil {
			return nil, err
		}

		attList = append(attList, &att)
	}

	return attList, nil
}

func (ds *Datastore) AddAttachments(ctx context.Context, attList []*model.ProductAttachment, prdID int) error {
	return ds.sql.InTx(ctx, func(tx *sqlx.Tx) error {
		return addAttchmentsWithTx(ctx, tx, attList, prdID)
	})
}

func (ds *Datastore) RemoveAttachment(ctx context.Context, prdID int, attID int) error {
	query := `
	DELETE product_attachments WHERE product_id = @p1 AND file_id = @p2
	`

	_, err := ds.sql.DB.ExecContext(ctx, query, prdID, attID)
	if err != nil {
		return errs.New(errs.Internal, err)
	}

	return nil
}

func addAttchmentsWithTx(ctx context.Context, tx *sqlx.Tx, attList []*model.ProductAttachment, prdID int) error {
	query := `
	INSERT INTO product_attachments
	(
		file_id, product_id, description
	)
	VALUES
	(
		:file_id, :product_id, :description
	)
	`

	for _, att := range attList {
		att.ProductID = prdID
		_, err := tx.NamedExecContext(ctx, query, &att)
		if err != nil {
			return errs.New(errs.Internal, err)
		}
	}

	return nil
}
