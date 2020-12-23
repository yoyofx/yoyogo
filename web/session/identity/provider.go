package identity

type IProvider interface {
	SetID(sessionId string)
	GetID() string
	Clear()
	SetMaxLifeTime(liftTime int64)
	SetContext(context interface{})
}
