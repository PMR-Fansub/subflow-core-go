package handler

import (
	"errors"
	"reflect"
	"runtime"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"subflow-core-go/internal/api/common"
	"subflow-core-go/internal/api/v1/service"
	"subflow-core-go/internal/config"
)

type Handler struct {
	service *service.Service
	config  *config.Config
}

func New(service *service.Service, cfg *config.Config) *Handler {
	return &Handler{
		service,
		cfg,
	}
}

func WrapHandlerWithAutoParse[Request any, Response any](handle func(*fiber.Ctx, Request) (Response, error)) fiber.Handler {
	validate := validator.New()

	return func(c *fiber.Ctx) error {
		var req Request
		var err error

		err = c.ParamsParser(&req)
		if err != nil {
			zap.S().Errorw(
				"Failed to parse params",
				"err", err,
			)
			return err
		}

		err = c.BodyParser(&req)
		if err != nil && !errors.Is(err, fiber.ErrUnprocessableEntity) {
			zap.S().Errorw(
				"Failed to parse body",
				"err", err,
			)
			return err
		}

		err = c.QueryParser(&req)
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
		resp, err := handle(c, req)
		if err != nil {
			return err
		}

		wrappedResp := common.MakeSuccessAPIResponse(resp)
		return c.JSON(wrappedResp)
	}
}
