package handler

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"subflow-core-go/internal/api/v1/service"
)

func (h *Handler) Register(c *fiber.Ctx, req service.CreateUserRequest) (string, error) {
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		zap.S().Errorw(
			"Failed to crypt password",
			"req", req,
			"err", err,
		)
		return "", err
	}
	req.RemoteAddr = c.IP()
	req.Password = string(pwdHash)
	_, err = h.service.CreateUser(context.Background(), req)
	return "", err
}
