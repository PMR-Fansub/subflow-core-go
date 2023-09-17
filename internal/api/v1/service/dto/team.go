package dto

import "subflow-core-go/pkg/ent"

type TeamInfo struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Status  int    `json:"status"`
	QQGroup string `json:"QQGroup"`
	Logo    string `json:"logo"`
	Desc    string `json:"desc"`
}

func GetTeamInfoFromEntities(teams []*ent.Team) []*TeamInfo {
	var teamInfos []*TeamInfo
	for _, t := range teams {
		teamInfos = append(teamInfos, GetTeamInfoFromEntity(t))
	}
	return teamInfos
}

func GetTeamInfoFromEntity(t *ent.Team) *TeamInfo {
	return &TeamInfo{
		ID:      t.ID,
		Name:    t.Name,
		Status:  t.Status,
		QQGroup: t.QqGroup,
		Logo:    t.Logo,
		Desc:    t.Desc,
	}
}
