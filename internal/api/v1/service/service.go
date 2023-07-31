package service

import (
	"subflow-core-go/internal/config"
	"subflow-core-go/pkg/ent"
)

type Service struct {
	db     *ent.Client
	config *config.Config
}

func New(db *ent.Client, cfg *config.Config) *Service {
	return &Service{
		db:     db,
		config: cfg,
	}
}
