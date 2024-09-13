package port

import (
	"encoding/json"
	"foxomni/internal/partner/model"
	"foxomni/pkg/errs"
	"foxomni/pkg/req"
	"foxomni/pkg/resp"
	"net/http"
)

func (hh *HTTPHandler) CreatePartnerGroup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var pnGrp model.PartnerGroup
	if err := json.NewDecoder(r.Body).Decode(&pnGrp); err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	_, err := hh.svc.CreatePartnerGroup(ctx, &pnGrp)
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	resp.WriteJSONMessage(w, "suceess", http.StatusCreated)

}

func (hh *HTTPHandler) GetPartnerGroupByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := req.GetIntFromVars(r, "partner_group_id")
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	grp, err := hh.svc.GetPartnerGroupByID(ctx, id)
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	resp.WriteJSONData(w, grp, http.StatusOK)
}

func (hh *HTTPHandler) GetPartnerGroupPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	page, err := hh.svc.GetPartnerGroupPage(ctx, r.URL.Query())
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	resp.WriteJSONData(w, page, http.StatusOK)
}
