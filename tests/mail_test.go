package tests

import (
	"github.com/yoyofx/yoyogo/pkg/email"
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
		smtpConnection := email.New("smtp.exmail.qq.com", "smtpUser@smtp.exmail.qq.com", "smtpPassword")
		res := smtpConnection.SendMail("from@domain.com", errValue, "This is subject", "Hi! <br><br> This is body")
		if !strings.Contains(res.Error(), errMessage) {
			t.Errorf("tests failed on Tos: %s", errValue)
		}
	}

}
