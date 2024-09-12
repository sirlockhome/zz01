package service

import (
	"context"
	"foxomni/internal/unit/datastore"
	"foxomni/internal/unit/model"
	"foxomni/pkg/pagination"
	"net/url"
)

type Service struct {
	ds *datastore.Datastore
}

func NewService(ds *datastore.Datastore) *Service {
	return &Service{
		ds: ds,
	}
}

func (svc *Service) CreateUnit(ctx context.Context, unit *model.Unit) (int, error) {
	return svc.ds.CreateUnit(ctx, unit)
}

func (svc *Service) UpdateUnit(ctx context.Context, unit *model.Unit) error {
	return svc.ds.UpdateUnit(ctx, unit)
}

func (svc *Service) DeleteUnit(ctx context.Context, id int) error {
	return svc.ds.DeleteUnit(ctx, id)
}

func (svc *Service) GetUnitByID(ctx context.Context, id int) (*model.Unit, error) {
	return svc.ds.GetUnitByID(ctx, id)
}

func (svc *Service) GetUnitPage(ctx context.Context, uv url.Values) (*model.UnitPage, error) {
	pq, err := pagination.NewWithURLValue(uv)
	if err != nil {
		return nil, err
	}

	count, err := svc.ds.GetUnitCount(ctx)
	if err != nil {
		return nil, err
	}

	page := &model.UnitPage{
		Page:       pq.Page,
		Size:       pq.Size,
		TotalCount: count,
		TotalPage:  pq.GetTotalPage(count),
		Hasmore:    pq.GetHasmore(count),
	}

	if count == 0 {
		return page, nil
	}

	unitList, err := svc.ds.GetUnits(ctx, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, err
	}

	page.Units = unitList
	return page, nil
}
