package service

import (
	"context"

	"subflow-core-go/internal/api/v1/service/dto"
	"subflow-core-go/pkg/ent"
	"subflow-core-go/pkg/ent/team"
)

func QueryTeamWithEdge(q *ent.TeamQuery) {
	q.Select(
		team.FieldName,
		team.FieldStatus,
		team.FieldQqGroup,
		team.FieldLogo,
	)
}
func GetTeamInfoFromEdge(teams []*ent.Team) []*dto.TeamInfo {
	var teamInfo []*dto.TeamInfo
	for _, t := range teams {
		teamInfo = append(
			teamInfo, &dto.TeamInfo{
				ID:      t.ID,
				Name:    t.Name,
				Status:  t.Status,
				QQGroup: t.QqGroup,
				Logo:    t.Logo,
				Desc:    t.Desc,
			},
		)
	}
	return teamInfo
}

func (s *Service) GetTeamsOfUser(ctx context.Context, u *ent.User) ([]*dto.TeamInfo, error) {
	t, err := u.
		QueryTeams().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return GetTeamInfoFromEdge(t), nil
}
