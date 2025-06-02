// Package migration
package migration

import (
	"warehouse-se/internal/appctx"
	"warehouse-se/pkg/databasex"
)

func MigrateDatabase() {
	cfg := appctx.NewConfig()

	databasex.DatabaseMigration(&databasex.Config{
		Driver:       cfg.WriteDB.Driver,
		Host:         cfg.WriteDB.Host,
		Port:         cfg.WriteDB.Port,
		Name:         cfg.WriteDB.Name,
		User:         cfg.WriteDB.User,
		Password:     cfg.WriteDB.Pass,
		Charset:      cfg.WriteDB.Charset,
		MaxIdleConns: cfg.WriteDB.MaxIdle,
		MaxOpenConns: cfg.WriteDB.MaxOpen,
		MaxLifetime:  cfg.WriteDB.MaxLifeTime,
		TimeZone:     cfg.WriteDB.Timezone,
	})
}
