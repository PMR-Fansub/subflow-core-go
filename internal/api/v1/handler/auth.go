package handler

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"subflow-core-go/internal/api/helper"
	"subflow-core-go/internal/api/v1/service"
)

type RegisterRequest struct {
	*service.CreateUserRequest
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterResponse struct{}

func (h *Handler) Register(c *fiber.Ctx, req RegisterRequest) (*RegisterResponse, error) {
	req.RemoteAddr = c.IP()
	_, err := h.service.CreateUser(context.Background(), req.CreateUserRequest)
	return nil, err
}

func (h *Handler) Login(c *fiber.Ctx, req LoginRequest) (*LoginResponse, error) {
	var resp LoginResponse

	user, err := h.service.VerifyPwdByUsername(context.Background(), req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	token, err := helper.SignJWT(
		h.config.Server.SigningKey, &helper.UserClaim{
			UID:      user.ID,
			Username: user.Username,
			Exp:      time.Now().Add(time.Hour * 72).Unix(),
		},
	)
	if err != nil {
		return nil, err
	}

	resp.Token = token
	return &resp, nil
}
