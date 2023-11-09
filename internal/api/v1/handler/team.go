package handler

import (
	"github.com/gofiber/fiber/v2"
	"subflow-core-go/internal/api/v1/dto"
)

type GetAllTeamsReq struct{}

type GetTeamUsersByIDReq struct {
	ID int `params:"id"`
}

type GetTeamByIDReq struct {
	ID int `params:"id"`
}

type GetTeamTasksByIDReq struct {
	ID int `params:"id"`
}

type CreateNewTeamReq struct {
	Name    string `json:"name" validate:"required"`
	QQGroup string `json:"QQGroup"`
	Logo    string `json:"logo"`
	Desc    string `json:"desc"`
}

type UpdateTeamReq struct {
	ID      int    `params:"id"`
	Status  int    `json:"status" validate:"oneof=1 2"`
	Name    string `json:"name"`
	QQGroup string `json:"QQGroup"`
	Logo    string `json:"logo"`
	Desc    string `json:"desc"`
}

type AddUserToTeamReq struct {
	ID  int `params:"id"`
	UID int `json:"uid"`
}

type AddUserToTeamResp struct {
}

// GetAllTeams godoc
//
//	@Summary	Get all teams
//	@Tags		teams
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	common.APIResponse{data=[]dto.TeamInfo}
//	@Router		/teams [get]
func (h *Handler) GetAllTeams(ctx *fiber.Ctx, _ GetAllTeamsReq) (res []dto.TeamInfo, err error) {
	teams, err := h.services.Team.GetAllTeams(ctx.Context())
	if err != nil {
		return
	}
	res = dto.GetTeamInfoFromEntities(teams)
	return
}

// GetTeamByID godoc
//
//	@Summary	Get team info by ID
//	@Tags		teams
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"team id"
//	@Success	200	{object}	common.APIResponse{data=dto.TeamInfo}
//	@Router		/teams/{id} [get]
func (h *Handler) GetTeamByID(ctx *fiber.Ctx, req GetTeamByIDReq) (res dto.TeamInfo, err error) {
	t, err := h.services.Team.GetTeamByID(ctx.Context(), req.ID)
	if err != nil {
		return
	}
	res = dto.GetTeamInfoFromEntity(t)
	return
}

// GetTeamUsersByID godoc
//
//	@Summary	Get all users for the specified team
//	@Tags		teams
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"team id"
//	@Success	200	{object}	common.APIResponse{data=[]dto.UserBasicInfo}
//	@Router		/teams/{id}/users [get]
func (h *Handler) GetTeamUsersByID(ctx *fiber.Ctx, req GetTeamUsersByIDReq) (res []dto.UserBasicInfo, err error) {
	us, err := h.services.Team.GetAllUsersOfTeamByID(ctx.Context(), req.ID)
	if err != nil {
		return
	}
	res = dto.GetBasicInfoFromUsers(us)
	return
}

// GetTeamTasksByID godoc
//
//	@Summary	Get all tasks for the specified team
//	@Tags		teams
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"team id"
//	@Success	200	{object}	common.APIResponse{data=[]dto.TaskInfo}
//	@Router		/teams/{id}/tasks [get]
func (h *Handler) GetTeamTasksByID(ctx *fiber.Ctx, req GetTeamTasksByIDReq) (res []dto.TaskInfo, err error) {
	tasks, err := h.services.Team.GetAllTasksOfTeamByID(ctx.Context(), req.ID)
	if err != nil {
		return
	}
	res = dto.GetTaskInfoFromEntities(tasks)
	return
}

// CreateNewTeam godoc
//
//	@Summary	Create new team
//	@Tags		teams
//	@Accept		json
//	@Produce	json
//	@Param		message	body		CreateNewTeamReq	true	"new team info"
//	@Success	200		{object}	common.APIResponse{data=dto.TeamInfo}
//	@Router		/teams [post]
func (h *Handler) CreateNewTeam(ctx *fiber.Ctx, req CreateNewTeamReq) (res dto.TeamInfo, err error) {
	t, err := h.services.Team.CreateNewTeam(
		ctx.Context(), dto.TeamInfo{
			Name:    req.Name,
			QQGroup: req.QQGroup,
			Logo:    req.Logo,
			Desc:    req.Desc,
		},
	)
	if err != nil {
		return
	}
	res = dto.GetTeamInfoFromEntity(t)
	return
}

// UpdateTeamInfoByID godoc
//
//	@Summary	Update the team for the specified team
//	@Tags		teams
//	@Accept		json
//	@Produce	json
//	@Param		id		path		int				true	"team id"
//	@Param		message	body		UpdateTeamReq	true	"team info"
//	@Success	200		{object}	common.APIResponse{data=dto.TeamInfo}
//	@Router		/teams/{id} [patch]
func (h *Handler) UpdateTeamInfoByID(ctx *fiber.Ctx, req UpdateTeamReq) (res dto.TeamInfo, err error) {
	t, err := h.services.Team.UpdateTeamByID(
		ctx.Context(), req.ID, dto.TeamInfo{
			Name:    req.Name,
			Status:  req.Status,
			QQGroup: req.QQGroup,
			Logo:    req.Logo,
			Desc:    req.Desc,
		},
	)
	if err != nil {
		return
	}
	res = dto.GetTeamInfoFromEntity(t)
	return
}

// AddUserToTeam godoc
//
//	@Summary	Add specified user to team
//	@Tags		teams
//	@Accept		json
//	@Produce	json
//	@Param		id		path		int					true	"team id"
//	@Param		message	body		AddUserToTeamReq	true	"user info"
//	@Success	200		{object}	common.APIResponse{data=AddUserToTeamResp}
//	@Router		/teams/{id}/users [post]
func (h *Handler) AddUserToTeam(ctx *fiber.Ctx, req AddUserToTeamReq) (res AddUserToTeamResp, err error) {
	u, err := h.services.User.GetUserByID(ctx.Context(), req.UID)
	if err != nil {
		return
	}
	err = h.services.Team.AddUsersForTeam(ctx.Context(), req.ID, u)
	if err != nil {
		return
	}
	return
}
