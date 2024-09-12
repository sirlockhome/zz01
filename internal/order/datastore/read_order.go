package datastore

import (
	"context"
	"database/sql"
	"errors"
	"foxomni/internal/common"
	"foxomni/internal/order/model"
	"foxomni/pkg/errs"
)

func (ds *Datastore) GetOrderByID(ctx context.Context, id int) (*model.Order, error) {
	query := `
	SELECT
		o.id, o.user_id, o.order_date, o.status, o.buyer_note, o.order_tracking_id
	FROM 
		orders o
	WHERE o.id = @p1
	`

	var order model.Order
	if err := ds.sql.DB.GetContext(ctx, &order, query, id); err != nil {
		return nil, errs.New(errs.Internal, err)
	}

	items, totalAmount, err := ds.getOrderItems(ctx, id)
	if err != nil {
		return nil, errs.New(errs.Internal, err)
	}

	shippingAddr, err := ds.getShippingAddress(ctx, order.ID)
	if err != nil {
		return nil, err
	}
	bb, err := ds.getBusinessBuyer(ctx, order.ID)
	if err != nil {
		return nil, err
	}

	order.BusinessBuyer = bb
	order.ShippingAddress = shippingAddr

	order.OrderItems = items
	order.TotalAmount = totalAmount

	return &order, nil
}

func (ds *Datastore) GetOrders(ctx context.Context, offset, limit int) ([]*model.Order, error) {
	query := `
	SELECT
		o.id, o.user_id, o.order_date, o.status, o.buyer_note, o.order_tracking_id
	FROM 
		orders o
	ORDER BY o.id DESC
	OFFSET @p1 ROWS
	FETCH NEXT @p2 ROWS ONLY
	`

	rows, err := ds.sql.DB.QueryxContext(ctx, query, offset, limit)
	if err != nil {
		return nil, errs.New(errs.Internal, err)
	}

	defer rows.Close()

	var orderList []*model.Order
	for rows.Next() {
		var order model.Order
		if err := rows.StructScan(&order); err != nil {
			return nil, errs.New(errs.Internal, err)
		}

		shippingAddr, err := ds.getShippingAddress(ctx, order.ID)
		if err != nil {
			return nil, err
		}
		bb, err := ds.getBusinessBuyer(ctx, order.ID)
		if err != nil {
			return nil, err
		}

		order.BusinessBuyer = bb
		order.ShippingAddress = shippingAddr

		orderList = append(orderList, &order)
	}

	return orderList, nil
}

func (ds *Datastore) GetOrderCount(ctx context.Context) (int, error) {
	query := `SELECT COUNT(*) FROM orders`

	var count int
	if err := ds.sql.DB.GetContext(ctx, &count, query); err != nil {
		return -1, errs.New(errs.Internal, err)
	}

	return count, nil
}

func (ds *Datastore) getOrderItems(ctx context.Context, orderID int) ([]*model.OrderItem, float64, error) {
	query := `
	SELECT
		oi.id, oi.order_id, oi.product_id, oi.product_name, oi.product_sku, oi.unit_id,
		oi.quantity, oi.price,
		oi.unit_name, p.thumbnail_id, uf.file_path AS thumbnail_path
	FROM
		order_items oi
	LEFT JOIN products p ON p.id = oi.product_id
	LEFT JOIN uploaded_files uf ON p.thumbnail_id = uf.file_id
	WHERE oi.order_id = @p1
	`

	rows, err := ds.sql.DB.QueryxContext(ctx, query, orderID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, 0, nil
		}
		return nil, -1, errs.New(errs.Internal, err)
	}

	defer rows.Close()

	var totalAmount float64 = 0
	var itemList []*model.OrderItem
	for rows.Next() {
		var item model.OrderItem
		if err := rows.StructScan(&item); err != nil {
			return nil, -1, errs.New(errs.Internal, err)
		}

		if item.ThumbnailPath != nil {
			link := common.Domain + *item.ThumbnailPath
			item.ThumbnailLink = &link
		}

		totalAmount += item.Price * float64(item.Quantity)
		item.TotalAmount = totalAmount

		itemList = append(itemList, &item)
	}

	return itemList, totalAmount, nil
}

func (ds *Datastore) getShippingAddress(ctx context.Context, orderID int) (*model.OrderAddress, error) {
	query := `
	SELECT 
		order_id, country, city, district, ward, street, house_number, address_line1, address_line2
	FROM order_addresses
	WHERE order_id = @p1 AND address_type = 'shipping'
	`

	var address model.OrderAddress
	if err := ds.sql.DB.GetContext(ctx, &address, query, orderID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, errs.New(errs.Internal, err)
	}

	return &address, nil
}

func (ds *Datastore) getBusinessBuyer(ctx context.Context, orderID int) (*model.BusinessBuyer, error) {
	query := `
	SELECT
		business_name, tax_id, contact_name, contact_phone, contact_position
	FROM 
		business_buyers
	WHERE order_id = @p1
	`

	var bb model.BusinessBuyer
	if err := ds.sql.DB.GetContext(ctx, &bb, query, orderID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, errs.New(errs.Internal, err)
	}

	return &bb, nil
}
