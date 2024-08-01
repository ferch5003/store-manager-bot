package router

import (
	"backend/cmd/api/handler"
	"backend/config"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var NewHistoryModule = fx.Module("user",
	// Register Handler
	fx.Provide(handler.NewHistoryHandler),

	// Register Router
	fx.Provide(
		fx.Annotate(
			NewUserRouter,
			fx.ResultTags(`group:"routers"`),
		),
	),
)

type historyRouter struct {
	App     fiber.Router
	config  *config.EnvVars
	Handler *handler.HistoryHandler
}

func NewUserRouter(app *fiber.App,
	config *config.EnvVars,
	historyHandler *handler.HistoryHandler) Router {
	return &historyRouter{
		App:     app,
		config:  config,
		Handler: historyHandler,
	}
}

func (h historyRouter) Register() {
	h.App.Use("/histories", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	h.App.Route("/histories", func(api fiber.Router) {
		api.Get("/", websocket.New(h.Handler.HistoryChat)).Name("get")
	}, "users.")
}
