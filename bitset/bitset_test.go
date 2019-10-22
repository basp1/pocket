package bitset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetA(T *testing.T) {
	b := New(10)
	assert.Equal(T, 1, len(b.GetValues()))

	b.Set(1, true)
	assert.Equal(T, uint64(2), b.GetValues()[0])

	b.Set(0, true)
	assert.Equal(T, uint64(3), b.GetValues()[0])

	b.Set(7, true)
	assert.Equal(T, uint64(131), b.GetValues()[0])

	b.Set(0, false)
	assert.Equal(T, uint64(130), b.GetValues()[0])

	b.Set(1, false)
	assert.Equal(T, uint64(128), b.GetValues()[0])

	b.Set(7, false)
	assert.Equal(T, uint64(0), b.GetValues()[0])
}

func TestSetB(T *testing.T) {
	b := New(1)
	assert.Equal(T, 1, len(b.GetValues()))

	b = New(32)
	assert.Equal(T, 1, len(b.GetValues()))

	b = New(33)
	assert.Equal(T, 2, len(b.GetValues()))

	b = New(63)
	assert.Equal(T, 2, len(b.GetValues()))

	b = New(64)
	assert.Equal(T, 2, len(b.GetValues()))

	b = New(65)
	assert.Equal(T, 3, len(b.GetValues()))
}

func TestSetC(T *testing.T) {
	b := New(100)
	assert.Equal(T, 4, len(b.GetValues()))

	for i := 0; i < 8; i++ {
		b.Set(i, true)
	}
	assert.Equal(T, uint64(255), b.GetValues()[0])

	for i := 8; i < 16; i++ {
		b.Set(i, true)
	}
	assert.Equal(T, uint64(65535), b.GetValues()[0])

	for i := 0; i < 8; i++ {
		b.Set(i, false)
	}
	assert.Equal(T, uint64(65280), b.GetValues()[0])

	for i := 8; i < 16; i++ {
		b.Set(i, false)
	}
	assert.Equal(T, uint64(0), b.GetValues()[0])

	b.Set(16, true)
	assert.Equal(T, uint64(65536), b.GetValues()[0])

	b.Set(33, true)
	assert.Equal(T, uint64(2), b.GetValues()[1])
}

func TestGetA(T *testing.T) {
	b := New(100)
	assert.Equal(T, 4, len(b.GetValues()))

	b.Set(3, true)
	assert.False(T, b.Get(0))
	assert.False(T, b.Get(1))
	assert.False(T, b.Get(2))
	assert.True(T, b.Get(3))
	assert.False(T, b.Get(4))
	assert.False(T, b.Get(5))

	b.Set(50, true)
	assert.False(T, b.Get(0))
	assert.False(T, b.Get(18))
	assert.True(T, b.Get(50))
	assert.False(T, b.Get(72))
}
