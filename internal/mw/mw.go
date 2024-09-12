package mw

import "foxomni/pkg/config"

type Middleware struct {
	conf config.Config
}

func NewMiddleware(conf config.Config) *Middleware {
	return &Middleware{
		conf: conf,
	}
}
