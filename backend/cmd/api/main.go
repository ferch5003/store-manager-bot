package main

import (
	"backend/cmd/api/bootstrap"
	"backend/cmd/api/router"
	"backend/config"
	"backend/internal/platform/files"
	"backend/internal/platform/postgresql"
	"backend/internal/platform/vertex"
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"log"
)

func main() {
	configurations, err := config.NewConfigurations()
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalln(err)
	}

	credentials, err := files.GetFile("config/client_secret.json")
	if err != nil {
		log.Fatalln(err)
	}

	vertexClient, err := vertex.NewVertexClient(
		ctx, configurations.ProjectID,
		configurations.LocationID,
		configurations.EndpointID,
		credentials)
	if err != nil {
		log.Fatalln(err)
	}

	app := fx.New(
		// creates: config.EnvVars
		fx.Supply(configurations),
		// creates: *zap.Logger
		fx.Supply(logger),
		// creates: *fiber.Router
		fx.Provide(
			fx.Annotate(
				router.NewRouter,
				fx.ParamTags( // Equivalent to *fiber.App, config.Envars, []Router `group:"routers"` in constructor
					``,
					``,
					`group:"routers"`),
			),
		),
		// creates: *fiber.App
		fx.Provide(bootstrap.NewFiberServer),
		// creates: context.Context
		fx.Provide(context.Background),

		// creates: *sqlx.DB
		fx.Provide(postgresql.NewConnection),

		// creates: *vertx.Client
		fx.Supply(vertexClient),

		// Provide modules
		router.NewHistoryModule,

		// Start web server.
		fx.Invoke(bootstrap.Start),
	)

	app.Run()
}
