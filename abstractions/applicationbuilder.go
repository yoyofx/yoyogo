package abstractions

type IApplicationBuilder interface {
	Build() interface{}
	SetHostBuildContext(*HostBuilderContext)
}
