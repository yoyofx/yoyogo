package tests

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"os"
	"testing"
)

type Profile struct {
	DNS string
	IP  string
}

func TestEnv(t *testing.T) {

	_ = os.Setenv("CUSTOM_ENV", "my env variable")
	_ = os.Setenv("REMOTE_HOST", "my host")

	config := abstractions.NewConfigurationBuilder().
		AddEnvironment().
		AddYamlFile("config").Build()

	//str := config.Get("env")
	//str2 := config.Get("profile.dns")
	//str3 := config.Get("profile.ip")
	//fmt.Println(str)
	//fmt.Println(str2)
	//assert.Equal(t, str3, "yoyogoDefault")
	var profile Profile
	config.GetConfigObject("profile", &profile)

}
