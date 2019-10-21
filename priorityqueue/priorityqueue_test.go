package priorityqueue

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriorityQueuetPushA(T *testing.T) {
	pq := New(func(a, b interface{}) interface{} {
		i := a.(int)
		j := b.(int)
		if i < j {
			return i
		} else {
			return j
		}
	})

	pq.Push(3)
	assert.Equal(T, 3, pq.GetTop().(int))
	assert.Equal(T, 1, pq.GetHeight())

	pq.Push(4)
	assert.Equal(T, 3, pq.GetTop().(int))
	assert.Equal(T, 2, pq.GetHeight())

	pq.Push(5)
	assert.Equal(T, 3, pq.GetTop().(int))
	assert.Equal(T, 2, pq.GetHeight())

	pq.Push(2)
	assert.Equal(T, 2, pq.GetTop().(int))
	assert.Equal(T, 3, pq.GetHeight())

	pq.Push(1)
	assert.Equal(T, 1, pq.GetTop().(int))
	assert.Equal(T, 3, pq.GetHeight())
}

func TestPriorityQueuetPushB(T *testing.T) {
	pq := New(func(a, b interface{}) interface{} {
		i := a.(int)
		j := b.(int)
		if i < j {
			return i
		} else {
			return j
		}
	})

	for i := 8; i >= 0; i-- {
		pq.Push(i)
	}

	assert.Equal(T, 0, pq.GetTop().(int))
	assert.Equal(T, 4, pq.GetHeight())

	pq.Push(9)
	assert.Equal(T, 0, pq.GetTop().(int))
	assert.Equal(T, 4, pq.GetHeight())
}

func TestPriorityQueuetPopA(T *testing.T) {
	pq := New(func(a, b interface{}) interface{} {
		i := a.(int)
		j := b.(int)
		if i < j {
			return i
		} else {
			return j
		}
	})

	pq.Push(18)
	pq.Push(19)
	pq.Push(20)
	assert.Equal(T, 18, pq.GetTop().(int))

	pq.Pop()
	assert.Equal(T, 19, pq.GetTop().(int))

	pq.Pop()
	assert.Equal(T, 20, pq.GetTop().(int))

	pq.Pop()
	assert.Equal(T, 0, pq.length)
}

func TestPriorityQueuetPopB(T *testing.T) {
	pq := New(func(a, b interface{}) interface{} {
		i := a.(int)
		j := b.(int)
		if i < j {
			return i
		} else {
			return j
		}
	})

	for i := 9; i >= 1; i-- {
		pq.Push(i)
	}

	assert.Equal(T, 1, pq.GetTop().(int))
	assert.Equal(T, 4, pq.GetHeight())

	pq.Pop()
	assert.Equal(T, 2, pq.GetTop().(int))
	assert.Equal(T, 4, pq.GetHeight())

	pq.Pop()
	assert.Equal(T, 3, pq.GetTop().(int))
	assert.Equal(T, 3, pq.GetHeight())

	pq.Pop()
	assert.Equal(T, 4, pq.GetTop().(int))
	assert.Equal(T, 3, pq.GetHeight())

	pq.Pop()
	assert.Equal(T, 5, pq.GetTop().(int))
	assert.Equal(T, 3, pq.GetHeight())
}

func TestPriorityQueuetPopC(T *testing.T) {
	pq := New(func(a, b interface{}) interface{} {
		i := a.(int)
		j := b.(int)
		if i < j {
			return i
		} else {
			return j
		}
	})

	N := 20

	for i := N; i >= 1; i-- {
		pq.Push(i)
	}

	assert.Equal(T, 1, pq.GetTop().(int))

	for i := 0; i < int(math.Floor(float64(N)/2)); i++ {
		pq.Pop()
	}

	assert.Equal(T, int(math.Floor(float64(N)/2+1)), pq.GetTop().(int))

	for i := 0; i < int(math.Floor(float64(N)/2-1)); i++ {
		pq.Pop()
	}

	assert.Equal(T, N, pq.GetTop().(int))
}
