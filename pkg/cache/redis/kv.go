package redis

import "time"

type KV struct {
	ops        Ops
	serializer ISerializer
}

// Append append string,to the string
func (kv KV) Append(key string, value string) int64 {
	v, _ := kv.ops.Append(key, value)
	return v
}

// MultiGet is like Get but accepts multiple keys:
func (kv KV) MultiGet(keys ...string) ([]interface{}, error) {
	return kv.ops.MultiGet(keys...)
}

// MultiSet is like Set but accepts multiple values:
//   - MSet("key1", "value1", "key2", "value2")
//   - MSet([]string{"key1", "value1", "key2", "value2"})
//   - MSet(map[string]interface{}{"key1": "value1", "key2": "value2"})
func (kv KV) MultiSet(values ...interface{}) error {
	return kv.ops.MultiSet(values...)
}

//Get string for the key
func (kv KV) SetString(key string, value string, duration time.Duration) error {
	return kv.ops.Set(key, value, duration)
}

//Get get string for the key
func (kv KV) GetString(key string) (string, error) {
	return kv.ops.Get(key)
}

//Set set serialization object to the key
func (kv KV) Set(key string, value interface{}, duration time.Duration) error {
	ss, e := kv.serializer.Serialization(value)
	if e != nil {
		return e
	}
	return kv.ops.SetValue(key, ss, duration)
}

//Get get serialization object for the key
func (kv KV) Get(key string, ptr interface{}) error {
	dv, e := kv.ops.GetValue(key)
	if e != nil {
		return e
	}
	return kv.serializer.Deserialization(dv, ptr)
}

//GetAndSet get value for the key , and that if it's null, set value to the key.
func (kv KV) GetAndSet(key string, value string) (string, error) {
	v, e := kv.ops.Get(key)
	if e != nil {
		v = value
		_ = kv.ops.Set(key, value, 0)
	}
	return v, e
}

//Increment Increment delta step value(-x,x) for the key
func (kv KV) Increment(key string, delta int64) (int64, error) {
	return kv.ops.IncrBy(key, delta)
}
