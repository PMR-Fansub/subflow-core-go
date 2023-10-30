package handler

import (
	"github.com/gofiber/fiber/v2"
	"subflow-core-go/internal/api/v1/service/dto"
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
func (h *Handler) GetAllTeams(ctx *fiber.Ctx, _ GetAllTeamsReq) ([]*dto.TeamInfo, error) {
	teams, err := h.services.Team.GetAllTeamsInfo(ctx.Context())
	if err != nil {
		return nil, err
	}
	return dto.GetTeamInfoFromEntities(teams), err
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
func (h *Handler) GetTeamByID(ctx *fiber.Ctx, req GetTeamByIDReq) (*dto.TeamInfo, error) {
	t, err := h.services.Team.GetTeamByID(ctx.Context(), req.ID)
	if err != nil {
		return nil, err
	}
	return dto.GetTeamInfoFromEntity(t), nil
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
func (h *Handler) GetTeamUsersByID(ctx *fiber.Ctx, req GetTeamUsersByIDReq) ([]*dto.UserBasicInfo, error) {
	us, err := h.services.Team.GetAllUsersOfTeamByID(ctx.Context(), req.ID)
	if err != nil {
		return nil, err
	}
	return dto.GetBasicInfoFromUsers(us), nil
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
func (h *Handler) GetTeamTasksByID(ctx *fiber.Ctx, req GetTeamTasksByIDReq) ([]*dto.TaskInfo, error) {
	tasks, err := h.services.Team.GetAllTasksOfTeamByID(ctx.Context(), req.ID)
	if err != nil {
		return nil, err
	}
	return dto.GetTaskInfoFromEntities(tasks), nil
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
func (h *Handler) CreateNewTeam(ctx *fiber.Ctx, req CreateNewTeamReq) (*dto.TeamInfo, error) {
	t, err := h.services.Team.CreateNewTeam(
		ctx.Context(), &dto.TeamInfo{
			Name:    req.Name,
			QQGroup: req.QQGroup,
			Logo:    req.Logo,
			Desc:    req.Desc,
		},
	)
	if err != nil {
		return nil, err
	}
	return dto.GetTeamInfoFromEntity(t), nil
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
func (h *Handler) UpdateTeamInfoByID(ctx *fiber.Ctx, req UpdateTeamReq) (*dto.TeamInfo, error) {
	t, err := h.services.Team.UpdateTeamInfoByID(
		ctx.Context(), req.ID, &dto.TeamInfo{
			Name:    req.Name,
			Status:  req.Status,
			QQGroup: req.QQGroup,
			Logo:    req.Logo,
			Desc:    req.Desc,
		},
	)
	if err != nil {
		return nil, err
	}
	return dto.GetTeamInfoFromEntity(t), nil
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
func (h *Handler) AddUserToTeam(ctx *fiber.Ctx, req AddUserToTeamReq) (*AddUserToTeamResp, error) {
	u, err := h.services.User.GetUserByID(ctx.Context(), req.UID)
	if err != nil {
		return nil, err
	}
	err = h.services.Team.AddUsersForTeam(ctx.Context(), req.ID, u)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
