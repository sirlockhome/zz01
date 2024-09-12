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

func (hh *HTTPHandler) AddAttachments(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	prdID, err := strconv.Atoi(mux.Vars(r)["product_id"])
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	var attList []*model.ProductAttachment
	if err := json.NewDecoder(r.Body).Decode(&attList); err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	err = hh.svc.AddAttachments(ctx, attList, prdID)
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	resp.WriteJSONMessage(w, "add attachments to product success", http.StatusCreated)
}

func (hh *HTTPHandler) RemoveAttachment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	prdID, err := strconv.Atoi(mux.Vars(r)["product_id"])
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	attID, err := strconv.Atoi(mux.Vars(r)["attachment_id"])
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	err = hh.svc.RemoveAttachment(ctx, prdID, attID)
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	resp.WriteJSONMessage(w, "remove attachment from product success", http.StatusOK)
}
