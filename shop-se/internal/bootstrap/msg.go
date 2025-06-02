// Package bootstrap
package bootstrap

import (
	"shop-se/internal/consts"
	"shop-se/pkg/logger"
	"shop-se/pkg/msgx"
)

func RegistryMessage() {
	err := msgx.Setup("msg.yaml", consts.ConfigPath)
	if err != nil {
		logger.Fatal(logger.MessageFormat("file message multi language load error %s", err.Error()))
	}

}
