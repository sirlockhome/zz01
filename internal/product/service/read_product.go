package service

import (
	"context"
	"foxomni/internal/product/model"
	"foxomni/pkg/pagination"
	"net/url"
)

func (svc *Service) GetProductByID(ctx context.Context, id int) (*model.Product, error) {
	return svc.ds.GetProductByID(ctx, id)
}

func (svc *Service) GetProductPage(ctx context.Context, uv url.Values) (*model.ProductPage, error) {
	pq, err := pagination.NewWithURLValue(uv)
	if err != nil {
		return nil, err
	}

	total, err := svc.ds.GetProductCount(ctx)
	if err != nil {
		return nil, err
	}

	page := &model.ProductPage{
		Page:       pq.Page,
		Size:       pq.Size,
		TotalPage:  pq.GetTotalPage(total),
		TotalCount: total,
		Hasmore:    pq.GetHasmore(total),
	}

	if total == 0 {
		return page, nil
	}

	prdList, err := svc.ds.GetProducts(ctx, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, err
	}

	page.Products = prdList
	return page, nil
}
