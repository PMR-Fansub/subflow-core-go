package handler

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"subflow-core-go/internal/api/helper"
	"subflow-core-go/internal/api/v1/service"
	"subflow-core-go/internal/api/v1/service/dto"
)

type RegisterRequest struct {
	*dto.CreateUserRequest
}

type LoginRequest struct {
	Username   string `json:"username" validate:"required"`
	Password   string `json:"password" validate:"required"`
	WithCookie bool   `json:"withCookie" default:"false"`
}

type LoginResponse struct {
	Token    string             `json:"token"`
	UserInfo *dto.UserBasicInfo `json:"userInfo"`
}

type RegisterResponse struct{}

// Register godoc
//
//	@Summary	Register a new user
//	@Tags		auth
//	@Accept		json
//	@Produce	json
//	@Param		message	body		RegisterRequest	true	"user info to create"
//	@Success	200		{object}	common.APIResponse{data=RegisterResponse}
//	@Router		/auth/register [post]
func (h *Handler) Register(c *fiber.Ctx, req RegisterRequest) (*RegisterResponse, error) {
	req.RemoteAddr = c.IP()
	_, err := h.service.CreateUser(context.Background(), req.CreateUserRequest)
	return nil, err
}

// Login godoc
//
//	@Summary	User login
//	@Tags		auth
//	@Accept		json
//	@Produce	json
//	@Param		message	body		LoginRequest	true	"login form"
//	@Success	200		{object}	common.APIResponse{data=LoginResponse}
//	@Router		/auth/login [post]
func (h *Handler) Login(c *fiber.Ctx, req LoginRequest) (*LoginResponse, error) {
	var resp LoginResponse

	user, err := h.service.VerifyPwdByUsername(context.Background(), req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	err = h.service.RefreshLastLoginTimeAndIP(context.Background(), user, time.Now(), c.IP())
	if err != nil {
		return nil, err
	}

	token, err := helper.SignJWT(
		h.config.Server.SigningKey, &helper.UserClaim{
			UID:      user.ID,
			Username: user.Username,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * 24 * time.Hour)),
			},
		},
	)
	if err != nil {
		return nil, err
	}

	resp.Token = token
	resp.UserInfo = service.GetBasicInfoFromUser(user)

	if req.WithCookie {
		c.Cookie(
			&fiber.Cookie{
				Name:     "auth_token",
				Value:    token,
				HTTPOnly: true,
				Secure:   true,
			},
		)
	}

	return &resp, nil
}
