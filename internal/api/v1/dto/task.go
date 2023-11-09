package dto

import (
	"time"

	"subflow-core-go/pkg/ent"
)

type TaskInfo struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	Type        int           `json:"type"`
	Status      int           `json:"status"`
	Desc        string        `json:"desc"`
	CreatedAt   time.Time     `json:"createdAt"`
	CompletedAt time.Time     `json:"completedAt"`
	Tags        []TaskTagInfo `json:"tags"`
}

func GetTaskInfoFromEntity(t *ent.Task) TaskInfo {
	return TaskInfo{
		ID:          t.ID,
		Name:        t.Name,
		Type:        t.Type,
		Status:      t.Status,
		Desc:        t.Desc,
		CreatedAt:   t.CreatedAt,
		CompletedAt: t.CompletedAt,
		Tags:        GetTaskTagInfoFromEntities(t.Edges.TaskTags),
	}
}

func GetTaskInfoFromEntities(ts ent.Tasks) []TaskInfo {
	infos := make([]TaskInfo, len(ts))
	for i, t := range ts {
		infos[i] = GetTaskInfoFromEntity(t)
	}
	return infos
}
