package port

import (
	"encoding/json"
	"foxomni/internal/order/model"
	"foxomni/pkg/errs"
	"foxomni/pkg/resp"
	"net/http"
)

func (hh *HTTPHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var order model.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	orderResp, err := hh.svc.CreateOrder(ctx, &order)
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	resp.WriteJSONData(w, orderResp, http.StatusCreated)
}
