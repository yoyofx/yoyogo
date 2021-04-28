package context

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"github.com/yoyofx/yoyogo/utils"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

type Input struct {
	Request        *http.Request
	RouterData     url.Values
	bodyCache      []byte
	queryString    url.Values
	formCache      url.Values
	allParamsCache url.Values
	RequestMaxSize int64
}

// NewInput return OrangeInput.
func NewInput(request *http.Request, maxMemory int64) *Input {
	//var buf []byte
	if request.ContentLength > maxMemory {
		panic(fmt.Sprintf("request body is too large, msxMemory is %d", maxMemory))
	}
	//queryStrings, _ := url.ParseQuery(request.URL.RawQuery)
	input := &Input{
		Request:        request,
		RouterData:     url.Values{},
		RequestMaxSize: maxMemory,
	}
	return input
}

// Reset init the
func (input *Input) Reset(request *http.Request) {
	input.Request = request
	input.RouterData = url.Values{}
	input.bodyCache = []byte{}
}

func (input *Input) GetReader() *http.Request {
	return input.Request
}

// Header returns request header item string by a given string.
// if non-existed, return empty string.
func (input *Input) Header(key string) string {
	return input.Request.Header.Get(key)
}

// RemoteIP RemoteAddr to an "IP" address
func (input *Input) RemoteIP() string {
	host, _, _ := net.SplitHostPort(input.Request.RemoteAddr)
	return host
}

// RealIP returns the first ip from 'X-Forwarded-For' or 'X-Real-IP' header key
// if not exists data, returns request.RemoteAddr
// fixed for #164
func (input *Input) RealIP() string {
	if ip := input.Request.Header.Get(HeaderXForwardedFor); ip != "" {
		return strings.Split(ip, ", ")[0]
	}
	if ip := input.Request.Header.Get(HeaderXRealIP); ip != "" {
		return ip
	}
	host, _, _ := net.SplitHostPort(input.Request.RemoteAddr)
	return host
}

// FullRemoteIP RemoteAddr to an "IP:port" address
func (input *Input) FullRemoteIP() string {
	fullIp := input.Request.RemoteAddr
	return fullIp
}

// IsAJAX returns if it is a ajax request
func (input *Input) IsAJAX() bool {
	return strings.Contains(input.Request.Header.Get(HeaderXRequestedWith), "XMLHttpRequest")
}

func (input *Input) IsWebsocket() bool {
	if strings.Contains(strings.ToLower(input.Request.Header.Get("Connection")), "upgrade") &&
		strings.EqualFold(input.Request.Header.Get("Upgrade"), "websocket") {
		return true
	}
	return false
}

// IsUpload returns boolean of whether file uploads in this request or not..
func (input *Input) IsUpload() bool {
	return strings.Contains(input.Request.Header.Get("Content-Type"), "multipart/form-data")
}

//Get Cookie by Name
func (input *Input) GetCookie(name string) string {
	cookie, err := input.Request.Cookie(name)
	if err != nil {
		return ""
	}
	return url.QueryEscape(cookie.Value)
}

//Get Http Method.
func (input *Input) Method() string {
	return input.Request.Method
}

// Path returns URL Path string.
func (input *Input) Path() string {
	return input.Request.URL.Path
}

// Url get request url
func (input *Input) Url() string {
	return input.Request.URL.String()
}

// Referer returns request referer.
func (input *Input) Referer() string {
	return input.Request.Header.Get("Referer")
}

// UserAgent returns http request UserAgent
func (input *Input) UserAgent() string {
	return input.Request.Header.Get("User-Agent")
}

// Scheme returns request scheme as "http" or "https".
func (input *Input) Scheme() string {
	if scheme := input.Request.Header.Get("X-Forwarded-Proto"); scheme != "" {
		return scheme
	}
	if input.Request.URL.Scheme != "" {
		return input.Request.URL.Scheme
	}
	if input.Request.TLS == nil {
		return "http"
	}
	return "https"
}

// Domain returns host name.
// Alias of Host method.
func (input *Input) Domain() string {
	return input.Host()
}

// Host returns host name.
// if no host info in request, return localhost.
func (input *Input) Host() string {
	if input.Request.Host != "" {
		if hostPart, _, err := net.SplitHostPort(input.Request.Host); err == nil {
			return hostPart
		}
		return input.Request.Host
	}
	return "localhost"
}

// ParseFormOrMultipartForm parseForm or parseMultiForm based on Content-type
func (input *Input) ParseFormOrMultipartForm(maxMemory int64) error {
	if input.formCache == nil {
		input.formCache = make(url.Values)
		// Parse the body depending on the content type.
		if strings.Contains(input.Header("Content-Type"), "multipart/form-data") {
			if err := input.Request.ParseMultipartForm(maxMemory); err != nil {
				return errors.New("Error parsing request body:" + err.Error())
			}
		} else if err := input.Request.ParseForm(); err != nil {
			return errors.New("Error parsing request body:" + err.Error())
		}
		input.formCache = input.Request.PostForm
	}
	return nil
}

// QueryStrings Get queryString
func (input *Input) QueryStrings() url.Values {
	if input.queryString == nil {
		if input.Request != nil {
			input.queryString = input.Request.URL.Query()
		} else {
			input.queryString = url.Values{}
		}
	}
	return input.queryString
}

func (input *Input) GetBody() []byte {
	if input.bodyCache == nil {
		buf, _ := ioutil.ReadAll(input.Request.Body)
		input.Request.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
		input.bodyCache = buf
	}
	return input.bodyCache
}

// Query Get Query String By Key
// GET /path?id=1234&name=Manu&value=
// c.Query("id") == "1234"
// c.Query("name") == "Manu"
// c.Query("value") == ""
// c.Query("wtf") == ""
func (input *Input) Query(key string) string {
	return input.QueryStrings().Get(key)
}

func (input *Input) QueryDefault(key string, defaultval string) string {
	val := input.QueryStrings().Get(key)
	if val == "" {
		val = defaultval
	}
	return val
}

// FormFile gets file from request.
func (input *Input) FormFile(key string) (multipart.File, *multipart.FileHeader, error) {
	return input.Request.FormFile(key)
}

// SaveFile saves the form file and
// returns the filename.
func (input *Input) SaveFile(name, saveDir string) (string, error) {
	fr, handle, err := input.FormFile(name)
	if err != nil {
		return "", err
	}
	defer fr.Close()
	fw, err := os.OpenFile(path.Join(saveDir, handle.Filename), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return "", err
	}
	defer fw.Close()

	_, err = io.Copy(fw, fr)
	return handle.Filename, err
}

func (input *Input) GetAllParam() url.Values {
	if input.allParamsCache == nil {
		input.allParamsCache = make(url.Values)
		err := input.ParseFormOrMultipartForm(input.RequestMaxSize)
		if err == nil {
			utils.MergeMap(input.allParamsCache, input.formCache)
			if input.Request.MultipartForm != nil {
				utils.MergeMap(input.allParamsCache, input.Request.MultipartForm.Value)
			}

		}
		utils.MergeMap(input.allParamsCache, input.QueryStrings())
		utils.MergeMap(input.allParamsCache, input.RouterData)
	}
	return input.allParamsCache
}

//Get Post Param
func (input *Input) Param(name string) string {
	form := input.GetAllParam()
	if form[name] != nil {
		return form[name][0]
	}
	return ""
}

func bodyAllowedForStatus(status int) bool {
	switch {
	case status >= 100 && status <= 199:
		return false
	case status == http.StatusNoContent:
		return false
	case status == http.StatusNotModified:
		return false
	}
	return true
}
