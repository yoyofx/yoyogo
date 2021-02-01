package tests

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/yoyofx/yoyogo/abstractions/health"
	"testing"
)

type downHealth struct{}

func (u downHealth) Health() health.ComponentStatus {
	return health.Down("downHealth").WithDetail("reason", "error:down")
}

type upHealth struct{}

func (u upHealth) Health() health.ComponentStatus {
	return health.Up("UpHealth").
		WithDetail("total", 1024).
		WithDetail("current", 50)

}

func TestHealth(t *testing.T) {
	var indicatorList []health.Indicator
	indicatorList = append(indicatorList, downHealth{}, upHealth{})
	builder := health.NewHealthIndicator(indicatorList)
	m := builder.Build()
	bytes, _ := json.Marshal(m)
	jsonstr := string(bytes)
	assert.Equal(t, jsonstr, `{"components":[{"details":{"reason":"error:down"},"name":"downHealth","status":"down"},{"details":{"current":50,"total":1024},"name":"UpHealth","status":"up"}],"status":"down"}`)
}
