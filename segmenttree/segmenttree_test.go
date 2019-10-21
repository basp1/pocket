package segmenttree

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func selectMinInt(a interface{}, b interface{}) interface{} {
	i := a.(int)
	j := b.(int)
	if i > j {
		return i
	} else {
		return j
	}
}

func TestSegmentTreeGetRangeA(T *testing.T) {
	u := New(int(math.MinInt32), selectMinInt, []interface{}{1})

	assert.Equal(T, 1, u.GetTop())
	assert.Equal(T, 1, u.GetRange(0, 0))
	assert.Equal(T, 1, u.Get(0))
}

func TestSegmentTreeGetRangeB(T *testing.T) {
	u := New(int(math.MinInt32), selectMinInt, []interface{}{3, 8, 6, 4, 2, 5, 9, 0, 7, 1})

	assert.Equal(T, 9, u.GetTop())
	assert.Equal(T, 9, u.GetRange(0, 9))
	assert.Equal(T, 9, u.GetRange(5, 9))
	assert.Equal(T, 8, u.GetRange(0, 4))
	assert.Equal(T, 8, u.GetRange(1, 1))
	assert.Equal(T, 8, u.GetRange(1, 4))
	assert.Equal(T, 7, u.GetRange(7, 8))
	assert.Equal(T, 5, u.GetRange(4, 5))
	assert.Equal(T, 6, u.GetRange(2, 5))
}

func TestSegmentTreeGetRangeC(T *testing.T) {
	u := New(int(math.MinInt32), selectMinInt, []interface{}{5, 4, 3, 2, 1})

	assert.Equal(T, 5, u.GetTop())
	assert.Equal(T, 3, u.GetRange(2, 4))
	assert.Equal(T, 5, u.Get(0))
	assert.Equal(T, 4, u.Get(1))
	assert.Equal(T, 3, u.Get(2))
}

func TestSegmentTreeSetA(T *testing.T) {
	values := []interface{}{3, 8, 6, 4, 2, 5, 9, 0, 7, 1}

	u := New(int(math.MinInt32), selectMinInt, values)

	v := Empty(int(math.MinInt32), selectMinInt, 10)

	assert.False(T, u.Equals(v))

	for i := 0; i < 10; i++ {
		v.Set(i, values[i])
	}

	assert.True(T, u.Equals(v))
}

func TestSegmentTreeSetB(T *testing.T) {
	values := []interface{}{3, 8, 6, 4, 2, 5, 9, 0, 7, 1}

	u := New(int(math.MinInt32), selectMinInt, values)

	u.Set(6, 0)

	assert.Equal(T, 8, u.GetTop())
	assert.Equal(T, 8, u.GetRange(0, 9))
	assert.Equal(T, 7, u.GetRange(5, 9))
	assert.Equal(T, 8, u.GetRange(0, 4))
	assert.Equal(T, 8, u.Get(1))
	assert.Equal(T, 8, u.GetRange(1, 4))
	assert.Equal(T, 7, u.GetRange(7, 8))
	assert.Equal(T, 5, u.GetRange(4, 5))
	assert.Equal(T, 6, u.GetRange(2, 5))
}
