package v1

import (
	"net/http"

	casbinware "github.com/gofiber/contrib/casbin"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	_ "subflow-core-go/docs"
	"subflow-core-go/internal/api/constants"
	"subflow-core-go/internal/api/v1/handler"
)

type Router struct {
	router           fiber.Router
	handler          *handler.Handler
	casbinMiddleware *casbinware.Middleware
	jwtMiddleware    fiber.Handler
}

func NewRouter(app *fiber.App, handlerV1 *handler.Handler, cm *casbinware.Middleware, jm fiber.Handler) *Router {
	return &Router{
		router:           app.Group("v1"),
		handler:          handlerV1,
		casbinMiddleware: cm,
		jwtMiddleware:    jm,
	}
}
func (r *Router) Setup() {
	r.SetupHealthCheck()
	r.SetupMetrics()
	r.SetupSwagger()

	r.SetupAuth()
	r.SetupUser()
	r.SetupTeam()
	r.SetupWorkFlow()
	r.SetupTask()
}

func (r *Router) SetupSwagger() fiber.Router {
	return r.router.Get("/swagger/*", fiberSwagger.WrapHandler)
}

func (r *Router) SetupMetrics() fiber.Router {
	return r.router.Get("metrics", adaptor.HTTPHandler(promhttp.Handler()))
}

func (r *Router) SetupHealthCheck() fiber.Router {
	return r.router.Get("healthz", health)
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

	userGrpWithAuth := r.router.Group("user", r.jwtMiddleware)
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

func (r *Router) SetupWorkFlow() {

}

func (r *Router) SetupTask() {

}

func health(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}
