package port

import (
	"foxomni/internal/category/service"
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
	cat := r.PathPrefix("/categories").Subrouter()
	cat.Use(mwf...)

	cat.HandleFunc("", hh.CreateCategory).Methods(http.MethodPost)
	cat.HandleFunc("", hh.GetCategoryPage).Methods(http.MethodGet)
	cat.HandleFunc("/{category_id}", hh.GetCategoryByID).Methods(http.MethodGet)
	cat.HandleFunc("/{category_id}", hh.UpdateCategory).Methods(http.MethodPut)
}
