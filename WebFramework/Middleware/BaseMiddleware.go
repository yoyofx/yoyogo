package Middleware

import "github.com/yoyofx/yoyogo/Abstractions"

type IConfigurationMiddleware interface {
	SetConfiguration(config Abstractions.IConfiguration)
}

type BaseMiddleware struct {
	// Configuration
	config Abstractions.IConfiguration
}

func (mdw *BaseMiddleware) SetConfiguration(config Abstractions.IConfiguration) {
	mdw.config = config
}
