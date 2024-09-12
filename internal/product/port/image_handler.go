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

func (hh *HTTPHandler) AddImages(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	prdID, err := strconv.Atoi(mux.Vars(r)["product_id"])
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	var imgList []*model.ProductImage
	if err := json.NewDecoder(r.Body).Decode(&imgList); err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	err = hh.svc.AddImages(ctx, imgList, prdID)
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	resp.WriteJSONMessage(w, "add images to product success", http.StatusCreated)
}

func (hh *HTTPHandler) RemoveImage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	prdID, err := strconv.Atoi(mux.Vars(r)["product_id"])
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	imgID, err := strconv.Atoi(mux.Vars(r)["image_id"])
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	err = hh.svc.RemoveImage(ctx, prdID, imgID)
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	resp.WriteJSONMessage(w, "remove image from product success", http.StatusOK)
}
