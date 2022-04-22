package tests

import (
	"fmt"
	"github.com/yoyofx/yoyogo/abstractions"
	"testing"
)

func TestEnv(t *testing.T) {
	config := abstractions.NewConfigurationBuilder().
		AddEnvironment().
		AddYamlFile("config").Build()
	str := config.Get("env")
	str2 := config.Get("profile.homebrew")
	str3 := config.Get("profile.default")
	fmt.Println(str)
	fmt.Println(str2)
	fmt.Println(str3)
}
