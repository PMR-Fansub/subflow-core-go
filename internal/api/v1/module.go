package v1

import (
	"go.uber.org/fx"
	"subflow-core-go/internal/api/v1/handler"
	"subflow-core-go/internal/api/v1/service"
)

var Module = fx.Provide(
	service.New,
	handler.New,
	// middleware.New,
	NewRouter,
)
