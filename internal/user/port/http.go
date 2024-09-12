package port

import (
	"encoding/json"
	"foxomni/internal/user/model"
	"foxomni/internal/user/service"
	"foxomni/pkg/errs"
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
	r.HandleFunc("/login", hh.Login).Methods(http.MethodPost)
}

func (hh *HTTPHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var userLogin model.UserLogin
	if err := json.NewDecoder(r.Body).Decode(&userLogin); err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	userAuth, err := hh.svc.Login(ctx, userLogin)
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	resp.WriteJSONData(w, userAuth, http.StatusOK)
}
