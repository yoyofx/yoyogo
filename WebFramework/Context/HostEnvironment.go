package Context

import "github.com/yoyofx/yoyogo/Abstractions/Env"

type HostEnvironment struct {
	ApplicationName string
	Version         string
	Profile         string
	Args            []string
	Addr            string
	Port            string
	Host            string
	PID             int
}

func (env HostEnvironment) IsDevelopment() bool {
	return env.Profile == Env.Dev
}

func (env HostEnvironment) IsStaging() bool {
	return env.Profile == Env.Test
}

func (env HostEnvironment) IsProduction() bool {
	return env.Profile == Env.Prod
}
