package bootstrap

import (
	"time"

	"order-se/pkg/logger"
	"order-se/pkg/postgres"

	config "order-se/internal/appctx"
)

func RegistryPostgres(cfg *config.Database) postgres.Adapter {
	db, err := postgres.NewAdapter(&postgres.Config{
		Host:         cfg.Host,
		Name:         cfg.Name,
		Password:     cfg.Pass,
		Port:         cfg.Port,
		User:         cfg.User,
		Timeout:      time.Duration(cfg.TimeoutSecond) * time.Second,
		MaxOpenConns: cfg.MaxOpen,
		MaxIdleConns: cfg.MaxIdle,
		MaxLifetime:  time.Duration(cfg.MaxLifeTimeMS) * time.Millisecond,
		Timezone:     cfg.Timezone,
	})

	if err != nil {
		logger.Fatal(
			err,
			logger.EventName("db"),
			logger.Any("host", cfg.Host),
			logger.Any("port", cfg.Port),
		)
	}

	return db
}
