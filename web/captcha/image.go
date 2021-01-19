package captcha

import (
	"github.com/yoyofx/yoyogo/pkg/captcha"
	"github.com/yoyofx/yoyogo/utils"
	"strings"
)

// CreateImage
func CreateImage(n int, size ...int) (text string, md5 string, imgByte []byte) {
	text = utils.GetRandStr(n)
	textMd5 := utils.Md5String(strings.ToUpper(text))
	width := 180
	height := 60
	if len(size) >= 2 {
		width = size[0]
		height = size[1]
	}
	imgByte = captcha.ImageText(width, height, text)

	return text, textMd5, imgByte
}

func Validation(text string, md5 string) bool {
	textMd5 := utils.Md5String(strings.ToUpper(text))
	return textMd5 == md5
}
