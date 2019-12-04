package Context

const (
	Dev  = "Dev"
	Prod = "Prod"
	Test = "Test"
)

type HostEnvironment struct {
	ApplicationName string
	Version         string
	AppMode         string
	Args            []string
	Addr            string
	Port            string
	PID             int
}

func (env HostEnvironment) IsDevelopment() bool {
	return env.AppMode == Dev
}

func (env HostEnvironment) IsStaging() bool {
	return env.AppMode == Test
}

func (env HostEnvironment) IsProduction() bool {
	return env.AppMode == Prod
}
