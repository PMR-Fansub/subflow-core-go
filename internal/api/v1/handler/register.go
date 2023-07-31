package handler

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"subflow-core-go/internal/api/v1/service"
)

func (h *Handler) Register(c *fiber.Ctx, req service.CreateUserRequest) (string, error) {
	req.RemoteAddr = c.IP()
	_, err := h.service.CreateUser(context.Background(), req)
	return "", err
}
