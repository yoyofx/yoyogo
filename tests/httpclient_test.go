package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
	"github.com/yoyofx/yoyogo/pkg/httpclient"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetHttp(t *testing.T) {
	httpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)

		if r.Header.Get("Content-Type") == "application/json" {
			if string(body) != `{"aid":1002,"auth":"ss"}` {
				t.Error("RunPost json type func error")
			}
		}

		if r.Header.Get("Content-Type") == "application/x-www-form-urlencoded" {
			if string(body) != `word=你好` {
				t.Error("RunPost WithFormRequest type func error 'word=你好' != ", string(body))
			}
		}

		if r.Header.Get("Content-Type") == "application/text" {
			w.WriteHeader(200)
			_, _ = w.Write([]byte("hello"))
		}
	}))
	defer httpServer.Close()

	c := httpclient.NewClient()
	request := c.WithJsonRequest(`{"aid":1002,"auth":"ss"}`).POST(httpServer.URL)
	_, _ = c.Do(request)

	request1 := c.WithFormRequest(map[string]interface{}{
		"word": "你好",
	}).POST(httpServer.URL)
	_, _ = c.Do(request1)

	request2 := c.WithRequest().Header("Content-Type", "application/text").GET(httpServer.URL)
	resp, _ := c.Do(request2)

	assert.Equal(t, resp.GetRequestTime().Seconds() < 5, true)
	assert.Equal(t, string(resp.Body), "hello")
}

func TestUriParser(t *testing.T) {

	url := "http://[DEMO1]/app/v1/getuser?id=1"

	parser := servicediscovery.NewUriParser(url)

	assert.Equal(t, parser.GetUriEntry().Protocol, "http")

	url1 := parser.Generate("127.0.0.1:8080")

	assert.Equal(t, url1, "http://127.0.0.1:8080/app/v1/getuser?id=1")

}

func TestHttpCleintFactory(t *testing.T) {
	//url := "https://mnurtestapi.mengniu.com.cn/operations/v1/0/user-info/role"
	//factory := httpclient.ClientFactory{}
	//client, err := factory.CreatHttpClient(url)
	//if err != nil {
	//	panic(err)
	//}
	//client.Send()
}
