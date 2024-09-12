package service

import (
	"context"
	"foxomni/internal/unit/model"
	"foxomni/pkg/pagination"
	"net/url"
)

func (svc *Service) GetUnitGroupByID(ctx context.Context, id int) (*model.UnitGroup, error) {
	return svc.ds.GetUnitGroupByID(ctx, id)
}

func (svc *Service) GetUnitGroupPage(ctx context.Context, uv url.Values) (*model.UnitGroupPage, error) {
	pq, err := pagination.NewWithURLValue(uv)
	if err != nil {
		return nil, err
	}

	count, err := svc.ds.GetUnitGroupCount(ctx)
	if err != nil {
		return nil, err
	}

	page := &model.UnitGroupPage{
		Page:       pq.Page,
		Size:       pq.Size,
		TotalPage:  pq.GetTotalPage(count),
		TotalCount: count,
		Hasmore:    pq.GetHasmore(count),
	}

	if count == 0 {
		return page, nil
	}

	grpList, err := svc.ds.GetUnitGroups(ctx, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, err
	}

	page.UnitGroups = grpList
	return page, nil
}

func (svc *Service) CreateUnitGroup(ctx context.Context, grp *model.UnitGroup) (int, error) {
	return svc.ds.CreateUnitGroup(ctx, grp)
}

func (svc *Service) AddUnitConversions(ctx context.Context, convList []*model.UnitConversion, grpID int) error {
	return svc.ds.AddConversions(ctx, convList, grpID)
}
