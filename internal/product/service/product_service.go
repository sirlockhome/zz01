package service

import "context"

func (svc *Service) ChangeThumbnail(ctx context.Context, prdID int, imgID string) error {
	return svc.ds.ChangeThumbnail(ctx, prdID, imgID)
}

func (svc *Service) ChangeActive(ctx context.Context, prdID int, isActive bool) error {
	return svc.ds.ChangeActive(ctx, prdID, isActive)
}
