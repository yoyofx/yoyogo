package YoyoGo

var ApplicationCycle = NewApplicationLife()

type ApplicationLife struct {
	StopApplicationEvent  chan ApplicationEvent
	StartApplicationEvent chan ApplicationEvent
}

func NewApplicationLife() *ApplicationLife {
	return &ApplicationLife{
		StopApplicationEvent:  make(chan ApplicationEvent),
		StartApplicationEvent: make(chan ApplicationEvent),
	}
}

func (life *ApplicationLife) StartApplication() {
	//life.StartApplicationEvent <- nil
}

func (life *ApplicationLife) StopApplication() {
	//life.StopApplicationEvent <- nil
}
