package service

import (
	"context"

	"subflow-core-go/internal/api/common"
	"subflow-core-go/internal/api/constants"
	"subflow-core-go/internal/api/v1/dto"
	"subflow-core-go/internal/config"
	"subflow-core-go/pkg/ent"
	"subflow-core-go/pkg/ent/team"
)

type TeamService interface {
	GetTeamByID(ctx context.Context, id int) (*ent.Team, error)
	GetTeamByName(ctx context.Context, name string) (*ent.Team, error)
	GetAllTeams(ctx context.Context) (ent.Teams, error)
	GetAllUsersOfTeamByID(ctx context.Context, id int) (ent.Users, error)
	GetAllTasksOfTeamByID(ctx context.Context, id int) (ent.Tasks, error)
	CreateNewTeam(ctx context.Context, info dto.TeamInfo) (*ent.Team, error)
	UpdateTeamByID(ctx context.Context, id int, info dto.TeamInfo) (*ent.Team, error)
	AddUsersForTeam(ctx context.Context, teamID int, u ...*ent.User) error
}

type TeamServiceImpl struct {
	db     *ent.Client
	config *config.Config
}

func NewTeamService(db *ent.Client, config *config.Config) TeamService {
	return &TeamServiceImpl{
		db,
		config,
	}
}

func (s *TeamServiceImpl) GetTeamByID(ctx context.Context, id int) (*ent.Team, error) {
	t, err := s.db.Team.
		Query().
		Where(team.ID(id)).
		Only(ctx)
	if err != nil && ent.IsNotFound(err) {
		return nil, &common.BusinessError{
			Code:    common.ResultNotFound,
			Message: "团队不存在",
		}
	}
	return t, err
}

func (s *TeamServiceImpl) GetTeamByName(ctx context.Context, name string) (*ent.Team, error) {
	t, err := s.db.Team.
		Query().
		Where(team.Name(name)).
		Only(ctx)
	if err != nil && ent.IsNotFound(err) {
		return nil, &common.BusinessError{
			Code:    common.ResultNotFound,
			Message: "团队不存在",
		}
	}
	return t, err
}

func (s *TeamServiceImpl) GetAllTeams(ctx context.Context) (ent.Teams, error) {
	teams, err := s.db.Team.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return teams, nil
}

func (s *TeamServiceImpl) GetAllUsersOfTeamByID(ctx context.Context, id int) (ent.Users, error) {
	t, err := s.GetTeamByID(ctx, id)
	if err != nil {
		return nil, err
	}
	u, err := t.
		QueryUsers().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *TeamServiceImpl) GetAllTasksOfTeamByID(ctx context.Context, id int) (ent.Tasks, error) {
	t, err := s.GetTeamByID(ctx, id)
	if err != nil {
		return nil, err
	}
	tasks, err := t.
		QueryTasks().
		WithTaskTags().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *TeamServiceImpl) CreateNewTeam(ctx context.Context, info dto.TeamInfo) (*ent.Team, error) {
	if t, _ := s.GetTeamByName(ctx, info.Name); t != nil {
		return nil, &common.BusinessError{
			Code:    common.ResultCreationFailed,
			Message: "团队名称已被使用",
		}
	}
	return s.db.Team.
		Create().
		SetName(info.Name).
		SetStatus(int(constants.TeamStatusOpen)).
		SetQqGroup(info.QQGroup).
		SetLogo(info.Logo).
		SetDesc(info.Desc).
		Save(ctx)
}

func (s *TeamServiceImpl) UpdateTeamByID(ctx context.Context, id int, info dto.TeamInfo) (*ent.Team, error) {
	t, err := s.GetTeamByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return t.
		Update().
		SetName(info.Name).
		SetQqGroup(info.QQGroup).
		SetLogo(info.Logo).
		SetDesc(info.Desc).
		Save(ctx)
}

func (s *TeamServiceImpl) AddUsersForTeam(ctx context.Context, teamID int, u ...*ent.User) error {
	t, err := s.GetTeamByID(ctx, teamID)
	if err != nil {
		return err
	}
	return t.
		Update().
		AddUsers(u...).
		Exec(ctx)
}
