package service

import (
	"context"

	"subflow-core-go/internal/api/common"
	"subflow-core-go/internal/api/constants"
	"subflow-core-go/internal/api/v1/service/dto"
	"subflow-core-go/pkg/ent"
	"subflow-core-go/pkg/ent/team"
)

func (s *Service) FindTeamByID(ctx context.Context, id int) (*ent.Team, error) {
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

func (s *Service) FindTeamByName(ctx context.Context, name string) (*ent.Team, error) {
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

func (s *Service) GetAllTeamsInfo(ctx context.Context) (ent.Teams, error) {
	teams, err := s.db.Team.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return teams, nil
}

func (s *Service) GetTeamInfoByID(ctx context.Context, id int) (*ent.Team, error) {
	t, err := s.db.Team.
		Query().
		Where(team.ID(id)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (s *Service) GetAllUsersOfTeamByID(ctx context.Context, id int) (ent.Users, error) {
	t, err := s.GetTeamInfoByID(ctx, id)
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

func (s *Service) GetAllTasksOfTeamByID(ctx context.Context, id int) (ent.Tasks, error) {
	t, err := s.GetTeamInfoByID(ctx, id)
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

func (s *Service) CreateNewTeam(ctx context.Context, info *dto.TeamInfo) (*ent.Team, error) {
	if t, _ := s.FindTeamByName(ctx, info.Name); t != nil {
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

func (s *Service) UpdateTeamInfoByID(ctx context.Context, id int, info *dto.TeamInfo) (*ent.Team, error) {
	t, err := s.FindTeamByID(ctx, id)
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

func (s *Service) AddUsersForTeam(ctx context.Context, teamID int, u ...*ent.User) error {
	t, err := s.FindTeamByID(ctx, teamID)
	if err != nil {
		return err
	}
	return t.
		Update().
		AddUsers(u...).
		Exec(ctx)
}
