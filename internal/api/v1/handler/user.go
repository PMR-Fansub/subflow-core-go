package handler

import (
	"github.com/gofiber/fiber/v2"
	"subflow-core-go/internal/api/helper"
	"subflow-core-go/internal/api/v1/dto"
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
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Security	ApiKeyAuth
//	@Success	200	{object}	common.APIResponse{data=dto.UserInfo}
//	@Router		/users [get]
func (h *Handler) GetCurrentUser(ctx *fiber.Ctx, _ GetUserReq) (res dto.UserInfo, err error) {
	claim, err := helper.GetClaimFromFiberCtx(ctx)
	if err != nil {
		return
	}
	user, err := h.services.User.GetUserByID(ctx.Context(), claim.UID)
	if err != nil {
		return
	}

	res = dto.GetInfoFromUser(user)
	return
}

// GetUserByID godoc
//
//	@Summary	Get user basic info by UID
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"user id"
//	@Success	200	{object}	common.APIResponse{data=dto.UserBasicInfo}
//	@Router		/users/{id} [get]
func (h *Handler) GetUserByID(ctx *fiber.Ctx, req GetUserByIDReq) (res dto.UserBasicInfo, err error) {
	user, err := h.services.User.GetUserByID(ctx.Context(), req.ID)
	if err != nil {
		return
	}
	res = dto.GetBasicInfoFromUser(user)
	return
}

// UpdateCurrentUser godoc
//
//	@Summary	Update current logged user info
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Security	ApiKeyAuth
//	@Param		message	body		UpdateCurUserReq	true	"user info to update"
//	@Success	200		{object}	common.APIResponse{data=UpdateUserInfoResp}
//	@Router		/users [patch]
func (h *Handler) UpdateCurrentUser(ctx *fiber.Ctx, req UpdateCurUserReq) (res UpdateUserInfoResp, err error) {
	claim, err := helper.GetClaimFromFiberCtx(ctx)
	if err != nil {
		return
	}
	err = h.services.User.UpdateUser(
		ctx.Context(), dto.UpdateUserReq{
			ID:       claim.UID,
			Nickname: req.Nickname,
		},
	)
	return
}

// UpdateUser godoc
//
//	@Summary	Update user info by UID (admin)
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Security	ApiKeyAuth
//	@Param		id		path		int					true	"user id"
//	@Param		message	body		UpdateCurUserReq	true	"user info to update"
//	@Success	200		{object}	common.APIResponse{data=UpdateUserInfoResp}
//	@Router		/users/{id} [patch]
func (h *Handler) UpdateUser(ctx *fiber.Ctx, req UpdateUserReq) (res UpdateUserInfoResp, err error) {
	err = h.services.User.UpdateUser(
		ctx.Context(), dto.UpdateUserReq{
			ID:       req.ID,
			Nickname: req.Nickname,
		},
	)
	return
}

// GetUserTeamsByID godoc
//
//	@Summary	Get all teams that the specified user belongs to
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"user id"
//	@Success	200	{object}	common.APIResponse{data=[]dto.TeamInfo}
//	@Router		/users/{id}/teams [get]
func (h *Handler) GetUserTeamsByID(ctx *fiber.Ctx, req GetUserTeamsByIDReq) (res []dto.TeamInfo, err error) {
	u, err := h.services.User.GetUserByID(ctx.Context(), req.UID)
	if err != nil {
		return
	}
	ts, err := h.services.User.GetTeamsOfUser(ctx.Context(), u)
	if err != nil {
		return
	}
	res = dto.GetTeamInfoFromEntities(ts)
	return
}
