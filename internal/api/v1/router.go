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
	endpointName := "auth"
	auth := r.router.Group(endpointName)
	auth.Post("register", handler.Wrap(r.handler.Register))
	auth.Post("login", handler.Wrap(r.handler.Login))
}

func (r *Router) SetupUser() {
	endpointName := "users"
	grp := r.router.Group(endpointName)
	grp.Get("/:id", handler.Wrap(r.handler.GetUserByID))
	grp.Get("/:id/teams", handler.Wrap(r.handler.GetUserTeamsByID))

	grpWithAuth := r.router.Group(endpointName, r.jwtMiddleware)
	grpWithAuth.Get("/", handler.Wrap(r.handler.GetCurrentUser))
	grpWithAuth.Patch("/", handler.Wrap(r.handler.UpdateCurrentUser))
	grpWithAuth.Patch(
		"/:id",
		r.casbinMiddleware.RequiresRoles(
			[]string{constants.RoleNameAdmin, constants.RoleNameSuperuser},
			casbinware.WithValidationRule(casbinware.AtLeastOneRule),
		),
		handler.Wrap(r.handler.UpdateUser),
	)
}

func (r *Router) SetupTeam() {
	endpointName := "teams"
	grp := r.router.Group(endpointName)
	grp.Get("/", handler.Wrap(r.handler.GetAllTeams))
	grp.Get("/:id", handler.Wrap(r.handler.GetTeamByID))
	grp.Get("/:id/users", handler.Wrap(r.handler.GetTeamUsersByID))
	grp.Get("/:id/tasks", handler.Wrap(r.handler.GetTeamTasksByID))

	grpWithAuth := r.router.Group(endpointName, r.jwtMiddleware)
	grpWithAuth.Post("/", handler.Wrap(r.handler.CreateNewTeam))          // TODO: more access control
	grpWithAuth.Patch("/:id", handler.Wrap(r.handler.UpdateTeamInfoByID)) // TODO: more access control
	grpWithAuth.Post("/:id/users", handler.Wrap(r.handler.AddUserToTeam)) // TODO: more access control
}

func (r *Router) SetupWorkFlow() {
	// endpointName := "workflows"
	// grp := r.router.Group(endpointName)
	// grp.Get("/", handler.Wrap(r.handler.GetAllWorkflows))
	// grp.Get("/:id", handler.Wrap(r.handler.GetWorkflowByID))
	// grp.Get("/:id/nodes", handler.Wrap(r.handler.GetWorkflowNodesByID))

	// grpWithAuth := r.router.Group(endpointName, r.jwtMiddleware)
	// grpWithAuth.Post("/", handler.Wrap(r.handler.CreateWorkflow))
	// grpWithAuth.Patch("/:id", handler.Wrap(r.handler.UpdateWorkflowByID))
}

func (r *Router) SetupTask() {
	// endpointName := "tasks"
	// grp := r.router.Group(endpointName)
	// grp.Get("/:id", r.handler.GetTaskByID)
	// grp.Get("/:id/records", r.handler.GetTaskRecordsByID)
}

func health(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}
