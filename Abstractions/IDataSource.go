package Abstractions

type IDataSource interface {
	GetName() string
	Open() (conn interface{}, put func(), err error)
	Close()
	Ping() bool
}
