package datastore

import (
	"context"
	"foxomni/internal/order/model"
	"foxomni/pkg/errs"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func (ds *Datastore) CreateOrder(ctx context.Context, order *model.Order) (int, string, error) {
	query := `
	INSERT INTO orders
	(
		user_id, buyer_note, order_tracking_id
	) OUTPUT INSERTED.id
	VALUES
	(
		:user_id, :buyer_note, :order_tracking_id
	)
	`

	order.OrderTrackingID = generateTrackingID()

	var idNewOrder int
	err := ds.sql.InTx(ctx, func(tx *sqlx.Tx) error {
		rows, err := tx.NamedQuery(query, order)
		if err != nil {
			return errs.New(errs.Internal, err)
		}

		defer rows.Close()
		if rows.Next() {
			if err := rows.Scan(&idNewOrder); err != nil {
				return errs.New(errs.Internal, err)
			}
		} else {
			return errs.New(errs.Internal, rows.Err())
		}

		err = createOrderItems(ctx, tx, order.OrderItems, idNewOrder)
		if err != nil {
			return err
		}

		err = createAddress(ctx, tx, idNewOrder, order.PartnerAddressID, "shipping")
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return -1, "", err
	}

	return idNewOrder, order.OrderTrackingID, nil
}

func createOrderItems(ctx context.Context, tx *sqlx.Tx, items []*model.OrderItem, orderID int) error {
	query := `
	INSERT INTO order_items
	(
		order_id, product_id, product_name, product_sku,
		quantity, price, unit_id, unit_name
	)
	SELECT 
		:order_id, :product_id, p.name, p.sku,
		:quantity, p.price, :unit_id, u.unit_name
	FROM 
		products p
	LEFT JOIN units u ON p.unit_price_id = u.id
	WHERE p.id = :product_id
	`

	for _, item := range items {
		item.OrderID = orderID
		_, err := tx.NamedExecContext(ctx, query, item)
		if err != nil {
			return errs.New(errs.Internal, err)
		}
	}

	return nil
}

func createAddress(ctx context.Context, tx *sqlx.Tx, orderID int, addrID int, addrType string) error {
	query := `
	INSERT INTO order_addresses
	(
		order_id, country, city, district, ward, street, house_number,
		address_line1, address_line2, address_type
	)
	SELECT
		@p1, country, city, district, ward, street, house_number,
		address_line1, address_line2, @p2 
	FROM partner_addresses WHERE id = @p3
	`

	_, err := tx.ExecContext(ctx, query, orderID, addrType, addrID)
	if err != nil {
		return errs.New(errs.Internal, err)
	}

	return nil
}

func generateTrackingID() string {
	return uuid.New().String()
}
