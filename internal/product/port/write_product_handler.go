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

func (hh *HTTPHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := strconv.Atoi(mux.Vars(r)["product_id"])
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	var prd model.ProductRequest
	if err := json.NewDecoder(r.Body).Decode(&prd); err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	prd.ID = id

	err = hh.svc.UpdateProduct(ctx, &prd)
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	resp.WriteJSONMessage(w, "updated product success", http.StatusOK)
}

func (hh *HTTPHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := strconv.Atoi(mux.Vars(r)["product_id"])
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	err = hh.svc.DeleteProduct(ctx, id)
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	resp.WriteJSONMessage(w, "deleted product success", http.StatusOK)
}

func (hh *HTTPHandler) ChangeThumbnail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := strconv.Atoi(mux.Vars(r)["product_id"])
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	var in map[string]string
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	err = hh.svc.ChangeThumbnail(ctx, id, in["image_id"])
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	resp.WriteJSONMessage(w, "change thumbnail success", http.StatusOK)
}
