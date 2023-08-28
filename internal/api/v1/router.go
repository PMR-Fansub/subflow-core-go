package v1

import (
	"net/http"

	casbinware "github.com/gofiber/contrib/casbin"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"subflow-core-go/internal/api/common"
	"subflow-core-go/internal/api/constants"
	"subflow-core-go/internal/api/helper"
	"subflow-core-go/internal/api/v1/handler"
	"subflow-core-go/internal/config"

	_ "subflow-core-go/docs"
)

type Router struct {
	router           fiber.Router
	handler          *handler.Handler
	casbinMiddleware *casbinware.Middleware
}

var jwtMiddleware fiber.Handler

func jwtErrorHandler(ctx *fiber.Ctx, err error) error {
	result := common.ResultUnauthorized
	resp := common.MakeAPIResponse(result, err.Error())
	return ctx.Status(result.HttpCode).JSON(resp)
}

func NewRouter(app *fiber.App, handlerV1 *handler.Handler, cfg *config.Config, cm *casbinware.Middleware) *Router {
	jwtMiddleware = jwtware.New(
		jwtware.Config{
			SigningMethod: "HS256",
			SigningKey:    []byte(cfg.Server.SigningKey),
			ErrorHandler:  jwtErrorHandler,
			TokenLookup:   "cookie:auth_token,header:Authorization",
			Claims:        &helper.UserClaim{},
		},
	)
	return &Router{
		router:           app.Group("v1"),
		handler:          handlerV1,
		casbinMiddleware: cm,
	}
}
func (r *Router) Setup() {
	r.router.Get("healthz", health)
	r.router.Get("metrics", adaptor.HTTPHandler(promhttp.Handler()))
	r.router.Get("/swagger/*", fiberSwagger.WrapHandler)

	r.SetupAuth()
	r.SetupUser()
	r.SetupTeam()
}

func (r *Router) SetupAuth() {
	auth := r.router.Group("auth")
	auth.Post("register", handler.WrapHandlerWithAutoParse(r.handler.Register))
	auth.Post("login", handler.WrapHandlerWithAutoParse(r.handler.Login))
}

func (r *Router) SetupUser() {
	userGrp := r.router.Group("user")
	userGrp.Get("/:id", handler.WrapHandlerWithAutoParse(r.handler.GetUserByID))
	userGrp.Get("/:id/teams", handler.WrapHandlerWithAutoParse(r.handler.GetUserTeamsByID))

	userGrpWithAuth := r.router.Group("user", jwtMiddleware)
	userGrpWithAuth.Get("/", handler.WrapHandlerWithAutoParse(r.handler.GetCurrentUser))
	userGrpWithAuth.Patch("/", handler.WrapHandlerWithAutoParse(r.handler.UpdateCurrentUser))
	userGrpWithAuth.Patch(
		"/:id",
		r.casbinMiddleware.RequiresRoles(
			[]string{constants.RoleNameAdmin, constants.RoleNameSuperuser},
			casbinware.WithValidationRule(casbinware.AtLeastOneRule),
		),
		handler.WrapHandlerWithAutoParse(r.handler.UpdateUser),
	)
}

func (r *Router) SetupTeam() {
	// teamGrp := r.router.Group("team")
}

func health(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}
