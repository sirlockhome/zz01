package port

import (
	"encoding/json"
	"foxomni/internal/unit/model"
	"foxomni/pkg/errs"
	"foxomni/pkg/req"
	"foxomni/pkg/resp"
	"net/http"
)

func (hh *HTTPHandler) GetUnitGroupByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := req.GetIntFromVars(r, "unit_group_id")
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	grp, err := hh.svc.GetUnitGroupByID(ctx, id)
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	resp.WriteJSONData(w, grp, http.StatusOK)
}

func (hh *HTTPHandler) GetUnitGroupPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	page, err := hh.svc.GetUnitGroupPage(ctx, r.URL.Query())
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	resp.WriteJSONData(w, page, http.StatusOK)
}

func (hh *HTTPHandler) CreateUnitGroup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var grp model.UnitGroup
	if err := json.NewDecoder(r.Body).Decode(&grp); err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	_, err := hh.svc.CreateUnitGroup(ctx, &grp)
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	resp.WriteJSONMessage(w, "success", http.StatusCreated)
}

func (hh *HTTPHandler) AddUnitConversions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	grpID, err := req.GetIntFromVars(r, "unit_group_id")
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	var convList []*model.UnitConversion
	if err := json.NewDecoder(r.Body).Decode(&convList); err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	err = hh.svc.AddUnitConversions(ctx, convList, grpID)
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	resp.WriteJSONMessage(w, "success", http.StatusCreated)
}
