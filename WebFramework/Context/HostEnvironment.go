package Context

const (
	Dev  = "Dev"
	Prod = "Prod"
	Test = "Test"
)

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
	return env.Profile == Dev
}

func (env HostEnvironment) IsStaging() bool {
	return env.Profile == Test
}

func (env HostEnvironment) IsProduction() bool {
	return env.Profile == Prod
}
