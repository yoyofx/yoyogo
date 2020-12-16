package actionresult

import (
	"encoding/json"
	"regexp"
	"strings"
	"unicode"
)

type JsonCamelCase struct {
	Value interface{}
}

// 首字母小写
func LowerFirstCode(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

func (c JsonCamelCase) MarshalJSON() ([]byte, error) {
	var keyMatchRegex = regexp.MustCompile(`\"(\w+)\":`)
	marshalled, err := json.Marshal(c.Value)
	converted := keyMatchRegex.ReplaceAllFunc(
		marshalled,
		func(match []byte) []byte {
			matchStr := string(match)
			key := matchStr[1 : len(matchStr)-2]
			resKey := LowerFirstCode(CaseToCamel(key))
			return []byte(`"` + resKey + `":`)
		},
	)
	return converted, err
}

// 下划线写法转为驼峰写法
func CaseToCamel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}