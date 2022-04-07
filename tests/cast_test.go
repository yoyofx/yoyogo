package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/yoyofx/yoyogo/utils"
	"github.com/yoyofx/yoyogo/utils/cast"
	"testing"
)

func TestStringToNumberConvert(t *testing.T) {
	num, _ := cast.Str2Number[int64]("1111")
	assert.Equal(t, num, int64(1111))

	var outNum int64
	_ = cast.Str2NPtr("1111", &outNum)
	assert.Equal(t, outNum, int64(1111))
}

func TestCondition(t *testing.T) {
	a := -1
	assert.Equal(t, utils.IFF(a > 0, a, 0), 0)

	assert.Equal(t, utils.IFN(a > 0, func() int { return a }, func() int { return 0 }), 0)
}
