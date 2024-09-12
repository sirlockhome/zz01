package user

import (
	"foxomni/internal/user/datastore"
	"foxomni/internal/user/port"
	"foxomni/internal/user/service"
	"foxomni/pkg/config"
	"foxomni/pkg/database"
	"foxomni/pkg/jwt"

	"github.com/gorilla/mux"
)

func InitHTTPRoutes(sql *database.SQL, jwt *jwt.Service, conf *config.Config, r *mux.Router, mwf ...mux.MiddlewareFunc) {
	ds := datastore.NewDatastore(sql)
	svc := service.NewService(ds, jwt, conf)
	handler := port.NewHTTPHandler(svc)

	handler.Routes(r, mwf...)
}
