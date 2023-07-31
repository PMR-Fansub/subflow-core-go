package handler

import (
	"errors"
	"time"

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

func AutoParsedHandler[Request any, Response any](handle func(*fiber.Ctx, Request) (Response, error)) fiber.Handler {
	validate := validator.New()

	return func(c *fiber.Ctx) error {
		var req Request

		if err := c.BodyParser(&req); err != nil && !errors.Is(err, fiber.ErrUnprocessableEntity) {
			zap.S().Warnw(
				"Failed to parse body",
				"err", err,
			)
			return err
		}

		if err := c.QueryParser(&req); err != nil {
			zap.S().Warnw(
				"Failed to parse query",
				"err", err,
			)
			return err
		}

		// More parsers here...

		if err := validate.Struct(req); err != nil {
			return err
		}

		resp, err := handle(c, req)
		if err != nil {
			return err
		}

		wrappedResp := common.APIResponse{
			Code:      common.ResultCodeSuccess,
			Success:   true,
			Message:   "success",
			Timestamp: time.Now().Unix(),
			Data:      resp,
		}
		return c.JSON(wrappedResp)
	}
}
