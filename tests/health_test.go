package tests

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	abshealth "github.com/yoyofx/yoyogo/abstractions/health"
	"github.com/yoyofx/yoyogo/pkg/health"

	"testing"
)

type downHealth struct{}

func (u downHealth) Health() abshealth.ComponentStatus {
	return abshealth.Down("downHealth").WithDetail("reason", "error:down")
}

type upHealth struct{}

func (u upHealth) Health() abshealth.ComponentStatus {
	return abshealth.Up("UpHealth").
		WithDetail("total", 1024).
		WithDetail("current", 50)
}

func TestHealth(t *testing.T) {
	var indicatorList []abshealth.Indicator
	indicatorList = append(indicatorList, downHealth{}, upHealth{}, health.DiskHealthIndicator{})
	builder := abshealth.NewHealthIndicator(indicatorList)
	m := builder.Build()
	bytes, _ := json.Marshal(m)
	jsonstr := string(bytes)
	assert.NotNil(t, jsonstr)
}
