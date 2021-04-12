package services

type IOCDemo struct {
}

func NewIOCDemo() *IOCDemo {
	return &IOCDemo{}
}

func (demo *IOCDemo) Print() string {
	return "IOC Demo test"
}
