package Test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yoyofx/yoyogo/Utils/jwt"
	"testing"
)

func TestCreateToken(t *testing.T) {
	SecretKey := []byte("AllYourBase")
	token, _ := jwt.CreateToken(SecretKey, "YDQ", 2222)
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
