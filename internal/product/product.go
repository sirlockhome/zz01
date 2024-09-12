package product

import (
	"foxomni/internal/product/datastore"
	"foxomni/internal/product/port"
	"foxomni/internal/product/service"
	"foxomni/pkg/database"

	"github.com/gorilla/mux"
)

func InitHTTPRoutes(sql *database.SQL, r *mux.Router, mwf ...mux.MiddlewareFunc) {
	ds := datastore.NewDatastore(sql)
	svc := service.NewService(ds)
	handler := port.NewHTTPHandler(svc)

	handler.Routes(r, mwf...)
}
