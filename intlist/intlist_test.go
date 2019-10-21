package intlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntlistPush(T *testing.T) {
	ii := New(10)

	ii.Push(1)
	ii.Push(7)
	ii.Push(3)
	ii.Push(2)

	assert.Equal(T, 10, ii.capacity)
	assert.Equal(T, 4, ii.length)
}

func TestIntlistContains(T *testing.T) {
	ii := New(10)

	ii.Push(1)
	ii.Push(7)
	ii.Push(3)
	ii.Push(2)
	ii.Push(0)

	assert.True(T, ii.Contains(1))
	assert.True(T, ii.Contains(2))
	assert.True(T, ii.Contains(3))
	assert.True(T, ii.Contains(7))

	assert.True(T, ii.Contains(0))
	assert.False(T, ii.Contains(5))
}

func TestIntlistPop(T *testing.T) {
	ii := New(10)

	ii.Push(1)
	ii.Push(7)
	ii.Push(3)
	ii.Push(2)

	assert.Equal(T, 2, ii.Pop())
	assert.Equal(T, 3, ii.Pop())
	assert.Equal(T, 7, ii.Pop())
	assert.Equal(T, 1, ii.Pop())
}

func TestIntlistPopAll(T *testing.T) {
	ii := New(10)

	ii.Push(1)
	ii.Push(7)
	ii.Push(3)
	ii.Push(2)

	values := ii.PopAll()

	assert.Equal(T, 0, ii.length)
	assert.Equal(T, 4, len(values))

	assert.Equal(T, 2, values[0])
	assert.Equal(T, 3, values[1])
	assert.Equal(T, 7, values[2])
	assert.Equal(T, 1, values[3])
}
