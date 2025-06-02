// Package bootstrap
package bootstrap

import (
	"warehouse-se/internal/appctx"
	"warehouse-se/pkg/logger"
	"warehouse-se/pkg/util"
)

func RegistryLogger(cfg *appctx.Config) {
	logger.Setup(logger.Config{
		Environment: util.EnvironmentTransform(cfg.App.Env),
		Debug:       cfg.App.Debug,
		Level:       cfg.Logger.Level,
		ServiceName: cfg.Logger.Name,
	})
}
