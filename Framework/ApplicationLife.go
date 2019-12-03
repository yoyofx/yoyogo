package YoyoGo

const (
	APPLICATION_LIFE_START = "application_life_start"
	APPLICATION_LIFE_STOP  = "application_life_stop"
)

//var ApplicationCycle = NewApplicationLife()

type ApplicationLife struct {
	eventPublisher     *ApplicationEventPublisher
	ApplicationStopped chan ApplicationEvent
	ApplicationStarted chan ApplicationEvent
}

func NewApplicationLife() *ApplicationLife {
	applife := &ApplicationLife{
		eventPublisher:     NewEventPublisher(),
		ApplicationStopped: make(chan ApplicationEvent),
		ApplicationStarted: make(chan ApplicationEvent),
	}
	applife.eventPublisher.Subscribe(APPLICATION_LIFE_START, applife.ApplicationStarted)
	applife.eventPublisher.Subscribe(APPLICATION_LIFE_STOP, applife.ApplicationStopped)
	return applife
}

func (life *ApplicationLife) StartApplication() {
	life.eventPublisher.Publish(APPLICATION_LIFE_START, "start")
}

func (life *ApplicationLife) StopApplication() {
	life.eventPublisher.Publish(APPLICATION_LIFE_STOP, "stop")
}
