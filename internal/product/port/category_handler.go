package port

import (
	"encoding/json"
	"foxomni/internal/product/model"
	"foxomni/pkg/errs"
	"foxomni/pkg/resp"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (hh *HTTPHandler) AddCategories(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	prdID, err := strconv.Atoi(mux.Vars(r)["product_id"])
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	var catList []*model.ProductCategory
	if err := json.NewDecoder(r.Body).Decode(&catList); err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	err = hh.svc.AddCategories(ctx, catList, prdID)
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	resp.WriteJSONMessage(w, "add categories to product success", http.StatusCreated)
}

func (hh *HTTPHandler) RemoveCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	prdID, err := strconv.Atoi(mux.Vars(r)["product_id"])
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	catID, err := strconv.Atoi(mux.Vars(r)["category_id"])
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	err = hh.svc.RemoveCategory(ctx, prdID, catID)
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	resp.WriteJSONMessage(w, "remove category from product success", http.StatusOK)
}
