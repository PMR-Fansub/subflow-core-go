package main

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"subflow-core-go/internal/api"
	v1 "subflow-core-go/internal/api/v1"
	"subflow-core-go/internal/config"
	"subflow-core-go/internal/datasource"
	"subflow-core-go/internal/logger"
)

func main() {
	fx.New(
		fx.Provide(config.New),
		fx.Provide(datasource.NewEntClient),
		fx.Provide(logger.New),
		fx.Provide(api.New),
		v1.Module,
		fx.WithLogger(
			func(log *zap.Logger) fxevent.Logger {
				return &fxevent.ZapLogger{Logger: log}
			},
		),
		fx.Invoke(api.Start),
	).Run()
}
