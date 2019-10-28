package bytebuffer

import (
	"math"
)

type ByteBuffer struct {
	data []byte

	local    int
	capacity int
}

func New(capacity int) *ByteBuffer {
	self := &ByteBuffer{}

	self.capacity = capacity
	self.data = make([]byte, capacity)
	self.local = 0

	self.Reserve(capacity)

	return self
}

func (self *ByteBuffer) GetPosition() int {
	return self.local
}

func (self *ByteBuffer) SetPosition(newPosition int) {
	self.Reserve(newPosition)
	self.local = newPosition
}

func (self *ByteBuffer) Flip() {
	self.SetPosition(0)
}

func (self *ByteBuffer) ToArray() []byte {
	return self.data[0:self.capacity]
}

func (self *ByteBuffer) Put(value byte) {
	i := self.GetPosition()
	self.SetPosition(i + 1)
	self.put(i, value)
}

func (self *ByteBuffer) PutUint16(value uint16) {
	i := self.GetPosition()
	self.SetPosition(i + 2)
	self.putUint16(i, value)
}

func (self *ByteBuffer) PutUint32(value uint32) {
	i := self.GetPosition()
	self.SetPosition(i + 4)
	self.putUint32(i, value)
}

func (self *ByteBuffer) PutUint64(value uint64) {
	i := self.GetPosition()
	self.SetPosition(i + 8)
	self.putUint64(i, value)
}

func (self *ByteBuffer) put(i int, value byte) {
	if i < 0 || i >= self.capacity {
		panic("'i' should be in [0;capacity)")
	}

	self.data[i] = value
}

func (self *ByteBuffer) putUint16(i int, value uint16) {
	if i < 0 || i >= self.capacity {
		panic("'i' should be in [0;capacity)")
	}
	for j := 0; j < 2; j++ {
		self.data[i+j] = byte(value & 0xff)
		value >>= 8
	}
}

func (self *ByteBuffer) putUint32(i int, value uint32) {
	if i < 0 || i >= self.capacity {
		panic("'i' should be in [0;capacity)")
	}
	for j := 0; j < 4; j++ {
		self.data[i+j] = byte(value & 0xff)
		value >>= 8
	}
}

func (self *ByteBuffer) putUint64(i int, value uint64) {
	if i < 0 || i >= self.capacity {
		panic("'i' should be in [0;capacity)")
	}
	for j := 0; j < 8; j++ {
		self.data[i+j] = byte(value & 0xff)
		value >>= 8
	}
}

func (self *ByteBuffer) Get() byte {
	i := self.GetPosition()
	self.SetPosition(i + 1)
	return self.get(i)
}

func (self *ByteBuffer) GetUint16() uint16 {
	i := self.GetPosition()
	self.SetPosition(i + 2)
	return self.getUint16(i)
}

func (self *ByteBuffer) GetUint32() uint32 {
	i := self.GetPosition()
	self.SetPosition(i + 4)
	return self.getUint32(i)
}

func (self *ByteBuffer) GetUint64() uint64 {
	i := self.GetPosition()
	self.SetPosition(i + 8)
	return self.getUint64(i)
}

func (self *ByteBuffer) get(i int) byte {
	if i < 0 || i >= self.capacity {
		panic("'i' should be in [0;capacity)")
	}
	return self.data[i]
}

func (self *ByteBuffer) getUint16(i int) uint16 {
	if i < 0 || i >= self.capacity {
		panic("'i' should be in [0;capacity)")
	}

	value := uint16(0)
	for j := 0; j < 2; j++ {
		value |= uint16(self.data[i+j] << uint16(j<<3))
	}
	return value
}

func (self *ByteBuffer) getUint32(i int) uint32 {
	if i < 0 || i >= self.capacity {
		panic("'i' should be in [0;capacity)")
	}

	value := uint32(0)
	for j := 0; j < 4; j++ {
		value |= uint32(self.data[i+j] << uint32(j<<3))
	}
	return value
}

func (self *ByteBuffer) getUint64(i int) uint64 {
	if i < 0 || i >= self.capacity {
		panic("'i' should be in [0;capacity)")
	}

	value := uint64(0)
	for j := 0; j < 8; j++ {
		value |= (uint64(self.data[i+j] << uint64(j<<3)))
	}
	return value
}

func (self *ByteBuffer) Reserve(newCapacity int) {
	// Use a long to prevent overflows
	arrayLen := uint64(len(self.data))
	newCap := uint64(newCapacity)

	if newCap > arrayLen {
		var newLen uint64
		if uint64(4) > arrayLen {
			newLen = 8
		} else {
			newLen = arrayLen * 2
		}

		for newCap > newLen {
			newLen <<= 1
		}

		if newLen > math.MaxInt32 {
			newLen = math.MaxInt32
		}

		if uint64(len(self.data)) < newLen {
			tmp := make([]byte, newLen)
			for i := 0; i < self.capacity; i++ {
				tmp[i] = self.data[i]
			}
			self.data = tmp
		}
	}

	self.capacity = newCapacity
}
