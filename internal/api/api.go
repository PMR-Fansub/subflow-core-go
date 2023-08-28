package api

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/contrib/fiberzap"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"subflow-core-go/internal/api/common"
	v1 "subflow-core-go/internal/api/v1"
	"subflow-core-go/internal/config"
)

const (
	AppName      = "SubFlow-Core"
	idleTimeout  = 10 * time.Second
	readTimeout  = 10 * time.Second
	writeTimeout = 10 * time.Second
)

//	@title		SubFlow API
//	@version	1.0

//	@license.name	GPLv3
//	@license.url	https://www.gnu.org/licenses/gpl-3.0.html

//	@host		localhost:8000
//	@BasePath	/v1

// @securityDefinitions.apiKey	ApiKeyAuth
// @In							header
// @Name						Authorization
func New() *fiber.App {
	return fiber.New(
		fiber.Config{
			AppName:      AppName,
			ErrorHandler: common.ErrorHandler,
			IdleTimeout:  idleTimeout,
			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout,
		},
	)
}

func Start(lifestyle fx.Lifecycle, app *fiber.App, cfg *config.Config, rv1 *v1.Router) {
	setupMiddlewares(app, cfg)

	rv1.Setup()

	lifestyle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				addr := fmt.Sprintf("%s:%d", cfg.Server.IpAddr, cfg.Server.Port)
				go func() {
					err := app.Listen(addr)
					if err != nil {
						zap.S().Warnw(
							"Failed to startup app",
							"err", err,
						)
					}
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return app.Shutdown()
			},
		},
	)
}

func setupMiddlewares(app *fiber.App, cfg *config.Config) {
	app.Use(
		cors.New(
			cors.Config{
				AllowCredentials: true,
				AllowOrigins:     cfg.Server.CookieHost,
			},
		),
	)
	app.Use(requestid.New())
	app.Use(
		fiberzap.New(
			fiberzap.Config{
				Logger: zap.L(),
				Fields: []string{"latency", "status", "method", "url", "requestId"},
			},
		),
	)
}
