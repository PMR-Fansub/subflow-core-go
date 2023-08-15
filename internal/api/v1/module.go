package v1

import (
	"go.uber.org/fx"
	"subflow-core-go/internal/api/helper"
	"subflow-core-go/internal/api/v1/handler"
	"subflow-core-go/internal/api/v1/middleware"
	"subflow-core-go/internal/api/v1/service"
)

var Module = fx.Provide(
	service.New,
	handler.New,
	helper.NewAdapter,
	middleware.NewCasbinEnforcer,
	middleware.NewCasbinMiddleware,
	NewRouter,
)
