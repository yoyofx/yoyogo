package Test

import (
	"fmt"
	"reflect"
	"testing"
)

type UserInfo struct {
	Name string `json:"name" w1:"12"`
	Age  int
}

func Test_StructGetFieldTag(t *testing.T) {
	user := &UserInfo{"John Doe The Fourth", 20}

	value := reflect.TypeOf(user).Elem()
	for i := 0; i < value.NumField(); i++ {
		f := value.Field(i)
		fmt.Printf("%d: %s %s %s \n", i,
			f.Name, f.Type, f.Tag.Get("json"))
	}
}
