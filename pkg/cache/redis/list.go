package redis

//the List elements type is string
type List struct {
	ops        Ops
	serializer ISerializer
}

//Index get index elements from key List
func (ls List) Index(key string, index int64) (string, error) {
	return ls.ops.LIndex(key, index)
}

//LeftPop left pop the element from key List
func (ls List) LeftPop(key string) (string, error) {
	return ls.ops.LPop(key)
}

//LeftPush left push element to key List
func (ls List) LeftPush(key string, value ...interface{}) (int64, error) {
	return ls.ops.LPush(key, value...)
}

//Range get range elements(strings.) with key List by start and end index
func (ls List) Range(key string, start int64, end int64) ([]string, error) {
	return ls.ops.LRange(key, start, end)
}

//Trim trim range(start-end index) with key List ,and then remove the others.
func (ls List) Trim(key string, start int64, end int64) error {
	return ls.ops.LTrim(key, start, end)
}

//RightPop right pop the element from key List ,that remove it.
func (ls List) RightPop(key string) (string, error) {
	return ls.ops.RPop(key)
}

//RightPush right push element to key List
func (ls List) RightPush(key string, values ...interface{}) (int64, error) {
	return ls.ops.RPush(key, values)
}

//Set set element to key List by index
func (ls List) Set(key string, index int64, value interface{}) error {
	return ls.ops.LSet(key, index, value)
}

//Remove remove count number elements from key List, if
func (ls List) Remove(key string, count int64, value interface{}) (int64, error) {
	return ls.ops.LRemove(key, count, value)
}

//Size the key List of size
func (ls List) Size(key string) (int64, error) {
	return ls.ops.LSize(key)
}

//Clear clear the key List
func (ls List) Clear(key string) (bool, error) {
	c, e := ls.ops.DeleteKey(key)
	return c > 0, e
}

//AddElements add serialization elements to key list
func (ls List) AddElements(key string, values ...interface{}) error {
	var rets []interface{}
	for _, value := range values {
		bytes, _ := ls.serializer.Serialization(value)
		rets = append(rets, bytes)
	}
	_, err := ls.RightPush(key, rets)
	return err
}

func (ls List) GetElement(key string, index int64, elem interface{}) error {
	strElem, _ := ls.Index(key, index)
	return ls.serializer.Deserialization([]byte(strElem), elem)
}

func (ls List) GetElements(key string, startIndex int64, endIndex int64, elements []interface{}) error {
	//strElemArray, _ := ls.Range(key, startIndex, endIndex)
	//for _, strElem := range strElemArray {
	//	//_ = ls.serializer.Deserialization([]byte(strElem), elem)
	//}
	return nil
}
