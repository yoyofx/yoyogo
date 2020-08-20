package Test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yoyofx/yoyogo/Utils/jwt"
	"testing"
	"time"
)

func TestCreateToken(t *testing.T) {
	SecretKey := []byte("12391JdeOW^%$#@")
	token, _ := jwt.CreateToken(SecretKey, "YDQ", 2222, int64(time.Now().Add(time.Hour*72).Unix()))
	fmt.Println(token)

	claims, err := jwt.ParseToken(token, SecretKey)
	if nil != err {
		fmt.Println(" err :", err)
	}
	fmt.Println("claims:", claims)
	fmt.Println("claims uid:", claims.(jwt.MapClaims)["uid"])

	assert.Equal(t, err, nil)
	assert.Equal(t, int(claims.(jwt.MapClaims)["uid"].(float64)), 2222)
	assert.Equal(t, claims.(jwt.MapClaims)["iss"], "YDQ")
}
