package handler

import (
	"github.com/gofiber/fiber/v2"
	"subflow-core-go/internal/api/helper"
)

type GetUserReq struct{}

func (h *Handler) GetUser(c *fiber.Ctx, req GetUserReq) (*helper.UserClaim, error) {
	claim, err := helper.GetClaimFromFiberCtx(c)
	return claim, err
}
