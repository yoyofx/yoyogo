package Abstractions

type IDataSource interface {
	Open() interface{}
	Close()
	Ping() bool
}
