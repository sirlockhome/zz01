package port

import (
	"foxomni/pkg/errs"
	"foxomni/pkg/req"
	"foxomni/pkg/resp"
	"net/http"
)

func (hh *HTTPHandler) ChangeOrderStatus(status string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		id, err := req.GetIntFromVars(r, "order_id")
		if err != nil {
			errs.HTTPErrorResponse(ctx, w, err)
			return
		}

		err = hh.svc.ChangeOrderStatus(ctx, id, status)
		if err != nil {
			errs.HTTPErrorResponse(ctx, w, err)
			return
		}

		resp.WriteJSONMessage(w, "change status success", http.StatusOK)
	}
}
