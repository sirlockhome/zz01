package service

import (
	"context"
	"foxomni/internal/order/model"
	"foxomni/pkg/errs"
)

type newOrderResp struct {
	NewOrderID         int    `json:"new_order_id"`
	NewOrderTrackingID string `json:"new_order_tracking_id"`
}

func (svc *Service) CreateOrder(ctx context.Context, order *model.Order) (*newOrderResp, error) {
	userID, ok := ctx.Value("user_id").(int)
	if !ok {
		return nil, errs.New(errs.Internal)
	}

	order.UserID = userID

	id, trkid, err := svc.ds.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	return &newOrderResp{
		NewOrderID:         id,
		NewOrderTrackingID: trkid,
	}, nil
}
