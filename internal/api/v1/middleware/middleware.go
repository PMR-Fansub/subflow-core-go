package middleware

import (
	"subflow-core-go/internal/api/v1/service"
	"subflow-core-go/internal/config"
)

type Middleware struct {
	service *service.Service
	config  *config.Config
}

func New(service *service.Service, cfg *config.Config) *Middleware {
	return &Middleware{
		service,
		cfg,
	}
}
