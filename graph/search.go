package graph

import (
	"container/list"

	"github.com/basp1/pocket/util"
)

func (self *Graph) DepthSearch(begin int, predicate func(int) bool) int {
	self.intlist.Resize(self.VertexCount)
	self.intlist.Clear()
	traversed := self.intlist

	var vertex int
	stack := []int{begin}

	for len(stack) > 0 {
		vertex, stack = util.PopInt(stack)

		if traversed.Contains(vertex) {
			continue
		}

		if predicate(vertex) {
			return vertex
		}

		traversed.Push(vertex)

		for _, neighbor := range self.GetAdjacent(vertex) {
			stack = append(stack, neighbor)
		}
	}

	return -1
}

func (self *Graph) BreadthSearch(begin int, predicate func(int) bool) int {
	self.intlist.Resize(self.VertexCount)
	self.intlist.Clear()
	traversed := self.intlist

	var vertex int
	dequeue := list.New()
	dequeue.PushBack(begin)

	for dequeue.Len() > 0 {
		front := dequeue.Front()
		vertex = front.Value.(int)
		dequeue.Remove(front)

		if traversed.Contains(vertex) {
			continue
		}

		if predicate(vertex) {
			return vertex
		}

		traversed.Push(vertex)

		for _, neighbor := range self.GetAdjacent(vertex) {
			dequeue.PushBack(neighbor)
		}
	}

	return -1
}
