package pagination

import (
	"foxomni/pkg/errs"
	"math"
	"net/url"
	"strconv"
)

type PaginationQuery struct {
	Page int
	Size int
}

const defaultSize = 100

const (
	ErrSizeIsNumber = "`size` is a number"
	ErrPageIsNumber = "`page` is a number"
)

func NewWithURLValue(uv url.Values) (*PaginationQuery, error) {
	pq := &PaginationQuery{}

	if err := pq.setPage(uv.Get("page")); err != nil {
		return nil, err
	}

	if err := pq.setSize(uv.Get("size")); err != nil {
		return nil, err
	}

	return pq, nil
}

func (pg *PaginationQuery) setSize(size string) error {
	if size == "" {
		pg.Size = defaultSize
		return nil
	}

	n, err := strconv.Atoi(size)
	if err != nil {
		return errs.New(errs.InvalidRequest, ErrSizeIsNumber)
	}

	pg.Size = n

	if n == 0 {
		pg.Size = defaultSize
	}

	return nil
}

func (pq *PaginationQuery) setPage(page string) error {
	if page == "" {
		pq.Page = 1
		return nil
	}

	n, err := strconv.Atoi(page)
	if err != nil {
		return errs.New(errs.InvalidRequest, ErrPageIsNumber)
	}

	pq.Page = n
	if n == 0 {
		pq.Page = 1
	}

	return nil
}

func (pq *PaginationQuery) GetOffset() int {
	if pq.Page == 0 {
		return 0
	}

	return (pq.Page - 1) * pq.Size
}

func (pq *PaginationQuery) GetLimit() int {
	return pq.Size
}

func (pq *PaginationQuery) GetTotalPage(totalCount int) int {
	d := float64(totalCount) / float64(pq.Size)
	return int(math.Ceil(d))
}

func (pq *PaginationQuery) GetHasmore(totalCount int) bool {
	return float64(pq.Page) < float64(totalCount/pq.Size)
}
