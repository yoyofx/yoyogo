package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/yoyofx/yoyogo/web/view"
	"os"
	"path/filepath"
	"testing"
)

func Test_ViewTextRender(t *testing.T) {

	res := view.TextRender("./testdata/view_test.tpl", map[string]interface{}{
		"word": "World!",
	})
	assert.Equal(t, res, "Hello World!")
}

func Test_ViewHtmlRender(t *testing.T) {
	res := view.HtmlRender("./testdata/html_test.tpl", map[string]interface{}{
		"Title":    "hello world",
		"Content":  "<div>tests html content</div>",
		"Content2": "<a href='#'>tests html link</a>",
	})

	willRes := `<html><head><title>hello world</title></head><body>&lt;div&gt;tests html content&lt;/div&gt;<a href='#'>tests html link</a><body><html>`

	assert.Equal(t, res, willRes)

}

func Test_ViewTemplate(t *testing.T) {
	view.SetTemplatePath("./testdata")
	path1 := view.ParseViewName("view_test")

	workDir, _ := os.Getwd()
	tplPath := filepath.Join(workDir, "testdata/view_test.tpl")
	assert.Equal(t, path1, tplPath)
}
