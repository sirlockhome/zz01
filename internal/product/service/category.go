package service

import (
	"context"
	"foxomni/internal/product/model"
)

func (svc *Service) AddCategories(ctx context.Context, catList []*model.ProductCategory, prdID int) error {
	return svc.ds.AddCategories(ctx, catList, prdID)
}

func (svc *Service) RemoveCategory(ctx context.Context, prdID int, catID int) error {
	return svc.ds.RemoveCategory(ctx, prdID, catID)
}
