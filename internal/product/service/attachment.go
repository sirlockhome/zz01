package service

import (
	"context"
	"foxomni/internal/product/model"
)

func (svc *Service) AddAttachments(ctx context.Context, attList []*model.ProductAttachment, prdID int) error {
	return svc.ds.AddAttachments(ctx, attList, prdID)
}

func (svc *Service) RemoveAttachment(ctx context.Context, prdID int, attID int) error {
	return svc.ds.RemoveAttachment(ctx, prdID, attID)
}
