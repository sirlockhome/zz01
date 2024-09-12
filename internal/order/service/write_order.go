package service

import "context"

func (svc *Service) ChangeOrderStatus(ctx context.Context, id int, status string) error {
	return svc.ds.ChangeStatusOrder(ctx, id, status)
}
