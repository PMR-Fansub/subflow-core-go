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
	"subflow-core-go/internal/config"
	"subflow-core-go/pkg/ent"
	"subflow-core-go/pkg/ent/user"
)

type UserService interface {
	FindUserByID(ctx context.Context, id int) (*ent.User, error)
	FindUserByUsername(ctx context.Context, username string) (*ent.User, error)
	FindUserByEmail(ctx context.Context, email string) (*ent.User, error)
	CreateUser(ctx context.Context, req *dto.CreateUserRequest) (*ent.User, error)
	VerifyPwdByUsername(ctx context.Context, username string, password string) (*ent.User, error)
	RefreshLastLoginTimeAndIP(ctx context.Context, u *ent.User, t time.Time, ip string) error
	UpdateUser(ctx context.Context, req *dto.UpdateUserReq) error
	GetTeamsOfUser(ctx context.Context, u *ent.User) (ent.Teams, error)
}

type UserServiceImpl struct {
	db     *ent.Client
	config *config.Config
}

func NewUserService(db *ent.Client, config *config.Config) UserService {
	return &UserServiceImpl{
		db,
		config,
	}
}

func (s *UserServiceImpl) FindUserByID(ctx context.Context, id int) (*ent.User, error) {
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

func (s *UserServiceImpl) FindUserByUsername(ctx context.Context, username string) (*ent.User, error) {
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

func (s *UserServiceImpl) FindUserByEmail(ctx context.Context, email string) (*ent.User, error) {
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

func (s *UserServiceImpl) CreateUser(ctx context.Context, req *dto.CreateUserRequest) (*ent.User, error) {
	if u, _ := s.FindUserByUsername(ctx, req.Username); u != nil {
		return nil, &common.BusinessError{
			Code:    common.ResultCreationFailed,
			Message: "用户名已存在",
		}
	}
	if u, _ := s.FindUserByEmail(ctx, req.Email); u != nil {
		return nil, &common.BusinessError{
			Code:    common.ResultCreationFailed,
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

func (s *UserServiceImpl) VerifyPwdByUsername(ctx context.Context, username string, pwd string) (*ent.User, error) {
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

func (s *UserServiceImpl) RefreshLastLoginTimeAndIP(ctx context.Context, u *ent.User, t time.Time, ip string) error {
	return u.Update().
		SetLastLoggedAt(t).
		SetLoginIP(ip).
		Exec(ctx)
}

func (s *UserServiceImpl) UpdateUser(ctx context.Context, req *dto.UpdateUserReq) error {
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

func (s *UserServiceImpl) GetTeamsOfUser(ctx context.Context, u *ent.User) (ent.Teams, error) {
	t, err := u.
		QueryTeams().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return t, nil
}
