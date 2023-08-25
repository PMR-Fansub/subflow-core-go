package service

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"subflow-core-go/internal/api/common"
	"subflow-core-go/internal/api/constants"
	"subflow-core-go/internal/api/v1/service/dto"
	"subflow-core-go/pkg/ent"
	"subflow-core-go/pkg/ent/user"
)

func GetBasicInfoFromUser(u *ent.User) *dto.UserBasicInfo {
	if u == nil {
		return nil
	}
	return &dto.UserBasicInfo{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
	}
}

func GetInfoFromUser(u *ent.User) *dto.UserInfo {
	if u == nil {
		return nil
	}
	return &dto.UserInfo{
		UserBasicInfo: GetBasicInfoFromUser(u),
		RegisterTime:  &u.RegisteredAt,
		RegisterIP:    u.RegisterIP,
		LoginTime:     &u.LastLoggedAt,
		LoginIP:       u.LoginIP,
	}
}

func (s *Service) FindUserByID(ctx context.Context, id int) (*ent.User, error) {
	u, err := s.db.User.
		Query().
		Where(user.ID(id)).
		Only(ctx)
	if err != nil && ent.IsNotFound(err) {
		return nil, &common.BusinessError{
			Code:    common.ResultNotFound,
			Message: "用户不存在",
		}
	}
	return u, err
}

func (s *Service) FindUserByUsername(ctx context.Context, username string) (*ent.User, error) {
	u, err := s.db.User.
		Query().
		Where(user.Username(username)).
		Only(ctx)
	if err != nil && ent.IsNotFound(err) {
		return nil, &common.BusinessError{
			Code:    common.ResultNotFound,
			Message: "用户不存在",
		}
	}
	return u, err
}

func (s *Service) FindUserByEmail(ctx context.Context, email string) (*ent.User, error) {
	u, err := s.db.User.
		Query().
		Where(user.Email(email)).
		Only(ctx)
	if err != nil && ent.IsNotFound(err) {
		return nil, &common.BusinessError{
			Code:    common.ResultNotFound,
			Message: "用户不存在",
		}
	}
	return u, err
}

func (s *Service) CreateUser(ctx context.Context, req *dto.CreateUserRequest) (*ent.User, error) {
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
		SetNickname(req.Username).
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

func (s *Service) RefreshLastLoginTimeAndIP(ctx context.Context, u *ent.User, t time.Time, ip string) error {
	return u.Update().
		SetLastLoggedAt(t).
		SetLoginIP(ip).
		Exec(ctx)
}

func (s *Service) UpdateUser(ctx context.Context, req *dto.UpdateUserReq) error {
	if len(req.Nickname) > 0 {
		err := s.db.User.
			UpdateOneID(req.ID).
			SetNickname(req.Nickname).
			Exec(ctx)
		switch {
		case ent.IsNotFound(err):
			return &common.BusinessError{
				Code:    common.ResultNotFound,
				Message: fmt.Sprintf("用户不存在 (UID: %d)", req.ID),
			}
		case err != nil:
			return err
		}
	}

	return nil
}
