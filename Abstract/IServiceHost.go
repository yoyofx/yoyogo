package Abstract

type IServiceHost interface {
	Run()
	Shutdown()
	StopApplicationNotify()
	SetAppMode(mode string)
}
