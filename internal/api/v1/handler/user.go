package handler

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"subflow-core-go/internal/api/helper"
	"subflow-core-go/internal/api/v1/service"
	"subflow-core-go/internal/api/v1/service/dto"
)

type GetUserReq struct{}

type GetUserByIDReq struct {
	ID int `params:"id"`
}

type UpdateCurUserReq struct {
	Nickname string `json:"nickname"`
}

type UpdateUserReq struct {
	ID int `params:"id"`
	UpdateCurUserReq
}

type GetUserTeamsByIDReq struct {
	UID int `params:"id"`
}

type UpdateUserInfoResp struct{}

func (h *Handler) GetCurrentUser(c *fiber.Ctx, req GetUserReq) (*dto.UserInfo, error) {
	claim, err := helper.GetClaimFromFiberCtx(c)
	if err != nil {
		return nil, err
	}
	user, err := h.service.FindUserByID(context.Background(), claim.UID)
	if err != nil {
		return nil, err
	}

	userInfo := service.GetInfoFromUser(user)
	return userInfo, err
}

func (h *Handler) GetUserByID(c *fiber.Ctx, req GetUserByIDReq) (*dto.UserBasicInfo, error) {
	user, err := h.service.FindUserByID(context.Background(), req.ID)
	if err != nil {
		return nil, err
	}
	info := service.GetBasicInfoFromUser(user)
	return info, err
}

func (h *Handler) UpdateCurrentUser(c *fiber.Ctx, req UpdateCurUserReq) (*UpdateUserInfoResp, error) {
	claim, err := helper.GetClaimFromFiberCtx(c)
	if err != nil {
		return nil, err
	}
	err = h.service.UpdateUser(
		context.Background(), &dto.UpdateUserReq{
			ID:       claim.UID,
			Nickname: req.Nickname,
		},
	)
	return nil, err
}

func (h *Handler) UpdateUser(c *fiber.Ctx, req UpdateUserReq) (*UpdateUserInfoResp, error) {
	err := h.service.UpdateUser(
		context.Background(), &dto.UpdateUserReq{
			ID:       req.ID,
			Nickname: req.Nickname,
		},
	)
	return nil, err
}

func (h *Handler) GetUserTeamsByID(ctx *fiber.Ctx, req GetUserTeamsByIDReq) ([]*dto.TeamInfo, error) {
	u, err := h.service.FindUserByID(context.Background(), req.UID)
	if err != nil {
		return nil, err
	}
	tInfo, err := h.service.GetTeamsOfUser(context.Background(), u)
	if err != nil {
		return nil, err
	}
	return tInfo, nil
}
