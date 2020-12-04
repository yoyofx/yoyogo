package Abstractions

type IDataSource interface {
	GetName() string
	Open() interface{}
	Close()
	Ping() bool
}
