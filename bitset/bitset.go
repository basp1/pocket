package bitset

import (
	"math"
)

type BitSet struct {
	values []uint64
	count  int
}

func New(count int) *BitSet {
	if count <= 0 {
		panic("count should be in [0,)")
	}

	self := &BitSet{}

	n := int(math.Ceil(float64(count) / 32))
	self.count = count
	self.values = make([]uint64, n)

	return self
}

func (self *BitSet) Get(index int) bool {
	if index < 0 || index >= self.count {
		panic("index should be in [0;count)")
	}

	i := index >> 5

	x := self.values[i]

	x >>= uint(index % 32)
	x &= 1

	return (1 == x)
}

func (self *BitSet) Set(index int, value bool) {
	if index < 0 || index >= self.count {
		panic("index should be in [0;count)")
	}

	i := index >> 5
	x := 1 << uint(index%32)
	if value {
		self.values[i] |= uint64(x)
	} else {
		x = ^x
		self.values[i] &= uint64(x)
	}
}

func (self *BitSet) GetValues() []uint64 {
	return self.values
}
