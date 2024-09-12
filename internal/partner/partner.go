package partner

import (
	"foxomni/internal/partner/datastore"
	"foxomni/internal/partner/port"
	"foxomni/internal/partner/service"
	"foxomni/pkg/database"

	"github.com/gorilla/mux"
)

func InitHTTPRoutes(sql *database.SQL, r *mux.Router, mwf ...mux.MiddlewareFunc) {
	ds := datastore.NewDatastore(sql)
	svc := service.NewService(ds)
	handler := port.NewHTTPHandler(svc)

	handler.Routes(r, mwf...)
}
