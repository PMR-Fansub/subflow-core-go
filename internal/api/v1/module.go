package v1

import (
	"go.uber.org/fx"
	"subflow-core-go/internal/api/helper"
	"subflow-core-go/internal/api/v1/handler"
	"subflow-core-go/internal/api/v1/middleware"
	"subflow-core-go/internal/api/v1/service"
)

var Module = fx.Module(
	"v1",
	service.Module,
	fx.Provide(
		handler.New,
		helper.NewAdapter,
		middleware.NewCasbinEnforcer,
		middleware.NewCasbinMiddleware,
		middleware.NewJwtMiddleware,
		NewRouter,
	),
)
