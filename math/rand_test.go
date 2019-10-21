package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTree(T *testing.T) {
	values := []interface{}{'a', 'b', 'c'}

	probs := Probs(10, values, []float32{0.2, 0.2, 0.6})

	assert.Equal(T, 'a', probs[0])
	assert.Equal(T, 'a', probs[1])
	assert.Equal(T, 'b', probs[2])
	assert.Equal(T, 'b', probs[3])
	assert.Equal(T, 'c', probs[4])
	assert.Equal(T, 'c', probs[5])
	assert.Equal(T, 'c', probs[6])
	assert.Equal(T, 'c', probs[7])
	assert.Equal(T, 'c', probs[8])
	assert.Equal(T, 'c', probs[9])
}
