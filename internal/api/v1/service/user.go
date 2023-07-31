package service

import (
	"context"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"subflow-core-go/internal/api/common"
	"subflow-core-go/pkg/ent"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GetAllUserRequest struct {
	PerPage int `json:"perPage"`
}

type GetAllUserResponse = []User

type CreateUserRequest struct {
	Username   string `json:"username" validate:"required"`
	Password   string `json:"password" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	RemoteAddr string
}

func (s *Service) GetAllUser(ctx context.Context, req GetAllUserRequest) (*GetAllUserResponse, error) {
	var users []User
	u, err := s.db.User.Query().Limit(req.PerPage).All(ctx)

	for _, user := range u {
		users = append(
			users, User{
				ID:   user.ID,
				Name: user.Username,
			},
		)
	}

	return &users, err
}

func (s *Service) CreateUser(ctx context.Context, req CreateUserRequest) (*ent.User, error) {
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
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
		SetPassword(string(pwdHash)).
		SetEmail(req.Email).
		SetStatus(int(common.UserStatusActive)).
		SetRegisterIP(req.RemoteAddr).
		Save(ctx)
}
