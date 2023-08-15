package middleware

import (
	"errors"
	"strconv"

	"github.com/casbin/casbin/v2"
	_ "github.com/go-sql-driver/mysql"
	casbinware "github.com/gofiber/contrib/casbin"
	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
	"subflow-core-go/internal/api/common"
	"subflow-core-go/internal/api/helper"
	"subflow-core-go/internal/config"
)

func NewCasbinEnforcer(cfg *config.Config, a *helper.CasbinAdapter) (*casbin.Enforcer, error) {
	enforcer, err := casbin.NewEnforcer("config/casbin_model.conf", a)
	if err != nil {
		return nil, err
	}
	return enforcer, nil
}

func NewCasbinMiddleware(enforcer *casbin.Enforcer) (*casbinware.Middleware, error) {

	middleware := casbinware.New(
		casbinware.Config{
			Enforcer: enforcer,
			Lookup: func(ctx *fiber.Ctx) string {
				claim, err := helper.GetClaimFromFiberCtx(ctx)
				if err != nil {
					zap.S().Info("Unable to get claim for casbin")
					return ""
				}
				roles, _ := enforcer.GetRolesForUser(strconv.Itoa(claim.UID))
				zap.S().Debugw(
					"User roles list",
					"uid", claim.UID,
					"roles", roles,
				)
				return strconv.Itoa(claim.UID)
			},
			Unauthorized: func(ctx *fiber.Ctx) error {
				zap.S().Info("Casbin unauthorized")
				return &common.BusinessError{
					Code: common.ResultUnauthorized,
				}
			},
			Forbidden: func(ctx *fiber.Ctx) error {
				zap.S().Info("Casbin forbidden")
				return &common.BusinessError{
					Code: common.ResultPermissionDenied,
				}
			},
		},
	)
	if middleware == nil {
		return nil, errors.New("create casbin middleware failed")
	}
	return middleware, nil
}
