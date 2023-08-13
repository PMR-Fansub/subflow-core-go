package helper

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"subflow-core-go/internal/api/common"
)

type UserClaim struct {
	UID      int    `json:"uid"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func SignJWT(signingKey string, claim *UserClaim) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
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
	claim, ok := user.Claims.(*UserClaim)
	if !ok {
		zap.S().Error("Cannot assert custom claim type")
		return nil, &common.BusinessError{
			Code: common.ResultUnauthorized,
		}
	}
	return claim, nil
}
