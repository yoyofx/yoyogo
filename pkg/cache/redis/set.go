package redis

type Set struct {
	ops Ops
}

func (set Set) Add(key string, members ...interface{}) (int64, error) {
	return set.ops.SAdd(key, members...)
}
func (set Set) Difference(keys ...string) ([]string, error) {
	return set.ops.SDiff(keys...)
}
func (set Set) Size(key string) (int64, error) {
	return set.ops.SCard(key)
}
func (set Set) Intersect(keys ...string) ([]string, error) {
	return set.ops.SInter(keys...)
}
func (set Set) IntersectAndStore(destination string, keys ...string) (int64, error) {
	return set.ops.SInterStore(destination, keys...)
}

func (set Set) IsMember(key string, member interface{}) (bool, error) {
	return set.ops.SIsMember(key, member)
}

func (set Set) Members(key string) ([]string, error) {
	return set.ops.SMembers(key)
}

func (set Set) Move(source string, destination string, member interface{}) (bool, error) {
	return set.ops.SMove(source, destination, member)
}

func (set Set) Pop(key string) (string, error) {
	return set.ops.SPop(key)
}

func (set Set) RandMembers(key string, count int64) ([]string, error) {
	return set.ops.SRandMembers(key, count)
}

func (set Set) Remove(key string, members ...interface{}) (int64, error) {
	return set.ops.SRem(key, members...)
}

func (set Set) Union(keys ...string) ([]string, error) {
	return set.ops.SUnion(keys...)
}

func (set Set) UnionAndStore(destination string, keys ...string) (int64, error) {
	return set.ops.SUnionStore(destination, keys...)
}
