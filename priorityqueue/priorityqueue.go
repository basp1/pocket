package priorityqueue

import (
	"math"
)

type PriorityQueue struct {
	length int
	values []interface{}

	selectFunc func(interface{}, interface{}) interface{}
}

func New(selectFunc func(interface{}, interface{}) interface{}) *PriorityQueue {
	self := &PriorityQueue{}

	self.length = 0
	self.values = []interface{}{}
	self.selectFunc = selectFunc

	return self
}

func (self *PriorityQueue) Push(value interface{}) {
	index := self.length

	if len(self.values) <= index {
		self.values = append(self.values, value)
	} else {
		self.values[index] = value
	}

	self.length += 1

	self.promote(index)
}

func (self *PriorityQueue) Pop() interface{} {
	if self.length <= 0 {
		panic("empty priorityqueue")
	}

	t := self.GetTop()

	if 1 == self.length {
		self.length = 0
	} else {
		last := self.values[self.length-1]
		self.values[0] = last
		self.length -= 1

		self.demote(0)
	}

	return t
}

func (self *PriorityQueue) GetTop() interface{} {
	if self.length <= 0 {
		panic("empty priorityqueue")
	}

	return self.values[0]
}

func (self *PriorityQueue) GetHeight() int {
	return 1 + int(math.Floor(math.Log(float64(self.length))/math.Log(2)))
}

func (self *PriorityQueue) promote(index int) {
	if index < 0 || index >= self.length {
		panic("index not in [0,length)")
	}

	if 0 == index {
		return
	}

	parent := int(math.Floor(float64(index) / 2))

	for index > 0 {
		t := self.values[index]

		if t != self.selectFunc(t, self.values[parent]) {
			break
		}

		self.values[index] = self.values[parent]
		self.values[parent] = t

		next := parent
		parent = int(math.Floor(float64(index) / 2))
		index = next
	}
}

func (self *PriorityQueue) demote(index int) {
	if index < 0 || index >= self.length {
		panic("index not in [0,length)")
	}

	if self.length == (1 + index) {
		return
	}

	value := self.values[index]

	for index < self.length {
		right := (1 + index) * 2
		var rightValue interface{}
		if right < self.length {
			rightValue = self.values[right]
		}

		left := right - 1
		var leftValue interface{}
		if left < self.length {
			leftValue = self.values[left]
		}

		child := -1
		if right < self.length && left < self.length && leftValue == self.selectFunc(leftValue, rightValue) {
			child = left
		} else if right < self.length {
			child = right
		} else if left < self.length {
			child = left
		}

		if child < 0 || value == self.selectFunc(value, self.values[child]) {
			break
		} else {
			self.values[index] = self.values[child]
			self.values[child] = value
			index = child
		}
	}
}
