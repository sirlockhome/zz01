package service

import "foxomni/internal/category/datastore"

type Service struct {
	ds *datastore.Datastore
}

func NewService(ds *datastore.Datastore) *Service {
	return &Service{
		ds: ds,
	}
}
