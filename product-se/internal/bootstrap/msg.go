// Package bootstrap
package bootstrap

import (
	"product-se/internal/consts"
	"product-se/pkg/logger"
	"product-se/pkg/msgx"
)

func RegistryMessage() {
	err := msgx.Setup("msg.yaml", consts.ConfigPath)
	if err != nil {
		logger.Fatal(logger.MessageFormat("file message multi language load error %s", err.Error()))
	}

}
