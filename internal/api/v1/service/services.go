package service

import (
	"go.uber.org/fx"
)

type Services struct {
	User UserService
	Team TeamService
}

type Params struct {
	fx.In

	User UserService
	Team TeamService
}

func NewServices(params Params) *Services {
	return &Services{
		params.User,
		params.Team,
	}
}
