package server

import (
	"fmt"
	"foxomni/internal/category"
	"foxomni/internal/mw"
	"foxomni/internal/order"
	"foxomni/internal/partner"
	"foxomni/internal/product"
	"foxomni/internal/unit"
	"foxomni/internal/user"
	"foxomni/pkg/config"
	"foxomni/pkg/database"
	"foxomni/pkg/jwt"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	port int
	r    *mux.Router
}

func initMw(conf config.Config) []mux.MiddlewareFunc {
	mw := mw.NewMiddleware(conf)
	var mwf []mux.MiddlewareFunc

	mwf = append(mwf, mw.ValidateJWT)
	return mwf
}

func NewHTTPServer(conf config.Config, sql *database.SQL) *HTTPServer {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	api.Use(mw.RequestLogging)

	jwt := jwt.NewService(conf.Auth)

	mwf := initMw(conf)

	product.InitHTTPRoutes(sql, api, mwf...)
	partner.InitHTTPRoutes(sql, api, mwf...)
	unit.InitHTTPRoutes(sql, api, mwf...)
	order.InitHTTPRoutes(sql, api, mwf...)
	user.InitHTTPRoutes(sql, jwt, &conf, api)
	category.InitHTTPRoutes(sql, api, mwf...)

	return &HTTPServer{
		port: conf.Server.Port,
		r:    r,
	}
}

func (s *HTTPServer) RunHTTPServer() error {
	fmt.Println("starting server...")

	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.r)
}
