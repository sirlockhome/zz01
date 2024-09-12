package port

import (
	"foxomni/pkg/errs"
	"foxomni/pkg/req"
	"foxomni/pkg/resp"
	"net/http"
)

func (hh *HTTPHandler) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := req.GetIntFromVars(r, "category_id")
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	cat, err := hh.svc.GetCategoryByID(ctx, id)
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	resp.WriteJSONData(w, cat, http.StatusOK)
}

func (hh *HTTPHandler) GetCategoryPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	page, err := hh.svc.GetCategoryPage(ctx, r.URL.Query())
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	resp.WriteJSONData(w, page, http.StatusOK)
}
