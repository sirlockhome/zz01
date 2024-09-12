package service

import (
	"context"
	"foxomni/internal/product/model"
)

func (svc *Service) CreateProduct(ctx context.Context, prd *model.ProductRequest) (int, error) {
	return svc.ds.CreateProduct(ctx, prd)
}

func (svc *Service) UpdateProduct(ctx context.Context, prd *model.ProductRequest) error {
	return svc.ds.UpdateProduct(ctx, prd)
}

func (svc *Service) DeleteProduct(ctx context.Context, id int) error {
	return svc.ds.DeleteProduct(ctx, id)
}
