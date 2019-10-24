package intlist

const (
	NIL   = -1
	EMPTY = -2
)

type Intlist struct {
	values   []int
	ip       int
	length   int
	capacity int
}

func New(capacity int) *Intlist {
	self := &Intlist{}

	self.values = []int{}
	self.ip = NIL
	self.length = 0
	self.capacity = capacity

	for i := 0; i < capacity; i++ {
		self.values = append(self.values, EMPTY)
	}

	return self
}

func (self *Intlist) Resize(newCapacity int) {
	if self.capacity > newCapacity {
		panic("capacity > newCapacity")
	}

	for i := self.capacity; i < newCapacity; i++ {
		self.values = append(self.values, EMPTY)
	}

	self.capacity = newCapacity
}

func (self *Intlist) IsEmpty() bool {
	return 0 == self.length
}

func (self *Intlist) GetTop() int {
	return self.ip
}

func (self *Intlist) Next(key int) int {
	if NIL == key {
		return key
	} else {
		return self.values[key]
	}
}

func (self *Intlist) Contains(key int) bool {
	if key < 0 || key >= self.capacity {
		panic("key not in [0;capacity)")
	}

	return EMPTY != self.values[key]
}

func (self *Intlist) Push(key int) {
	if key < 0 || key >= self.capacity {
		panic("key not in [0;capacity)")
	}

	if self.Contains(key) {
		return
	}

	self.values[key] = self.ip
	self.ip = key

	self.length += 1
}

func (self *Intlist) Pop() int {
	if self.length <= 0 {
		panic("empty intlist")
	}

	key := self.ip
	self.ip = self.values[self.ip]
	self.values[key] = EMPTY
	self.length -= 1

	return key
}

func (self *Intlist) PopAll() []int {
	n := self.length
	values := []int{}

	for i := 0; i < n; i++ {
		values = append(values, self.Pop())
	}

	return values
}

func (self *Intlist) Clear() {
	for self.length > 0 {
		self.Pop()
	}
}
