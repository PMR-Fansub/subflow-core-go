package helper

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"subflow-core-go/internal/api/common"
)

type UserClaim struct {
	UID      int    `json:"uid"`
	Username string `json:"username"`
	Exp      int64  `json:"exp"`
}

func SignJWT(signingKey string, claim *UserClaim) (string, error) {
	claimBytes, err := json.Marshal(claim)
	if err != nil {
		zap.S().Errorw(
			"Marshal failed",
			"err", err,
		)
		return "", err
	}
	var claimMap jwt.MapClaims
	err = json.Unmarshal(claimBytes, &claimMap)
	if err != nil {
		zap.S().Errorw(
			"Unmarshal failed",
			"err", err,
		)
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimMap)
	t, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

func GetClaimFromFiberCtx(c *fiber.Ctx) (*UserClaim, error) {
	user, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		zap.S().Error("Cannot get jwt claim from context")
		return nil, &common.BusinessError{
			Code: common.ResultUnauthorized,
		}
	}
	claimMap, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		zap.S().Error("Cannot get jwt claim from context")
		return nil, &common.BusinessError{
			Code: common.ResultUnauthorized,
		}
	}

	var claim UserClaim
	bytes, err := json.Marshal(claimMap)
	if err != nil {
		zap.S().Errorw(
			"Marshal error",
			"err", err,
		)
		return nil, err
	}
	err = json.Unmarshal(bytes, &claim)
	if err != nil {
		zap.S().Errorw(
			"Unmarshal error",
			"err", err,
		)
		return nil, err
	}
	return &claim, nil
}
