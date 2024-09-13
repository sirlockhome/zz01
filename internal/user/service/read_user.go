package service

import (
	"context"
	"foxomni/internal/user/model"
	"foxomni/pkg/errs"
)

func (svc *Service) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	return svc.ds.GetUserByID(ctx, id)
}

func (svc *Service) GetMe(ctx context.Context) (*model.User, error) {
	id, ok := ctx.Value("user_id").(int)
	if !ok {
		return nil, errs.New(errs.Unauthorized)
	}

	return svc.ds.GetUserByID(ctx, id)
}
