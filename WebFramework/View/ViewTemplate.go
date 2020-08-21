package View

import (
	"errors"
	"github.com/yoyofx/yoyogo/Utils"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
)

var templatePath string

func SetTemplatePath(path string) {
	templatePath = path
}

type Template struct {
	ViewName    string
	ViewPath    string
	IncludeTmpl []string
	ShowData    map[string]interface{}
	mutex       sync.Mutex
}

func New(viewName string, includeFiles ...string) *Template {
	viewPath := ParseViewName(viewName)
	includeTmpl := make([]string, 0)
	for _, item := range includeFiles {
		includeTmpl = append(includeTmpl, ParseViewName(item))
	}
	appTmpl := &Template{
		ViewName:    viewName,
		ViewPath:    viewPath,
		ShowData:    make(map[string]interface{}),
		IncludeTmpl: includeTmpl,
	}

	return appTmpl
}

// view data is map[string]interface{}
func (t *Template) ViewData(values interface{}) *Template {
	objT := reflect.TypeOf(values)
	objV := reflect.ValueOf(values)

	if objT.Kind() == reflect.Ptr {
		objT = objT.Elem()
		objV = objV.Elem()
	}
	switch objT.Kind() {
	case reflect.Struct:
		t.mutex.Lock()
		defer t.mutex.Unlock()
		for i := 0; i < objT.NumField(); i++ {
			objName := objT.Field(i).Name
			objValue := objV.Field(i).Interface()
			t.ShowData[objName] = objValue
		}
	case reflect.Map:
		t.mutex.Lock()
		defer t.mutex.Unlock()
		item := objV.MapRange()
		for item.Next() {
			k := item.Key()
			v := item.Value()
			if k.Kind() != reflect.String {
				typePanic()
			}
			t.ShowData[k.String()] = v.Interface()
		}
	default:
		typePanic()
	}

	return t
}

// view data is name and value , that value is interface{}
func (t *Template) ViewDataKV(name string, value interface{}) *Template {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.ShowData[name] = value
	return t
}

func (t *Template) Render() (viewHtmlRes string, err error) {
	isFile, _ := Utils.FileExists(t.ViewPath)
	if isFile == false {
		errMsg := t.ViewPath + " not found"
		return viewHtmlRes, errors.New(errMsg)
	}

	viewHtmlRes = HtmlRender(t.ViewPath, t.ShowData, t.IncludeTmpl...)
	return viewHtmlRes, nil
}

func (t *Template) RenderText() (viewHtmlRes string, err error) {
	isFile, _ := Utils.FileExists(t.ViewPath)
	if isFile == false {
		errMsg := t.ViewPath + " not found"
		return viewHtmlRes, errors.New(errMsg)
	}

	viewHtmlRes = TextRender(t.ViewPath, t.ShowData, t.IncludeTmpl...)
	return viewHtmlRes, nil
}

func ParseViewName(viewName string) string {
	if templatePath == "" {
		panic("view path not found")
	}
	workDir, _ := os.Getwd()
	templatePathRe := filepath.Join(workDir, templatePath)
	viewName = strings.Replace(viewName, ".", string(filepath.Separator), -1)
	viewPath := templatePathRe + string(filepath.Separator) + viewName + ".tpl"
	return viewPath
}

func typePanic() {
	panic("viewData must be map[string]interface{} or struct ")
}
