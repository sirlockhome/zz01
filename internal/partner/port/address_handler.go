package port

import (
	"encoding/json"
	"foxomni/internal/partner/model"
	"foxomni/pkg/errs"
	"foxomni/pkg/req"
	"foxomni/pkg/resp"
	"net/http"
)

func (hh *HTTPHandler) GetAddresses(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	pnID, err := req.GetIntFromVars(r, "partner_id")
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	addrList, err := hh.svc.GetAddresses(ctx, pnID)
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	resp.WriteJSONData(w, addrList, http.StatusOK)
}

func (hh *HTTPHandler) AddAddress(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	pnID, err := req.GetIntFromVars(r, "partner_id")
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	var addr model.Address
	if err := json.NewDecoder(r.Body).Decode(&addr); err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	addr.PartnerID = pnID

	err = hh.svc.AddAddress(ctx, &addr)
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	resp.WriteJSONMessage(w, "add new address to partner success", http.StatusCreated)
}

func (hh *HTTPHandler) RemoveAddress(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	pnID, err := req.GetIntFromVars(r, "partner_id")
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	addrID, err := req.GetIntFromVars(r, "address_id")
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	err = hh.svc.RemoveAddress(ctx, addrID, pnID)
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	resp.WriteJSONMessage(w, "remove address from partner success", http.StatusOK)
}
