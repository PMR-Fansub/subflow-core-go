package service

import (
	"context"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
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

func (s *Service) CreateUser(ctx context.Context, req *CreateUserRequest) (*ent.User, error) {
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

	encryptedPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		zap.S().Errorw(
			"Failed to crypt password",
			"req", req,
			"err", err,
		)
		return nil, err
	}
	return s.db.User.
		Create().
		SetUsername(req.Username).
		SetPassword(string(encryptedPwd)).
		SetEmail(req.Email).
		SetStatus(int(constants.UserStatusActive)).
		SetRegisterIP(req.RemoteAddr).
		Save(ctx)
}

func (s *Service) VerifyPwdByUsername(ctx context.Context, username string, pwd string) (*ent.User, error) {
	u, err := s.FindUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pwd))
	if err != nil {
		return nil, &common.BusinessError{
			Code: common.ResultUnauthorized,
		}
	}
	return u, nil
}
