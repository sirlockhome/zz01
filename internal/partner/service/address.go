package service

import (
	"context"
	"foxomni/internal/partner/model"
)

func (svc *Service) GetAddresses(ctx context.Context, pnID int) ([]*model.Address, error) {
	return svc.ds.GetAddresses(ctx, pnID)
}

func (svc *Service) AddAddress(ctx context.Context, addr *model.Address) error {
	return svc.ds.AddAddress(ctx, addr)
}

func (svc *Service) RemoveAddress(ctx context.Context, addrID int, pnID int) error {
	return svc.ds.RemoveAddress(ctx, addrID, pnID)
}
