package service

import (
	"context"
	"foxomni/internal/partner/model"
	"foxomni/pkg/pagination"
	"net/url"
)

func (svc *Service) CreatePartnerGroup(ctx context.Context, grp *model.PartnerGroup) (int, error) {
	return svc.ds.CreatePartnerGroup(ctx, grp)
}

func (svc *Service) GetPartnerGroupByID(ctx context.Context, id int) (*model.PartnerGroup, error) {
	return svc.ds.GetPartnerGroupByID(ctx, id)
}

func (svc *Service) GetPartnerGroupPage(ctx context.Context, uv url.Values) (*model.PartnerGroupPage, error) {
	pq, err := pagination.NewWithURLValue(uv)
	if err != nil {
		return nil, err
	}

	count, err := svc.ds.GetPartnerCount(ctx)
	if err != nil {
		return nil, err
	}

	page := &model.PartnerGroupPage{
		Page:       pq.Page,
		Size:       pq.Size,
		TotalCount: count,
		TotalPage:  pq.GetTotalPage(count),
		Hasmore:    pq.GetHasmore(count),
	}

	if count == 0 {
		return page, nil
	}

	grpList, err := svc.ds.GetPartnerGroups(ctx, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, err
	}

	page.Groups = grpList
	return page, nil
}
