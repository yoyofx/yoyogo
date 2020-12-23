package identity

type IProvider interface {
	SetID(sessionId string)
	GetID() string
	Clear()
	SetName(name string)
	SetMaxLifeTime(liftTime int64)
	SetContext(context interface{})
}
