package service

import (
	"context"
	"foxomni/internal/category/model"
	"foxomni/pkg/pagination"
	"net/url"
)

func (svc *Service) GetCategoryByID(ctx context.Context, id int) (*model.Category, error) {
	return svc.ds.GetCategoryByID(ctx, id)
}

func (svc *Service) GetCategoryPage(ctx context.Context, uv url.Values) (*model.CategoryPage, error) {
	pq, err := pagination.NewWithURLValue(uv)
	if err != nil {
		return nil, err
	}

	count, err := svc.ds.GetCategoryCount(ctx)
	if err != nil {
		return nil, err
	}

	page := &model.CategoryPage{
		Page:       pq.Page,
		Size:       pq.Size,
		TotalPage:  pq.GetTotalPage(count),
		TotalCount: count,
		Hasmore:    pq.GetHasmore(count),
	}

	if count == 0 {
		return page, nil
	}

	catList, err := svc.ds.GetCategories(ctx, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, err
	}

	page.Categories = catList

	return page, nil
}
