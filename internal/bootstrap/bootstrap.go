package bootstrap

import (
	"context"

	"go.uber.org/fx"
	"honnef.co/go/tools/config"
)

var Module = fx.Options(
	util.Module,
	handlers.Module,
	// routes.Module,
	// services.Module,
	// repositories.Module,
	// middlewares.Module,
	fx.Invoke(register),
)

func register(
	lifecycle fx.Lifecycle,
	cfg *config.Config,
	logger *util.Logger,
	handler *util.RequestHandler,
	// route routes.Routes,
	// middleware middlewares.Middlewares,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Info("Starting Application")
			logger.Info("---------------------")
			logger.Info("------- CLEAN -------")
			logger.Info("---------------------")

			go func() {
				middleware.Setup()
				route.Setup()
				if err := handler.App.Run(":" + cfg.ServerPort); err != nil {
					logger.Errorf("Failed to run server on port %s", cfg.ServerPort)
					panic(err)
				}
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			logger.Info("Stopping Application")
			//if err := logger.SugaredLogger.Sync(); err != nil {
			//	logger.Errorf("flushing loggers, err: %v", err)
			//}
			return nil
		},
	})
}
