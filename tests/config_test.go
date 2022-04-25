package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/yoyofx/yoyogo/abstractions"
	"os"
	"testing"
)

type Profile struct {
	DNS string `config:"dns"`
	IP  string `config:"ip"`
	NS  string `config:"namespace"`
}

func TestEnv(t *testing.T) {
	_ = os.Setenv("MYNAMESPACE", "space.yoyogo.run")
	_ = os.Setenv("CUSTOM_ENV", "my env variable")
	_ = os.Setenv("REMOTE_HOST", "my host")

	config := abstractions.NewConfigurationBuilder().
		AddEnvironment().
		AddYamlFile("config").Build()

	var profile Profile
	config.GetConfigObject("profile", &profile)

	assert.Equal(t, profile.NS, "space.yoyogo.run")
	assert.Equal(t, profile.DNS, "my host")
	assert.Equal(t, profile.IP, "10.0.1.12")

	env := config.Get("env")
	dns := config.Get("profile.dns")
	ip := config.Get("profile.ip")

	assert.Equal(t, env, "my env variable")
	assert.Equal(t, dns, "my host")
	assert.Equal(t, ip, "10.0.1.12")

}
