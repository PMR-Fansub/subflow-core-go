package v1

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"subflow-core-go/internal/api/v1/handler"
)

type Router struct {
	router  fiber.Router
	handler *handler.Handler
}

func NewRouter(app *fiber.App, handlerV1 *handler.Handler) *Router {
	return &Router{
		router:  app.Group("v1"),
		handler: handlerV1,
	}
}
func (r *Router) Setup() {
	r.router.Get("healthz", health)
	r.router.Get("metrics", adaptor.HTTPHandler(promhttp.Handler()))

	r.SetupAuth()
	r.SetupUser()
}

func (r *Router) SetupAuth() {
	auth := r.router.Group("auth")
	auth.Post("register", handler.AutoParsedHandler(r.handler.Register))
}

func (r *Router) SetupUser() {
	// user := r.router.Group("user")
	// user.Get("/", r.handler.GetAllUser)
	// user.Post("/", r.handler.CreateUser)
}

func health(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}
