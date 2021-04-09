package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
	"github.com/yoyofx/yoyogo/pkg/httpclient"
	"github.com/yoyofx/yoyogo/pkg/servicediscovery/memory"
	"github.com/yoyofx/yoyogo/pkg/servicediscovery/strategy"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
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
	//test server
	httpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte("ok"))
	}))
	//httpServer.URL = "http://127.0.0.1:8080"
	defer httpServer.Close()

	//test client
	url := httpServer.URL
	uri := strings.Split(url, ":")
	port, _ := strconv.ParseUint(uri[2], 10, 64)
	url = strings.Replace(url, "127.0.0.1", "[operations]", -1)
	factory := httpclient.ClientFactory{
		Selector: servicediscovery.Selector{DiscoveryCache: &memory.MemoryCache{Services: []string{"localhost"}, Port: port},
			Strategy: strategy.NewRound()}}
	req, err := factory.BuilderSelectorRequest(url, "GET")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, req.GetUrl(), fmt.Sprintf("http://localhost:%v", port))
	httpclient.AddHttpClient("operations", func(options *httpclient.ClientOptions) {
		options.AddClientBaseUrl("http://localhost")
	})
	client, err := factory.CreateClient("operations")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, req.GetUrl(), fmt.Sprintf("http://localhost:%v", port))
	res, err := client.Do(req)
	assert.Equal(t, string(res.Body), "ok")
}

func TestHttpClientFactoryBaseUrl(t *testing.T) {
	httpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte("ok"))
	}))
	fmt.Print(httpServer.URL)
	//httpServer.URL = "http://127.0.0.1:8080"
	defer httpServer.Close()
	httpclient.AddHttpClient("demo", func(options *httpclient.ClientOptions) {
		options.AddClientBaseUrl(httpServer.URL)
	})
	factory := httpclient.ClientFactory{}
	client, err := factory.CreateClient("demo")
	if err != nil {
		panic(err)
	}
	req := &httpclient.Request{}
	req.GET("")
	req.SetTimeout(10)
	res, err := client.Do(req)
	fmt.Print(res)
	assert.Equal(t, string(res.Body), "ok")
}
