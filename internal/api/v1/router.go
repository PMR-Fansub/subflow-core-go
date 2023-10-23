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
	auth.Post("register", handler.WithAutoParse(r.handler.Register))
	auth.Post("login", handler.WithAutoParse(r.handler.Login))
}

func (r *Router) SetupUser() {
	grp := r.router.Group("users")
	grp.Get("/:id", handler.WithAutoParse(r.handler.GetUserByID))
	grp.Get("/:id/teams", handler.WithAutoParse(r.handler.GetUserTeamsByID))

	grpWithAuth := r.router.Group("users", r.jwtMiddleware)
	grpWithAuth.Get("/", handler.WithAutoParse(r.handler.GetCurrentUser))
	grpWithAuth.Patch("/", handler.WithAutoParse(r.handler.UpdateCurrentUser))
	grpWithAuth.Patch(
		"/:id",
		r.casbinMiddleware.RequiresRoles(
			[]string{constants.RoleNameAdmin, constants.RoleNameSuperuser},
			casbinware.WithValidationRule(casbinware.AtLeastOneRule),
		),
		handler.WithAutoParse(r.handler.UpdateUser),
	)
}

func (r *Router) SetupTeam() {
	grp := r.router.Group("teams")
	grp.Get("/", handler.WithAutoParse(r.handler.GetAllTeams))
	grp.Get("/:id", handler.WithAutoParse(r.handler.GetTeamByID))
	grp.Get("/:id/users", handler.WithAutoParse(r.handler.GetTeamUsersByID))
	grp.Get("/:id/tasks", handler.WithAutoParse(r.handler.GetTeamTasksByID))

	grpWithAuth := r.router.Group("teams", r.jwtMiddleware)
	grpWithAuth.Post("/", handler.WithAutoParse(r.handler.CreateNewTeam))          // TODO: more access control
	grpWithAuth.Patch("/:id", handler.WithAutoParse(r.handler.UpdateTeamInfoByID)) // TODO: more access control
	grpWithAuth.Post("/:id/users", handler.WithAutoParse(r.handler.AddUserToTeam)) // TODO: more access control
}

func (r *Router) SetupWorkFlow() {
	// grp := r.router.Group("workflows")
	// grp.Get("/", handler.WithAutoParse(r.handler.GetAllWorkflows))
	// grp.Get("/:id", handler.WithAutoParse(r.handler.GetWorkflowByID))
	// grp.Get("/:id/nodes", handler.WithAutoParse(r.handler.GetWorkflowNodesByID))

	// grpWithAuth := r.router.Group("workflows", r.jwtMiddleware)
	// grpWithAuth.Post("/", handler.WithAutoParse(r.handler.CreateWorkflow))
	// grpWithAuth.Patch("/:id", handler.WithAutoParse(r.handler.UpdateWorkflowByID))
}

func (r *Router) SetupTask() {
	// grp := r.router.Group("tasks")
	// grp.Get("/:id", r.handler.GetTaskByID)
	// grp.Get("/:id/records", r.handler.GetTaskRecordsByID)
}

func health(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}
