package service

import (
	"context"
	"foxomni/internal/category/model"
)

func (svc *Service) CreateCategory(ctx context.Context, cat *model.Category) (int, error) {
	return svc.ds.CreateCategory(ctx, cat)
}

func (svc *Service) UpdateCategory(ctx context.Context, cat *model.Category) error {
	return svc.ds.UpdateCategory(ctx, cat)
}
