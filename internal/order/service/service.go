package service

import (
	"context"
	"foxomni/internal/order/datastore"
	"foxomni/internal/order/model"
	"foxomni/pkg/pagination"
	"net/url"
)

type Service struct {
	ds *datastore.Datastore
}

func NewService(ds *datastore.Datastore) *Service {
	return &Service{
		ds: ds,
	}
}

func (svc *Service) GetOrderByID(ctx context.Context, id int) (*model.Order, error) {
	return svc.ds.GetOrderByID(ctx, id)
}

func (svc *Service) GetOrderPage(ctx context.Context, uv url.Values) (*model.OrderPage, error) {
	pq, err := pagination.NewWithURLValue(uv)
	if err != nil {
		return nil, err
	}

	count, err := svc.ds.GetOrderCount(ctx)
	if err != nil {
		return nil, err
	}

	page := &model.OrderPage{
		Page:      pq.Page,
		Size:      pq.Size,
		TotalPage: pq.GetTotalPage(count),
		TotalSize: count,
		Hasmore:   pq.GetHasmore(count),
	}

	if count == 0 {
		return page, nil
	}

	orderList, err := svc.ds.GetOrders(ctx, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, err
	}

	page.Orders = orderList
	return page, nil
}
