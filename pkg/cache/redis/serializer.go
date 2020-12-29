package redis

type ISerializer interface {
	Serialization(value interface{}) ([]byte, error)
	Deserialization(byt []byte, ptr interface{}) (err error)
}
