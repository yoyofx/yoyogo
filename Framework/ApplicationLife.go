package YoyoGo

type ApplicationLife struct {
	StopApplicationEvent  chan int
	StartApplicationEvent chan int
}

func (life ApplicationLife) StartApplication() {
	life.StartApplicationEvent <- 1
}

func (life ApplicationLife) StopApplication() {
	life.StopApplicationEvent <- 1
}
