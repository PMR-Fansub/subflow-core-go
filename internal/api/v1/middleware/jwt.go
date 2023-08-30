package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"subflow-core-go/internal/api/common"
	"subflow-core-go/internal/api/helper"
	"subflow-core-go/internal/config"
)

func NewJwtMiddleware(cfg *config.Config) (fiber.Handler, error) {
	return jwtware.New(
		jwtware.Config{
			SigningMethod: "HS256",
			SigningKey:    []byte(cfg.Server.SigningKey),
			ErrorHandler:  jwtErrorHandler,
			TokenLookup:   "cookie:auth_token,header:Authorization",
			Claims:        &helper.UserClaim{},
		},
	), nil
}

func jwtErrorHandler(ctx *fiber.Ctx, err error) error {
	result := common.ResultUnauthorized
	resp := common.MakeAPIResponse(result, err.Error())
	return ctx.Status(result.HttpCode).JSON(resp)
}
