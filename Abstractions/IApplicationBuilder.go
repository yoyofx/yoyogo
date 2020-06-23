package Abstractions

type IApplicationBuilder interface {
	Build() interface{}
	SetHostBuildContext(*HostBuildContext)
}
