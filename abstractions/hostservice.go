package abstractions

type IHostService interface {
	Run() error
	Stop() error
}
