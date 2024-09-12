package port

import (
	"foxomni/internal/order/service"
	"foxomni/pkg/errs"
	"foxomni/pkg/req"
	"foxomni/pkg/resp"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPHandler struct {
	svc *service.Service
}

func NewHTTPHandler(svc *service.Service) *HTTPHandler {
	return &HTTPHandler{
		svc: svc,
	}
}

func (hh *HTTPHandler) Routes(r *mux.Router, mwf ...mux.MiddlewareFunc) {
	order := r.PathPrefix("/orders").Subrouter()
	order.Use(mwf...)

	order.HandleFunc("", hh.CreateOrder).Methods(http.MethodPost)
	order.HandleFunc("", hh.GetOrderPage).Methods(http.MethodGet)
	order.HandleFunc("/{order_id}", hh.GetOrderByID).Methods(http.MethodGet)

	order.HandleFunc("/{order_id}/confirm", hh.ChangeOrderStatus("confirmed")).Methods(http.MethodPost)
	order.HandleFunc("/{order_id}/reject", hh.ChangeOrderStatus("rejected")).Methods(http.MethodPost)
	order.HandleFunc("/{order_id}/cancel", hh.ChangeOrderStatus("canceled")).Methods(http.MethodPost)
}

func (hh *HTTPHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := req.GetIntFromVars(r, "order_id")
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	order, err := hh.svc.GetOrderByID(ctx, id)
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	resp.WriteJSONData(w, order, http.StatusOK)
}

func (hh *HTTPHandler) GetOrderPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	page, err := hh.svc.GetOrderPage(ctx, r.URL.Query())
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	resp.WriteJSONData(w, page, http.StatusOK)
}
