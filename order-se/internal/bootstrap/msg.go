// Package bootstrap
package bootstrap

import (
	"order-se/internal/consts"
	"order-se/pkg/logger"
	"order-se/pkg/msgx"
)

func RegistryMessage() {
	err := msgx.Setup("msg.yaml", consts.ConfigPath)
	if err != nil {
		logger.Fatal(logger.MessageFormat("file message multi language load error %s", err.Error()))
	}

}
