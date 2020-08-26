package Test

import (
	"github.com/yoyofx/yoyogo/Internal/EMail"
	"strings"
	"testing"
)

func TestTos(t *testing.T) {
	errMessage := "tos if invalid"

	errValues := []string{
		"",
		"qwerty",
		"qwe;rty",
		"qwe;rty;com",
		// "qwe@rty@com",
		// "@rty",
		// "qwe@",
	}

	for _, errValue := range errValues {
		smtpConnection := EMail.New("smtp.exmail.qq.com", "smtpUser@smtp.exmail.qq.com", "smtpPassword")
		res := smtpConnection.SendMail("from@domain.com", errValue, "This is subject", "Hi! <br><br> This is body")
		if !strings.Contains(res.Error(), errMessage) {
			t.Errorf("Test failed on Tos: %s", errValue)
		}
	}

}
