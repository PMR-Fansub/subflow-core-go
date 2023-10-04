package dto

import "subflow-core-go/pkg/ent"

type TaskTagInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func GetTaskTagInfoFromEntity(t *ent.TaskTag) *TaskTagInfo {
	return &TaskTagInfo{
		ID:   t.ID,
		Name: t.Name,
		Desc: t.Desc,
	}
}

func GetTaskTagInfoFromEntities(ts ent.TaskTags) []*TaskTagInfo {
	infos := make([]*TaskTagInfo, len(ts))
	for i, t := range ts {
		infos[i] = GetTaskTagInfoFromEntity(t)
	}
	return infos
}
