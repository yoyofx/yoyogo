package redis

type Hash struct {
	ops        Ops
	serializer ISerializer
}

//Delete delete hashset for key and hash key
func (h Hash) Delete(key string, hashKey string) error {
	_, e := h.ops.HDel(key, hashKey)
	return e
}

//GetEntries get all hash keys and values
func (h Hash) GetEntries(key string) (map[string]string, error) {
	return h.ops.HGetAll(key)
}

//GetString get string for key and hash key
func (h Hash) GetString(key string, hashKey string) (string, error) {
	return h.ops.HGet(key, hashKey)
}

//Get get value(struct , int64, float64) for the key and hash key
func (h Hash) Get(key string, hashKey string, value interface{}) error {
	v, e := h.ops.HGet(key, hashKey)
	if e != nil {
		return e
	}
	e = h.serializer.Deserialization([]byte(v), value)
	return e
}

//Exists has key and hash key
func (h Hash) Exists(key string, hashKey string) bool {
	v, _ := h.ops.HExists(key, hashKey)
	return v
}

//MultiGet get all values for the keys
func (h Hash) MultiGet(key string, hashKeys ...string) ([]interface{}, error) {
	return h.ops.HMGet(key, hashKeys...)
}

//GetHashKeys get hask keys for the key
func (h Hash) GetHashKeys(key string) ([]string, error) {
	return h.ops.HKeys(key)
}

//Put put value (struct , int64, float64) to the key and hash key
func (h Hash) Put(key string, hashKey string, value interface{}) error {
	ss, e := h.serializer.Serialization(value)
	if e != nil {
		return e
	}
	_, e = h.ops.HSet(key, hashKey, ss)
	return e
}

//Put put value (struct , int64, float64) to the key and hash key; cmd by SetNx .
func (h Hash) PutIfAbsent(key string, hashKey string, value interface{}) (bool, error) {
	ss, e := h.serializer.Serialization(value)
	if e != nil {
		return false, e
	}
	v, e1 := h.ops.HSetNX(key, hashKey, ss)
	return v, e1
}

//Increment increment value by delta step to the key and hash key
func (h Hash) Increment(key string, hashKey string, delta int64) int64 {
	v, _ := h.ops.HIncrBy(key, hashKey, delta)
	return v
}

//Size get len by key's all hash size
func (h Hash) Size(key string) int64 {
	v, _ := h.ops.HLen(key)
	return v
}
