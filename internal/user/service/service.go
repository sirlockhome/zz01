package service

import (
	"foxomni/internal/user/datastore"
	"foxomni/pkg/config"
	"foxomni/pkg/jwt"
)

type Service struct {
	ds   *datastore.Datastore
	jwt  *jwt.Service
	conf *config.Config
}

func NewService(ds *datastore.Datastore, jwt *jwt.Service, conf *config.Config) *Service {
	return &Service{
		ds:   ds,
		jwt:  jwt,
		conf: conf,
	}
}
