// Package bootstrap
package bootstrap

import (
	"shop-se/internal/appctx"
	"shop-se/pkg/logger"
	"shop-se/pkg/util"
)

func RegistryLogger(cfg *appctx.Config) {
	logger.Setup(logger.Config{
		Environment: util.EnvironmentTransform(cfg.App.Env),
		Debug:       cfg.App.Debug,
		Level:       cfg.Logger.Level,
		ServiceName: cfg.Logger.Name,
	})
}
