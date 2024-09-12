package port

import (
	"encoding/json"
	"fmt"
	"foxomni/internal/partner/model"
	"foxomni/internal/partner/service"
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
	partner := r.PathPrefix("/partners").Subrouter()
	partner.Use(mwf...)

	partner.HandleFunc("", hh.CreatePartner).Methods(http.MethodPost)
	partner.HandleFunc("", hh.GetPartnerPage).Methods(http.MethodGet)
	partner.HandleFunc("/{partner_id}", hh.GetPartnerByID).Methods(http.MethodGet)
	partner.HandleFunc("/{partner_id}", hh.UpdatePartner).Methods(http.MethodPut)
	partner.HandleFunc("/{partner_id}", hh.DeletePartner).Methods(http.MethodDelete)

	partner.HandleFunc("/{partner_id}/addresses", hh.GetAddresses).Methods(http.MethodPost)
	partner.HandleFunc("/{partner_id}/addresses", hh.GetAddresses).Methods(http.MethodGet)
	partner.HandleFunc("/{partner_id}/addresses/{address_id}", hh.RemoveAddress).Methods(http.MethodDelete)
}

func (hh *HTTPHandler) GetPartnerByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := req.GetIntFromVars(r, "partner_id")
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	partner, err := hh.svc.GetPartnerByID(ctx, id)
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	resp.WriteJSONData(w, partner, http.StatusOK)
}

func (hh *HTTPHandler) GetPartnerPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	page, err := hh.svc.GetPartnerPage(ctx, r.URL.Query())
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	resp.WriteJSONData(w, page, http.StatusOK)
}

func (hh *HTTPHandler) CreatePartner(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var pn model.PartnerRequest
	if err := json.NewDecoder(r.Body).Decode(&pn); err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	idNewPn, err := hh.svc.CreatePartner(ctx, &pn)
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	resp.WriteJSONMessage(w, fmt.Sprintf("new partner with id `%d`", idNewPn), http.StatusCreated)
}

func (hh *HTTPHandler) UpdatePartner(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := req.GetIntFromVars(r, "partner_id")
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	var pn model.PartnerRequest
	if err := json.NewDecoder(r.Body).Decode(&pn); err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	pn.ID = id
	err = hh.svc.UpdatePartner(ctx, &pn)
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
	}

	resp.WriteJSONMessage(w, "updated partner success", http.StatusOK)
}

func (hh *HTTPHandler) DeletePartner(w http.ResponseWriter, r *http.Request) {
}
