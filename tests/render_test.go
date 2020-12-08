package tests

import (
	"bytes"
	"encoding/xml"
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/ugorji/go/codec"
	"github.com/yoyofx/yoyogo/tests/testdata/protoexample"
	"github.com/yoyofx/yoyogo/web/actionresult"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TODO unit tests
// test errors

func TestRenderMsgPack(t *testing.T) {
	w := httptest.NewRecorder()
	data := map[string]interface{}{
		"foo": "bar",
	}

	(actionresult.MsgPack{data}).WriteContentType(w)
	assert.Equal(t, "application/msgpack; charset=utf-8", w.Header().Get("Content-Type"))

	err := (actionresult.MsgPack{data}).Render(w)

	assert.NoError(t, err)

	h := new(codec.MsgpackHandle)
	assert.NotNil(t, h)
	buf := bytes.NewBuffer([]byte{})
	assert.NotNil(t, buf)
	err = codec.NewEncoder(buf, h).Encode(data)

	assert.NoError(t, err)
	assert.Equal(t, w.Body.String(), buf.String())
	assert.Equal(t, "application/msgpack; charset=utf-8", w.Header().Get("Content-Type"))
}

func TestRenderJSON(t *testing.T) {
	w := httptest.NewRecorder()
	data := map[string]interface{}{
		"foo":  "bar",
		"html": "<b>",
	}

	(actionresult.Json{data}).WriteContentType(w)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	err := (actionresult.Json{data}).Render(w)

	assert.NoError(t, err)
	assert.Equal(t, "{\"foo\":\"bar\",\"html\":\"\\u003cb\\u003e\"}\n", w.Body.String())
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
}

func TestRenderJSONPanics(t *testing.T) {
	w := httptest.NewRecorder()
	data := make(chan int)

	// json: unsupported type: chan int
	assert.Panics(t, func() { assert.NoError(t, (actionresult.Json{data}).Render(w)) })
}

func TestRenderIndentedJSON(t *testing.T) {
	w := httptest.NewRecorder()
	data := map[string]interface{}{
		"foo": "bar",
		"bar": "foo",
	}

	err := (actionresult.IndentedJson{data}).Render(w)

	assert.NoError(t, err)
	assert.Equal(t, "{\n    \"bar\": \"foo\",\n    \"foo\": \"bar\"\n}", w.Body.String())
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
}

func TestRenderIndentedJSONPanics(t *testing.T) {
	w := httptest.NewRecorder()
	data := make(chan int)

	// json: unsupported type: chan int
	err := (actionresult.IndentedJson{data}).Render(w)
	assert.Error(t, err)
}

func TestRenderSecureJSON(t *testing.T) {
	w1 := httptest.NewRecorder()
	data := map[string]interface{}{
		"foo": "bar",
	}

	(actionresult.SecureJson{"while(1);", data}).WriteContentType(w1)
	assert.Equal(t, "application/json; charset=utf-8", w1.Header().Get("Content-Type"))

	err1 := (actionresult.SecureJson{"while(1);", data}).Render(w1)

	assert.NoError(t, err1)
	assert.Equal(t, "{\"foo\":\"bar\"}", w1.Body.String())
	assert.Equal(t, "application/json; charset=utf-8", w1.Header().Get("Content-Type"))

	w2 := httptest.NewRecorder()
	datas := []map[string]interface{}{{
		"foo": "bar",
	}, {
		"bar": "foo",
	}}

	err2 := (actionresult.SecureJson{"while(1);", datas}).Render(w2)
	assert.NoError(t, err2)
	assert.Equal(t, "while(1);[{\"foo\":\"bar\"},{\"bar\":\"foo\"}]", w2.Body.String())
	assert.Equal(t, "application/json; charset=utf-8", w2.Header().Get("Content-Type"))
}

func TestRenderSecureJSONFail(t *testing.T) {
	w := httptest.NewRecorder()
	data := make(chan int)

	// json: unsupported type: chan int
	err := (actionresult.SecureJson{"while(1);", data}).Render(w)
	assert.Error(t, err)
}

func TestRenderJsonpJSON(t *testing.T) {
	w1 := httptest.NewRecorder()
	data := map[string]interface{}{
		"foo": "bar",
	}

	(actionresult.Jsonp{"x", data}).WriteContentType(w1)
	assert.Equal(t, "application/javascript; charset=utf-8", w1.Header().Get("Content-Type"))

	err1 := (actionresult.Jsonp{"x", data}).Render(w1)

	assert.NoError(t, err1)
	assert.Equal(t, "x({\"foo\":\"bar\"});", w1.Body.String())
	assert.Equal(t, "application/javascript; charset=utf-8", w1.Header().Get("Content-Type"))

	w2 := httptest.NewRecorder()
	datas := []map[string]interface{}{{
		"foo": "bar",
	}, {
		"bar": "foo",
	}}

	err2 := (actionresult.Jsonp{"x", datas}).Render(w2)
	assert.NoError(t, err2)
	assert.Equal(t, "x([{\"foo\":\"bar\"},{\"bar\":\"foo\"}]);", w2.Body.String())
	assert.Equal(t, "application/javascript; charset=utf-8", w2.Header().Get("Content-Type"))
}

func TestRenderJsonpJSONError2(t *testing.T) {
	w := httptest.NewRecorder()
	data := map[string]interface{}{
		"foo": "bar",
	}
	(actionresult.Jsonp{"", data}).WriteContentType(w)
	assert.Equal(t, "application/javascript; charset=utf-8", w.Header().Get("Content-Type"))

	e := (actionresult.Jsonp{"", data}).Render(w)
	assert.NoError(t, e)

	assert.Equal(t, "{\"foo\":\"bar\"}", w.Body.String())
	assert.Equal(t, "application/javascript; charset=utf-8", w.Header().Get("Content-Type"))
}

func TestRenderJsonpJSONFail(t *testing.T) {
	w := httptest.NewRecorder()
	data := make(chan int)

	// json: unsupported type: chan int
	err := (actionresult.Jsonp{"x", data}).Render(w)
	assert.Error(t, err)
}

func TestRenderAsciiJSON(t *testing.T) {
	w1 := httptest.NewRecorder()
	data1 := map[string]interface{}{
		"lang": "GO语言",
		"tag":  "<br>",
	}

	err := (actionresult.AsciiJson{data1}).Render(w1)

	assert.NoError(t, err)
	assert.Equal(t, "{\"lang\":\"GO\\u8bed\\u8a00\",\"tag\":\"\\u003cbr\\u003e\"}", w1.Body.String())
	assert.Equal(t, "application/json", w1.Header().Get("Content-Type"))

	w2 := httptest.NewRecorder()
	data2 := float64(3.1415926)

	err = (actionresult.AsciiJson{data2}).Render(w2)
	assert.NoError(t, err)
	assert.Equal(t, "3.1415926", w2.Body.String())
}

func TestRenderAsciiJSONFail(t *testing.T) {
	w := httptest.NewRecorder()
	data := make(chan int)

	// json: unsupported type: chan int
	assert.Error(t, (actionresult.AsciiJson{data}).Render(w))
}

func TestRenderPureJSON(t *testing.T) {
	w := httptest.NewRecorder()
	data := map[string]interface{}{
		"foo":  "bar",
		"html": "<b>",
	}
	err := (actionresult.PureJson{data}).Render(w)
	assert.NoError(t, err)
	assert.Equal(t, "{\"foo\":\"bar\",\"html\":\"<b>\"}\n", w.Body.String())
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
}

type xmlmap map[string]interface{}

// Allows type H to be used with xml.Marshal
func (h xmlmap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{
		Space: "",
		Local: "map",
	}
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	for key, value := range h {
		elem := xml.StartElement{
			Name: xml.Name{Space: "", Local: key},
			Attr: []xml.Attr{},
		}
		if err := e.EncodeElement(value, elem); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func TestRenderYAML(t *testing.T) {
	w := httptest.NewRecorder()
	data := `
a : Easy!
b:
	c: 2
	d: [3, 4]
	`
	(actionresult.YAML{data}).WriteContentType(w)
	assert.Equal(t, "application/x-yaml; charset=utf-8", w.Header().Get("Content-Type"))

	err := (actionresult.YAML{data}).Render(w)
	assert.NoError(t, err)
	assert.Equal(t, "\"\\na : Easy!\\nb:\\n\\tc: 2\\n\\td: [3, 4]\\n\\t\"\n", w.Body.String())
	assert.Equal(t, "application/x-yaml; charset=utf-8", w.Header().Get("Content-Type"))
}

type fail struct{}

// Hook MarshalYAML
func (ft *fail) MarshalYAML() (interface{}, error) {
	return nil, errors.New("fail")
}

func TestRenderYAMLFail(t *testing.T) {
	w := httptest.NewRecorder()
	err := (actionresult.YAML{&fail{}}).Render(w)
	assert.Error(t, err)
}

// test Protobuf rendering
func TestRenderProtoBuf(t *testing.T) {
	w := httptest.NewRecorder()
	reps := []int64{int64(1), int64(2)}
	label := "test"

	data := &protoexample.Test{
		Label: &label,
		Reps:  reps,
	}

	(actionresult.ProtoBuf{data}).WriteContentType(w)
	protoData, err := proto.Marshal(data)
	assert.NoError(t, err)
	assert.Equal(t, "application/x-protobuf", w.Header().Get("Content-Type"))
	err = (actionresult.ProtoBuf{data}).Render(w)

	assert.NoError(t, err)
	assert.Equal(t, string(protoData), w.Body.String())
	assert.Equal(t, "application/x-protobuf", w.Header().Get("Content-Type"))
}

func TestRenderXML(t *testing.T) {
	w := httptest.NewRecorder()
	data := xmlmap{
		"foo": "bar",
	}

	(actionresult.XML{data}).WriteContentType(w)
	assert.Equal(t, "application/xml; charset=utf-8", w.Header().Get("Content-Type"))

	err := (actionresult.XML{data}).Render(w)

	assert.NoError(t, err)
	assert.Equal(t, "<map><foo>bar</foo></map>", w.Body.String())
	assert.Equal(t, "application/xml; charset=utf-8", w.Header().Get("Content-Type"))
}

func TestRenderRedirect(t *testing.T) {
	req, err := http.NewRequest("GET", "/test-redirect", nil)
	assert.NoError(t, err)

	data1 := actionresult.Redirect{
		Code:     http.StatusMovedPermanently,
		Request:  req,
		Location: "/new/location",
	}

	w := httptest.NewRecorder()
	err = data1.Render(w)
	assert.NoError(t, err)

	data2 := actionresult.Redirect{
		Code:     http.StatusOK,
		Request:  req,
		Location: "/new/location",
	}

	w = httptest.NewRecorder()
	assert.PanicsWithValue(t, "Cannot redirect with status code 200", func() { data2.Render(w) })

	data3 := actionresult.Redirect{
		Code:     http.StatusCreated,
		Request:  req,
		Location: "/new/location",
	}

	w = httptest.NewRecorder()
	err = data3.Render(w)
	assert.NoError(t, err)

	// only improve coverage
	data2.WriteContentType(w)
}

func TestRenderData(t *testing.T) {
	w := httptest.NewRecorder()
	data := []byte("#!PNG some raw data")

	err := (actionresult.Data{
		ContentType: "image/png",
		Data:        data,
	}).Render(w)

	assert.NoError(t, err)
	assert.Equal(t, "#!PNG some raw data", w.Body.String())
	assert.Equal(t, "image/png", w.Header().Get("Content-Type"))
}

func TestRenderString(t *testing.T) {
	w := httptest.NewRecorder()

	(actionresult.Text{
		Format: "hello %s %d",
		Data:   []interface{}{},
	}).WriteContentType(w)
	assert.Equal(t, "text/plain; charset=utf-8", w.Header().Get("Content-Type"))

	err := (actionresult.Text{
		Format: "hola %s %d",
		Data:   []interface{}{"manu", 2},
	}).Render(w)

	assert.NoError(t, err)
	assert.Equal(t, "hola manu 2", w.Body.String())
	assert.Equal(t, "text/plain; charset=utf-8", w.Header().Get("Content-Type"))
}

func TestRenderStringLenZero(t *testing.T) {
	w := httptest.NewRecorder()

	err := (actionresult.Text{
		Format: "hola %s %d",
		Data:   []interface{}{},
	}).Render(w)

	assert.NoError(t, err)
	assert.Equal(t, "hola %s %d", w.Body.String())
	assert.Equal(t, "text/plain; charset=utf-8", w.Header().Get("Content-Type"))
}

func TestRenderHTMLTemplate(t *testing.T) {
	w := httptest.NewRecorder()
	templ := template.Must(template.New("t").Parse(`Hello {{.name}}`))

	htmlRender := actionresult.HTMLProduction{Template: templ}
	instance := htmlRender.Instance("t", map[string]interface{}{
		"name": "alexandernyquist",
	})

	err := instance.Render(w)

	assert.NoError(t, err)
	assert.Equal(t, "Hello alexandernyquist", w.Body.String())
	assert.Equal(t, "text/html; charset=utf-8", w.Header().Get("Content-Type"))
}

func TestRenderHTMLTemplateEmptyName(t *testing.T) {
	w := httptest.NewRecorder()
	templ := template.Must(template.New("").Parse(`Hello {{.name}}`))

	htmlRender := actionresult.HTMLProduction{Template: templ}
	instance := htmlRender.Instance("", map[string]interface{}{
		"name": "alexandernyquist",
	})

	err := instance.Render(w)

	assert.NoError(t, err)
	assert.Equal(t, "Hello alexandernyquist", w.Body.String())
	assert.Equal(t, "text/html; charset=utf-8", w.Header().Get("Content-Type"))
}

func TestRenderHTMLDebugFiles(t *testing.T) {
	w := httptest.NewRecorder()
	htmlRender := actionresult.HTMLDebug{Files: []string{"testdata/template/hello.tmpl"},
		Glob:    "",
		Delims:  actionresult.Delims{Left: "{[{", Right: "}]}"},
		FuncMap: nil,
	}
	instance := htmlRender.Instance("hello.tmpl", map[string]interface{}{
		"name": "thinkerou",
	})

	err := instance.Render(w)

	assert.NoError(t, err)
	assert.Equal(t, "<h1>Hello thinkerou</h1>", w.Body.String())
	assert.Equal(t, "text/html; charset=utf-8", w.Header().Get("Content-Type"))
}

func TestRenderHTMLDebugGlob(t *testing.T) {
	w := httptest.NewRecorder()
	htmlRender := actionresult.HTMLDebug{Files: nil,
		Glob:    "testdata/template/hello*",
		Delims:  actionresult.Delims{Left: "{[{", Right: "}]}"},
		FuncMap: nil,
	}
	instance := htmlRender.Instance("hello.tmpl", map[string]interface{}{
		"name": "thinkerou",
	})

	err := instance.Render(w)

	assert.NoError(t, err)
	assert.Equal(t, "<h1>Hello thinkerou</h1>", w.Body.String())
	assert.Equal(t, "text/html; charset=utf-8", w.Header().Get("Content-Type"))
}

func TestRenderHTMLDebugPanics(t *testing.T) {
	htmlRender := actionresult.HTMLDebug{Files: nil,
		Glob:    "",
		Delims:  actionresult.Delims{"{{", "}}"},
		FuncMap: nil,
	}
	assert.Panics(t, func() { htmlRender.Instance("", nil) })
}
