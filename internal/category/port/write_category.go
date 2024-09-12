package port

import (
	"encoding/json"
	"foxomni/internal/category/model"
	"foxomni/pkg/errs"
	"foxomni/pkg/req"
	"foxomni/pkg/resp"
	"net/http"
)

func (hh *HTTPHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var cat model.Category
	if err := json.NewDecoder(r.Body).Decode(&cat); err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	_, err := hh.svc.CreateCategory(ctx, &cat)
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	resp.WriteJSONMessage(w, "add new category success", http.StatusCreated)
}

func (hh *HTTPHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := req.GetIntFromVars(r, "category_id")
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	var cat model.Category
	if err := json.NewDecoder(r.Body).Decode(&cat); err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}
	cat.ID = id

	err = hh.svc.UpdateCategory(ctx, &cat)
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	resp.WriteJSONMessage(w, "update category success", http.StatusOK)
}
