package httpclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"path/filepath"
	"strings"
	"sync"
)

type Request struct {
	url               string
	method            string
	params            map[string]interface{}
	requestBody       []byte
	files             []UploadFile
	host              string
	header            http.Header
	contentType       string
	errorRaw          string
	timeout           int
	mutex             sync.Mutex
	cooJar            *cookiejar.Jar
	skipHttps         bool
	disableKeepAlives bool
	cookieData        map[string]*http.Cookie
}

// Header 设置header头信息
func (c *Request) Header(key, value string) *Request {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.header[key] = []string{value}
	if strings.ToLower(key) == "content-type" {
		c.contentType = value
	}
	return c
}

func (c *Request) GetUrl() string {
	return c.url
}

func (c *Request) POST(url string) *Request {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.method = "POST"
	c.url = url
	return c
}

func (c *Request) GET(url string) *Request {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.method = "GET"
	c.url = url
	return c
}

func (c *Request) ContentType(value string) *Request {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.header.Set("Content-Type", value)
	c.contentType = value
	return c
}

// WithHost 添加请求头host
func (c *Request) WithHost(host string) *Request {
	c.host = host
	return c
}

// WithFile 添加文件form-data方式发送
func (c *Request) AddFile(name string, filePath string) *Request {
	c.files = append(c.files, UploadFile{name, filePath})
	return c
}

// SkipHttps 跳过https证书校验
func (c *Request) SkipHttps() *Request {
	c.skipHttps = true
	return c
}

// DisableKeepAlives 关闭KeepAlives
func (c *Request) DisableKeepAlives() *Request {
	c.disableKeepAlives = true
	return c
}

// ClearParam 清空重置请求参数
func (c *Request) ClearParam() *Request {
	c.requestBody = []byte{}
	c.params = map[string]interface{}{}
	return c
}

func (c *Request) SetTimeout(timeout int) *Request {
	c.timeout = timeout
	return c
}

// cookie保持 通过配置请求id将cookie保持
func (c *Request) WithCookie() *Request {
	if c.cooJar == nil {
		cooJar, _ := cookiejar.New(nil)
		c.cooJar = cooJar
	}
	return c
}

func (c *Request) Error() error {
	if c.errorRaw == "" {
		return nil
	}
	return errors.New(c.errorRaw)
}

func (c *Request) setCookieData(name string, cookie *http.Cookie) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.cookieData[name] = cookie

}

// 直接传递body中的参数
func (c *Request) WithBody(bodyStream string) *Request {
	c.requestBody = []byte(bodyStream)
	return c
}

// 参数设置 表单请求支持json和form两种类型
func (c *Request) FormParams(obj map[string]interface{}) *Request {
	c.params = obj
	return c
}

func (c *Request) paraseParams() *Request {
	if len(c.files) > 0 {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		for _, uploadFile := range c.files {
			fileBytes, err := ioutil.ReadFile(uploadFile.Filepath)
			if err != nil {
				fmt.Printf("can not open file %s", uploadFile.Filepath)
				continue
			}
			part, err := writer.CreateFormFile(uploadFile.Name, filepath.Base(uploadFile.Filepath))
			if err != nil {
				fmt.Printf("CreateFormFile error %v", err)
				continue
			}
			part.Write(fileBytes)
		}

		for k, v := range c.params {
			stringVal := fmt.Sprintf("%v", v)
			if err := writer.WriteField(k, stringVal); err != nil {
				fmt.Printf("WriteField error %v", err)
				continue
			}
		}

		c.contentType = writer.FormDataContentType()
		if err := writer.Close(); err != nil {
			panic(err)
		}

		c.requestBody = body.Bytes()
		return c
	}

	if strings.Contains(c.contentType, "application/x-www-form-urlencoded") {
		params := ""
		val := make([]interface{}, 0, len(c.params))
		for k, item := range c.params {
			params += k + "=%v&"
			val = append(val, item)
		}
		paramsLen := len(params) - 1
		params = params[0:paramsLen]
		params = fmt.Sprintf(params, val...)
		c.requestBody = []byte(params)

		return c
	}

	if strings.Contains(c.contentType, "application/json") {
		params, err := json.Marshal(c.params)
		if err != nil {
			c.errorRaw += err.Error() + "|"
		}
		c.requestBody = params
		return c
	}
	c.errorRaw += "Use params you  must set  Content-Type eq 'application/json' or 'application/x-www-form-urlencoded'"

	return c
}
