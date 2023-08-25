package dto

import (
	"time"
)

type CreateUserRequest struct {
	Username   string `json:"username" validate:"required"`
	Password   string `json:"password" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	RemoteAddr string
}

type UpdateUserReq struct {
	ID       int
	Nickname string
}

type UserBasicInfo struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
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
