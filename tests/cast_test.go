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

type Number interface {
	~int | ~int64 | ~float64 | ~float32 | ~uint | ~uint64
}

type Queryable[T Number] []T

func (query Queryable[T]) Count() int {
	return len([]T(query))
}

func (query Queryable[T]) Sum() T {
	var sum T
	for _, value := range query {
		sum += value
	}
	return sum
}
func (query Queryable[T]) ToList() []T {
	return query
}

func (query Queryable[T]) Filter(f func(T) bool) Queryable[T] {
	var result []T
	for _, elem := range query {
		if f(elem) {
			result = append(result, elem)
		}
	}
	return result
}

func TestMySlice(t *testing.T) {
	var s Queryable[int] = []int{0, 1, 2, 3, 4}
	sum := s.Filter(func(elem int) bool {
		return elem%2 == 0 // 奇数
	}).Sum()
	assert.Equal(t, sum, 6)
}

type Vector[T Number] struct {
	data []T
}

func NewVector[T Number](array []T) *Vector[T] {
	return &Vector[T]{data: array}
}

func (v *Vector[T]) Push(x T) { v.data = append([]T(v.data), x) }

func (v *Vector[T]) ToList() []T {
	return v.data
}

func (v *Vector[T]) ToQueryable() Queryable[T] {
	return v.data
}

func TestMyVector(t *testing.T) {
	vector := NewVector([]int{0, 1, 2, 3, 4})
	vector.Push(5)
	vector.Push(6)
	vector.Push(7)
	vector.Push(8)
	vector.Push(9)

	query := vector.ToQueryable().Filter(func(x int) bool {
		return x%2 == 0
	})

	assert.ElementsMatch(t, query.ToList(), []int{0, 2, 4, 6, 8})
	assert.Equal(t, query.Count(), 5)
}
