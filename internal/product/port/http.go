package port

import (
	"encoding/json"
	"fmt"
	"foxomni/internal/product/model"
	"foxomni/internal/product/service"
	"foxomni/pkg/errs"
	"foxomni/pkg/resp"
	"net/http"
	"strconv"

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
	prdGroup := r.PathPrefix("/products").Subrouter()

	prdGroup.HandleFunc("", hh.CreateProduct).Methods(http.MethodPost)
	prdGroup.HandleFunc("", hh.GetProductPage).Methods(http.MethodGet)
	prdGroup.HandleFunc("/{product_id}", hh.GetProductByID).Methods(http.MethodGet)
	prdGroup.HandleFunc("/{product_id}", hh.UpdateProduct).Methods(http.MethodPut)
	prdGroup.HandleFunc("/{product_id}", hh.DeleteProduct).Methods(http.MethodDelete)

	prdGroup.HandleFunc("/{product_id}/thumbnail", hh.ChangeThumbnail).Methods(http.MethodPut)

	prdGroup.HandleFunc("/{product_id}/images", hh.AddImages).Methods(http.MethodPost)
	prdGroup.HandleFunc("/{product_id}/images/{image_id}", hh.RemoveImage).Methods(http.MethodDelete)

	prdGroup.HandleFunc("/{product_id}/attachments", hh.AddAttachments).Methods(http.MethodPost)
	prdGroup.HandleFunc("/{product_id}/attachments/{attachment_id}", hh.RemoveAttachment).Methods(http.MethodDelete)

	prdGroup.HandleFunc("/{product_id}/categories", hh.AddCategories).Methods(http.MethodPost)
	prdGroup.HandleFunc("/{product_id}/categories/{category_id}", hh.RemoveCategory).Methods(http.MethodDelete)
}

func (hh *HTTPHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := strconv.Atoi(mux.Vars(r)["product_id"])
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	prd, err := hh.svc.GetProductByID(ctx, id)
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	resp.WriteJSONData(w, prd, http.StatusOK)
}

func (hh *HTTPHandler) GetProductPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	page, err := hh.svc.GetProductPage(ctx, r.URL.Query())
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	resp.WriteJSONData(w, page, http.StatusOK)
}

func (hh *HTTPHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var prd model.ProductRequest

	err := json.NewDecoder(r.Body).Decode(&prd)
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	idNewPd, err := hh.svc.CreateProduct(ctx, &prd)
	if err != nil {
		errs.HTTPErrorResponse(ctx, w, err)
		return
	}

	resp.WriteJSONMessage(w, fmt.Sprintf("new product with id `%d`", idNewPd), http.StatusCreated)
}
