package dto

import (
	"time"

	"subflow-core-go/pkg/ent"
)

type CreateUserRequest struct {
	Username   string `json:"username" validate:"required"`
	Password   string `json:"password" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	RemoteAddr string `swaggerignore:"true"`
}

type UpdateUserReq struct {
	ID       int
	Nickname string
}

type UserBasicInfo struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

type UserInfo struct {
	*UserBasicInfo
	RegisterTime *time.Time `json:"registerTime"`
	RegisterIP   string     `json:"registerIP"`
	LoginTime    *time.Time `json:"loginTime"`
	LoginIP      string     `json:"loginIP"`
}

func GetBasicInfoFromUser(u *ent.User) *UserBasicInfo {
	if u == nil {
		return nil
	}
	return &UserBasicInfo{
		ID:       u.ID,
		Username: u.Username,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
	}
}

func GetBasicInfoFromUsers(us []*ent.User) []*UserBasicInfo {
	infos := make([]*UserBasicInfo, len(us))
	for i, u := range us {
		infos[i] = GetBasicInfoFromUser(u)
	}
	return infos
}

func GetInfoFromUser(u *ent.User) *UserInfo {
	if u == nil {
		return nil
	}
	return &UserInfo{
		UserBasicInfo: GetBasicInfoFromUser(u),
		RegisterTime:  &u.RegisteredAt,
		RegisterIP:    u.RegisterIP,
		LoginTime:     &u.LastLoggedAt,
		LoginIP:       u.LoginIP,
	}
}
