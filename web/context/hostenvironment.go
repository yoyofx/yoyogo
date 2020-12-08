package context

import "github.com/yoyofx/yoyogo/abstractions/hostenv"

type HostEnvironment struct {
	ApplicationName string
	Version         string
	Profile         string
	Args            []string
	Addr            string
	Port            string
	Host            string
	PID             int
	MetaData        map[string]string
}

func (env HostEnvironment) IsDevelopment() bool {
	return env.Profile == hostenv.Dev
}

func (env HostEnvironment) IsStaging() bool {
	return env.Profile == hostenv.Test
}

func (env HostEnvironment) IsProduction() bool {
	return env.Profile == hostenv.Prod
}
