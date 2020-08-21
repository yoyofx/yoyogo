package View

import (
	"sync"
)

type DefaultEngine struct {
	TemplatePath    string
	shareViewData   map[string]interface{}
	includeTmplName []string
	mutx            sync.Mutex
}

func NewDefaultViewEngine() *DefaultEngine {
	return &DefaultEngine{}
}

func (c *DefaultEngine) SetTemplatePath(option *Option) {
	c.TemplatePath = option.Path
	c.includeTmplName = option.Includes
	SetTemplatePath(option.Path)
}

// 全局通用模板变量
func (c *DefaultEngine) ShareViewData(key string, viewData interface{}) {
	c.mutx.Lock()
	defer c.mutx.Unlock()
	c.shareViewData[key] = viewData
}

// 添加引入模版
func (c *DefaultEngine) AddIncludeTmpl(viewName string) {
	c.mutx.Lock()
	defer c.mutx.Unlock()
	c.includeTmplName = append(c.includeTmplName, viewName)
}

// ViewHtml 快速渲染html视图
func (c *DefaultEngine) ViewHtml(viewName string, viewDataset ...interface{}) (string, error) {
	tmpl := New(viewName, c.includeTmplName...)
	for _, item := range viewDataset {
		tmpl = tmpl.ViewData(item)
	}
	htmlRes, err := tmpl.Render()
	if err != nil {
		return htmlRes, err
	}

	return htmlRes, nil
}
