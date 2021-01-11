package redis

import "time"

type Client struct {
	//------- data struct -----------
	kv   KV
	list List
	hash Hash
	set  Set
	zset ZSet
	geo  Geo
	lock Lock
	//-----------------------
	ops             Ops
	valueSerializer ISerializer
}

//NewClient new redis client by Options
func NewClient(options *Options) IClient {
	var ops Ops
	if options.Addrs == nil {
		ops = NewStandaloneOps(options)
	} else {
		ops = NewClusterOps(options)
	}
	kv := KV{ops: ops}
	list := List{ops: ops}
	hash := Hash{ops: ops}
	set := Set{ops: ops}
	zset := ZSet{ops: ops}
	geo := Geo{ops: ops}
	lock := Lock{ops: ops}

	return &Client{ops: ops, kv: kv, list: list, hash: hash, set: set, zset: zset, geo: geo, lock: lock}
}

//SetSerializer set value serializer
func (c *Client) SetSerializer(serializer ISerializer) {
	c.valueSerializer = serializer
}

//GetKVOps Returns the operations performed on simple values (or Strings in Redis terminology).
func (c *Client) GetKVOps() KV {
	c.kv.serializer = c.valueSerializer
	return c.kv
}

//GetListOps Returns the operations performed on list values.
func (c *Client) GetListOps() List {
	c.list.serializer = c.valueSerializer
	return c.list
}

//GetHashOps Returns the operations performed on hash values.
func (c *Client) GetHashOps() Hash {
	c.hash.serializer = c.valueSerializer
	return c.hash
}

//GetSetOps Returns the operations performed on set values.
func (c *Client) GetSetOps() Set {
	return c.set
}

//GetZSetOps Returns the operations performed on zset values (also known as sorted sets).
func (c *Client) GetZSetOps() ZSet {
	return c.zset
}

//GetGeoOps Geo Returns the operations performed on geo values (also known GIS system).
func (c *Client) GetGeoOps() Geo {
	return c.geo
}

//GetLockOps Returns the operations performed on locker values.
func (c *Client) GetLockOps() Lock {
	return c.lock
}

// Ping return PONG and error
func (c *Client) Ping() (string, error) {
	return c.ops.Ping()
}

// SetExpire cmd by expire
func (c *Client) SetExpire(key string, expiration time.Duration) (bool, error) {
	return c.ops.SetExpire(key, expiration)
}

// SetExpire cmd by TTL
func (c *Client) GetExpire(key string) (time.Duration, error) {
	return c.ops.TTL(key)
}

// Delete delete the key, cmd by del
func (c *Client) Delete(key string) bool {
	n, _ := c.ops.DeleteKey(key)
	return n > 0
}

// HasKey exists the key, cmd by exists
func (c *Client) HasKey(key string) bool {
	h, _ := c.ops.Exists(key)
	return h
}

// RandomKey return random Key for db
func (c *Client) RandomKey() (string, error) {
	return c.ops.RandomKey()
}

// RandomKey return random Key for db
func (c *Client) Close() error {
	return c.ops.Close()
}
