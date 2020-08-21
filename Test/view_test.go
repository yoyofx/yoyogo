package Test

import (
	"github.com/stretchr/testify/assert"
	"github.com/yoyofx/yoyogo/WebFramework/View"
	"os"
	"path/filepath"
	"testing"
)

func Test_ViewTextRender(t *testing.T) {

	res := View.TextRender("./testdata/view_test.tpl", map[string]interface{}{
		"word": "World!",
	})
	assert.Equal(t, res, "Hello World!")
}

func Test_ViewHtmlRender(t *testing.T) {
	res := View.HtmlRender("./testdata/html_test.tpl", map[string]interface{}{
		"Title":    "hello world",
		"Content":  "<div>Test html content</div>",
		"Content2": "<a href='#'>Test html link</a>",
	})

	willRes := `<html><head><title>hello world</title></head><body>&lt;div&gt;Test html content&lt;/div&gt;<a href='#'>Test html link</a><body><html>`

	assert.Equal(t, res, willRes)

}

func Test_ViewTemplate(t *testing.T) {
	View.SetTemplatePath("./testdata")
	path1 := View.ParseViewName("view_test")

	workDir, _ := os.Getwd()
	tplPath := filepath.Join(workDir, "testdata/view_test.tpl")
	assert.Equal(t, path1, tplPath)
}
