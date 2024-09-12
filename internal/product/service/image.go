package service

import (
	"context"
	"foxomni/internal/product/model"
)

func (svc *Service) AddImages(ctx context.Context, imgList []*model.ProductImage, prdID int) error {
	return svc.ds.AddImages(ctx, imgList, prdID)
}

func (svc *Service) RemoveImage(ctx context.Context, prdID int, imgID int) error {
	return svc.ds.RemoveImage(ctx, prdID, imgID)
}
