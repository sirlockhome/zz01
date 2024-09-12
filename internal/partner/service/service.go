package service

import (
	"context"
	"foxomni/internal/partner/datastore"
	"foxomni/internal/partner/model"
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

func (svc *Service) CreatePartner(ctx context.Context, pn *model.PartnerRequest) (int, error) {
	return svc.ds.CreatePartner(ctx, pn)
}

func (svc *Service) UpdatePartner(ctx context.Context, pn *model.PartnerRequest) error {
	return svc.ds.UpdatePartner(ctx, pn)
}

func (svc *Service) DeletePartner(ctx context.Context, id int) error {
	return svc.ds.DeletePartner(ctx, id)
}

func (svc *Service) GetPartnerByID(ctx context.Context, id int) (*model.Partner, error) {
	return svc.ds.GetPartnerByID(ctx, id)
}

func (svc *Service) GetPartnerPage(ctx context.Context, uv url.Values) (*model.PartnerPage, error) {
	pq, err := pagination.NewWithURLValue(uv)
	if err != nil {
		return nil, err
	}

	count, err := svc.ds.GetPartnerCount(ctx)
	if err != nil {
		return nil, err
	}

	page := &model.PartnerPage{
		Page:       pq.Page,
		Size:       pq.Size,
		TotalCount: count,
		TotalPage:  pq.GetTotalPage(count),
		Hasmore:    pq.GetHasmore(count),
	}

	if count == 0 {
		return page, nil
	}

	pnList, err := svc.ds.GetPartners(ctx, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, err
	}

	page.Partners = pnList

	return page, nil
}
