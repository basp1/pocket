package bytebuffer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByteBufferA(T *testing.T) {
	bb := New(6)
	assert.Equal(T, 0, bb.GetPosition())

	bb.Put(1)
	assert.Equal(T, 1, bb.GetPosition())

	bb.PutUint32(2)
	assert.Equal(T, 5, bb.GetPosition())

	bb.Flip()
	assert.Equal(T, 1, int(bb.Get()))
	assert.Equal(T, 2, int(bb.GetUint32()))

	bb.Flip()
	bb.PutUint32(3)
	bb.PutUint64(4)
	assert.True(T, len(bb.ToArray()) >= 12)
	assert.Equal(T, 12, bb.GetPosition())
	assert.Equal(T, 3, int(bb.getUint32(0)))
	assert.Equal(T, 4, int(bb.getUint64(4)))

	bb.SetPosition(4)
	assert.Equal(T, 4, int(bb.GetUint64()))

	bb.Flip()
	assert.Equal(T, 3, int(bb.GetUint32()))
}
