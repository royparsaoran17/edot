// Package bootstrap
package bootstrap

import (
	"warehouse-se/internal/consts"
	"warehouse-se/pkg/logger"
	"warehouse-se/pkg/msgx"
)

func RegistryMessage() {
	err := msgx.Setup("msg.yaml", consts.ConfigPath)
	if err != nil {
		logger.Fatal(logger.MessageFormat("file message multi language load error %s", err.Error()))
	}

}
