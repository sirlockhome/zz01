package port

import (
	"encoding/json"
	"fmt"
	"foxomni/internal/unit/model"
	"foxomni/internal/unit/service"
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
	unit := r.PathPrefix("/units").Subrouter()
	unitGroup := r.PathPrefix("/unit-groups").Subrouter()
	unitGroup.Use(mwf...)
	unit.Use(mwf...)

	unit.HandleFunc("", hh.CreateUnit).Methods(http.MethodPost)
	unit.HandleFunc("", hh.GetUnitPage).Methods(http.MethodGet)
	unit.HandleFunc("/{unit_id}", hh.GetUnitByID).Methods(http.MethodGet)
	unit.HandleFunc("/{unit_id}", hh.UpdatedUnit).Methods(http.MethodPut)

	unitGroup.HandleFunc("", hh.CreateUnitGroup).Methods(http.MethodPost)
	unitGroup.HandleFunc("", hh.GetUnitGroupPage).Methods(http.MethodGet)
	unitGroup.HandleFunc("/{unit_group_id}", hh.GetUnitGroupByID).Methods(http.MethodGet)
	unitGroup.HandleFunc("/{unit_group_id}/conversions", hh.AddUnitConversions).Methods(http.MethodPost)
}

func (hh *HTTPHandler) CreateUnit(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var unit model.Unit
	if err := json.NewDecoder(r.Body).Decode(&unit); err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	idNewUnit, err := hh.svc.CreateUnit(ctx, &unit)
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	resp.WriteJSONMessage(w, fmt.Sprintf("new unit with id `%d`", idNewUnit), http.StatusCreated)
}

func (hh *HTTPHandler) UpdatedUnit(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := req.GetIntFromVars(r, "unit_id")
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	var unit model.Unit
	if err := json.NewDecoder(r.Body).Decode(&unit); err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}
	unit.ID = id

	err = hh.svc.UpdateUnit(ctx, &unit)
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}
}

func (hh *HTTPHandler) GetUnitPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	page, err := hh.svc.GetUnitPage(ctx, r.URL.Query())
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	resp.WriteJSONData(w, page, http.StatusOK)
}

func (hh *HTTPHandler) GetUnitByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := req.GetIntFromVars(r, "unit_id")
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	unit, err := hh.svc.GetUnitByID(ctx, id)
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	resp.WriteJSONData(w, unit, http.StatusOK)
}
