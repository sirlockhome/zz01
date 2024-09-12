package unit

import (
	"foxomni/internal/unit/datastore"
	"foxomni/internal/unit/port"
	"foxomni/internal/unit/service"
	"foxomni/pkg/database"

	"github.com/gorilla/mux"
)

func InitHTTPRoutes(sql *database.SQL, r *mux.Router, mwf ...mux.MiddlewareFunc) {
	ds := datastore.NewDatastore(sql)
	svc := service.NewService(ds)
	handler := port.NewHTTPHandler(svc)

	handler.Routes(r, mwf...)
}
