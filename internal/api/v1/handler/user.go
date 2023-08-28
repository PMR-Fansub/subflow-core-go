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

// GetCurrentUser godoc
//
//	@Summary	Get current logged user info
//	@Tags		user
//	@Accept		json
//	@Produce	json
//	@Security	ApiKeyAuth
//	@Success	200	{object}	common.APIResponse{data=dto.UserInfo}
//	@Router		/user [get]
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

// GetUserByID godoc
//
//	@Summary	Get user basic info by UID
//	@Tags		user
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"user id"
//	@Success	200	{object}	common.APIResponse{data=dto.UserBasicInfo}
//	@Router		/user/{id} [get]
func (h *Handler) GetUserByID(c *fiber.Ctx, req GetUserByIDReq) (*dto.UserBasicInfo, error) {
	user, err := h.service.FindUserByID(context.Background(), req.ID)
	if err != nil {
		return nil, err
	}
	info := service.GetBasicInfoFromUser(user)
	return info, err
}

// UpdateCurrentUser godoc
//
//	@Summary	Update current logged user info
//	@Tags		user
//	@Accept		json
//	@Produce	json
//	@Security	ApiKeyAuth
//	@Param		message	body		UpdateCurUserReq	true	"user info to update"
//	@Success	200		{object}	common.APIResponse{data=UpdateUserInfoResp}
//	@Router		/user [patch]
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

// UpdateUser godoc
//
//	@Summary	Update user info by UID (admin)
//	@Tags		user
//	@Accept		json
//	@Produce	json
//	@Security	ApiKeyAuth
//	@Param		id		path		int					true	"user id"
//	@Param		message	body		UpdateCurUserReq	true	"user info to update"
//	@Success	200		{object}	common.APIResponse{data=UpdateUserInfoResp}
//	@Router		/user/{id} [patch]
func (h *Handler) UpdateUser(c *fiber.Ctx, req UpdateUserReq) (*UpdateUserInfoResp, error) {
	err := h.service.UpdateUser(
		context.Background(), &dto.UpdateUserReq{
			ID:       req.ID,
			Nickname: req.Nickname,
		},
	)
	return nil, err
}

// GetUserTeamsByID godoc
//
//	@Summary	Get all teams that the specified user belongs to
//	@Tags		user
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"user id"
//	@Success	200	{object}	common.APIResponse{data=[]dto.TeamInfo}
//	@Router		/user/{id}/teams [get]
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
