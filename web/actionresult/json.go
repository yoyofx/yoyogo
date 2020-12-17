package actionresult

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/yoyofx/yoyogo/web/actionresult/extension"
	"html/template"
	"net/http"
)

type Json struct {
	Data interface{}
}

type IndentedJson struct {
	Data interface{}
}

type SecureJson struct {
	Prefix string
	Data   interface{}
}

type Jsonp struct {
	Callback string
	Data     interface{}
}

// AsciiJSON contains the given interface object.
type AsciiJson struct {
	Data interface{}
}

// PureJSON contains the given interface object.
type PureJson struct {
	Data interface{}
}

var jsonContentType = []string{"application/json; charset=utf-8"}
var jsonpContentType = []string{"application/javascript; charset=utf-8"}
var jsonAsciiContentType = []string{"application/json"}

func writeJSON(w http.ResponseWriter, obj interface{}) error {
	writeContentType(w, jsonContentType)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(&obj)
	return err
}

func writeJsonCamel(w http.ResponseWriter, obj interface{}) error {
	writeContentType(w, jsonContentType)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(&extension.JsonCamelCase{Value: obj})
	return err
}

var (
	jsonEncoder extension.Encoder
)

func SetJsonSerializeEncoder(encoder extension.Encoder) {
	jsonEncoder = encoder
}

// actionresult (JSON) writes data with custom ContentType.
func (d Json) Render(w http.ResponseWriter) (err error) {
	writeContentType(w, jsonContentType)
	if err = jsonEncoder.Encode(w, d.Data); err != nil {
		panic(err)
	}
	return
}

func (d Json) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}

// actionresult (IndentedJSON) marshals the given interface object and writes it with custom ContentType.
func (r IndentedJson) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)
	jsonBytes, err := json.MarshalIndent(r.Data, "", "    ")
	if err != nil {
		return err
	}
	_, err = w.Write(jsonBytes)
	return err
}

func (r IndentedJson) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}

// SecureJSON 用来防止json劫持。
// 如果给定的结构体是数值型，默认预置“while(1)” 到response body
// actionresult (SecureJSON) marshals the given interface object and writes it with custom ContentType.
func (r SecureJson) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)
	jsonBytes, err := json.Marshal(r.Data)
	if err != nil {
		return err
	}
	// if the jsonBytes is array values
	if bytes.HasPrefix(jsonBytes, []byte("[")) && bytes.HasSuffix(jsonBytes, []byte("]")) {
		_, err = w.Write([]byte(r.Prefix))
		if err != nil {
			return err
		}
	}
	_, err = w.Write(jsonBytes)
	return err
}

func (r SecureJson) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}

// callback(extension)
// actionresult (Jsonp JSON) marshals the given interface object and writes it and its callback with custom ContentType.
func (r Jsonp) Render(w http.ResponseWriter) (err error) {
	r.WriteContentType(w)
	ret, err := json.Marshal(r.Data)
	if err != nil {
		return err
	}

	if r.Callback == "" {
		_, err = w.Write(ret)
		return err
	}

	callback := template.JSEscapeString(r.Callback)
	_, err = w.Write([]byte(callback))
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("("))
	if err != nil {
		return err
	}
	_, err = w.Write(ret)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(");"))
	if err != nil {
		return err
	}

	return nil
}

func (r Jsonp) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonpContentType)
}

// 使用Ascii JSON 生成仅有 ASCII字符的JSON，非ASCII字符将会被转义
// actionresult (AsciiJSON) marshals the given interface object and writes it with custom ContentType.
func (r AsciiJson) Render(w http.ResponseWriter) (err error) {
	r.WriteContentType(w)
	ret, err := json.Marshal(r.Data)
	if err != nil {
		return err
	}

	var buffer bytes.Buffer
	for _, r := range string(ret) {
		cvt := string(r)
		if r >= 128 {
			cvt = fmt.Sprintf("\\u%04x", int64(r))
		}
		buffer.WriteString(cvt)
	}

	_, err = w.Write(buffer.Bytes())
	return err
}

func (r AsciiJson) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonAsciiContentType)
}

// 原始字符串，不通过Html编码
// actionresult (PureJSON) writes custom ContentType and encodes the given interface object.
func (r PureJson) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	return encoder.Encode(r.Data)
}

// WriteContentType (PureJSON) writes custom ContentType.
func (r PureJson) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}
