package v1

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"subflow-core-go/internal/api/common"
	"subflow-core-go/internal/api/v1/handler"
	"subflow-core-go/internal/config"
)

type Router struct {
	router  fiber.Router
	handler *handler.Handler
}

var jwtMiddleware fiber.Handler

func jwtErrorHandler(ctx *fiber.Ctx, err error) error {
	result := common.ResultUnauthorized
	resp := common.MakeAPIResponse(result, err.Error())
	return ctx.Status(result.HttpCode).JSON(resp)
}

func NewRouter(app *fiber.App, handlerV1 *handler.Handler, cfg *config.Config) *Router {
	jwtMiddleware = jwtware.New(
		jwtware.Config{
			SigningMethod: "HS256",
			SigningKey:    []byte(cfg.Server.SigningKey),
			ErrorHandler:  jwtErrorHandler,
			TokenLookup:   "cookie:auth_token,header:Authorization",
		},
	)
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
	auth.Post("register", handler.WrapHandlerWithAutoParse(r.handler.Register))
	auth.Post("login", handler.WrapHandlerWithAutoParse(r.handler.Login))
}

func (r *Router) SetupUser() {
	user := r.router.Group("user", jwtMiddleware)
	user.Get("/", handler.WrapHandlerWithAutoParse(r.handler.GetCurrentUser))
	user.Patch("/", handler.WrapHandlerWithAutoParse(r.handler.UpdateCurrentUser))
	user.Patch("/:id", handler.WrapHandlerWithAutoParse(r.handler.UpdateUser))
}

func health(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}
