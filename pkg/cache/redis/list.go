package redis

type List struct {
	ops Ops
}

func (ls List) Index(key string, index int64) interface{} {

	return nil
}

func (ls List) LeftPop(key string) interface{} {

	return nil
}

func (ls List) LeftPush(key string, value ...interface{}) interface{} {

	return nil
}

func (ls List) Range(key string, start int64, end int64) []interface{} {

	return nil
}

func (ls List) RightPop(key string) interface{} {

	return nil
}

func (ls List) RightPush(key string, value ...interface{}) interface{} {

	return nil
}

func (ls List) Set(key string, index int64, value interface{}) {

}

func (ls List) Size(key string) (int64, error) {
	return 0, nil
}
