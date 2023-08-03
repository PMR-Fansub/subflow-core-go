package service

import (
	"context"

	"subflow-core-go/internal/api/common"
	"subflow-core-go/internal/api/constants"
	"subflow-core-go/pkg/ent"
	"subflow-core-go/pkg/ent/user"
)

type CreateUserRequest struct {
	Username   string `json:"username" validate:"required"`
	Password   string `json:"password" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	RemoteAddr string
}

func (s *Service) FindUserByUsername(ctx context.Context, username string) (*ent.User, error) {
	return s.db.User.
		Query().
		Where(user.Username(username)).
		Only(ctx)
}

func (s *Service) FindUserByEmail(ctx context.Context, email string) (*ent.User, error) {
	return s.db.User.
		Query().
		Where(user.Email(email)).
		Only(ctx)
}

func (s *Service) CreateUser(ctx context.Context, req CreateUserRequest) (*ent.User, error) {
	if u, _ := s.FindUserByUsername(ctx, req.Username); u != nil {
		return nil, &common.BusinessError{
			Code:    common.ResultCreateUserFailed,
			Message: "用户名已存在",
		}
	}
	if u, _ := s.FindUserByEmail(ctx, req.Email); u != nil {
		return nil, &common.BusinessError{
			Code:    common.ResultCreateUserFailed,
			Message: "邮箱已被使用",
		}
	}
	return s.db.User.
		Create().
		SetUsername(req.Username).
		SetPassword(req.Password).
		SetEmail(req.Email).
		SetStatus(int(constants.UserStatusActive)).
		SetRegisterIP(req.RemoteAddr).
		Save(ctx)
}
