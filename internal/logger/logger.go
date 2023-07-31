package logger

import (
	"go.uber.org/zap"
	"subflow-core-go/internal/config"
)

func New(cfg *config.Config) (*zap.Logger, error) {
	var logger *zap.Logger
	var err error
	switch cfg.Server.Type {
	case config.ServerTypeProd:
		logger, err = zap.NewProduction()
	case config.ServerTypeDev:
		logger, err = zap.NewDevelopment()
	default:
		zap.S().Fatalf("Unsupported server type: %v", cfg.Server.Type)
	}
	if err != nil {
		return nil, err
	}
	zap.ReplaceGlobals(logger)
	zap.S().Infow(
		"Logger initialized",
		"type", cfg.Server.Type,
	)
	return logger, nil
}
