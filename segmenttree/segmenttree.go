package segmenttree

import (
	"math"
)

type SegmentTree struct {
	capacity   int
	limit      interface{}
	selectFunc func(interface{}, interface{}) interface{}
	values     []interface{}
}

func New(limit interface{}, selectFunc func(interface{}, interface{}) interface{}, values []interface{}) *SegmentTree {

	self := Empty(limit, selectFunc, len(values))

	for i := 0; i < len(values); i++ {
		self.values[self.capacity-1+i] = values[i]
	}

	for i := self.capacity - 2; i >= 0; i-- {
		self.values[i] = selectFunc(self.values[1+(i<<1)], self.values[2+(i<<1)])
	}

	return self
}

func Empty(limit interface{}, selectFunc func(interface{}, interface{}) interface{}, n int) *SegmentTree {
	self := &SegmentTree{}

	self.limit = limit
	self.selectFunc = selectFunc
	self.capacity = 1 << uint(math.Ceil(math.Log2(float64(n))))
	self.values = make([]interface{}, 0)

	for i := 0; i < (self.capacity << 1); i++ {
		self.values = append(self.values, limit)
	}

	return self
}

func (self *SegmentTree) Equals(that *SegmentTree) bool {
	if self.capacity != that.capacity {
		return false
	}

	for i := 0; i < self.capacity; i++ {
		if self.values[i] != that.values[i] {
			return false
		}
	}

	return true
}

func (self *SegmentTree) Set(index int, newValue interface{}) {
	if index >= self.capacity || index < 0 {
		panic("index should be in [0, capacity)")
	}

	index += (self.capacity - 1)
	self.values[index] = newValue

	for index > 0 {
		value := self.values[index]
		var neighbor int
		if 1 == index%2 {
			neighbor = index + 1
			index >>= 1
		} else {
			neighbor = index - 1
			index = (index >> 1) - 1
		}

		newValue = self.selectFunc(value, self.values[neighbor])
		if self.values[index] != newValue {
			self.values[index] = newValue
		} else {
			break
		}
	}
}

func (self *SegmentTree) GetTop() interface{} {
	return self.values[0]
}

func (self *SegmentTree) Get(index int) interface{} {
	return self.GetRange(index, index)
}

func (self *SegmentTree) GetRange(left int, right int) interface{} {
	if left > right {
		panic("left should <= right")
	}

	if left >= self.capacity || left < 0 {
		panic("left should be in [0, capacity)")
	}

	if right >= self.capacity || right < 0 {
		panic("right should be in [0, capacity)")
	}

	left += (self.capacity - 1)
	right += (self.capacity - 1)
	leftValue := self.limit
	rightValue := self.limit

	for left < right {
		if 0 == left%2 {
			leftValue = self.selectFunc(leftValue, self.values[left])
		}

		left >>= 1

		if 1 == right%2 {
			rightValue = self.selectFunc(self.values[right], rightValue)
		}
		right = (right >> 1) - 1
	}

	if left == right {
		leftValue = self.selectFunc(leftValue, self.values[left])
	}

	return self.selectFunc(leftValue, rightValue)
}
