package handler

import (
	"errors"
	"reflect"
	"runtime"

	"github.com/casbin/casbin/v2"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"subflow-core-go/internal/api/common"
	"subflow-core-go/internal/api/v1/service"
	"subflow-core-go/internal/config"
)

type Handler struct {
	services *service.Services
	config   *config.Config
	enforcer *casbin.Enforcer
}

func New(services *service.Services, cfg *config.Config, enforcer *casbin.Enforcer) *Handler {
	return &Handler{
		services,
		cfg,
		enforcer,
	}
}

func Wrap[Request any, Response any](handle func(*fiber.Ctx, Request) (Response, error)) fiber.Handler {
	validate := validator.New()

	return func(ctx *fiber.Ctx) error {
		var req Request
		var err error

		err = ctx.ParamsParser(&req)
		if err != nil {
			zap.S().Errorw(
				"Failed to parse params",
				"err", err,
			)
			return err
		}

		err = ctx.BodyParser(&req)
		if err != nil && !errors.Is(err, fiber.ErrUnprocessableEntity) {
			zap.S().Errorw(
				"Failed to parse body",
				"err", err,
			)
			return err
		}

		err = ctx.QueryParser(&req)
		if err != nil {
			zap.S().Errorw(
				"Failed to parse query",
				"err", err,
			)
			return err
		}

		// More parsers here...

		if err := validate.Struct(req); err != nil {
			return err
		}

		zap.S().Debugw(
			"Call handler",
			"handler", runtime.FuncForPC(reflect.ValueOf(handle).Pointer()).Name(),
			"req", req,
		)
		resp, err := handle(ctx, req)
		if err != nil {
			return err
		}

		wrappedResp := common.MakeSuccessAPIResponse(resp)
		return ctx.JSON(wrappedResp)
	}
}
