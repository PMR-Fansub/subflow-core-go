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

func GetTeamInfoFromEntities(ts []*ent.Team) []TeamInfo {
	teamInfos := make([]TeamInfo, len(ts))
	for i, t := range ts {
		teamInfos[i] = GetTeamInfoFromEntity(t)
	}
	return teamInfos
}

func GetTeamInfoFromEntity(t *ent.Team) TeamInfo {
	return TeamInfo{
		ID:      t.ID,
		Name:    t.Name,
		Status:  t.Status,
		QQGroup: t.QqGroup,
		Logo:    t.Logo,
		Desc:    t.Desc,
	}
}
