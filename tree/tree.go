package tree

import (
	"pocket/graph"
)

type Tree struct {
	graph *graph.Graph
	root  int
}

type Node struct {
	parent int
	value  interface{}
}

func New() *Tree {
	self := &Tree{}

	self.graph = graph.New()

	root := &Node{-1, nil}

	self.root = self.graph.AddVertex(root)

	return self
}

func (self *Tree) GetRoot() int {
	return self.root
}

func (self *Tree) GetGraph() *graph.Graph {
	return self.graph
}

func (self *Tree) Add(parent int, value interface{}, edge interface{}) int {
	node := &Node{parent, value}

	vertex := self.graph.AddVertex(node)

	self.graph.AddEdge(parent, vertex, edge)

	return vertex
}

func (self *Tree) GetVertex(vertex int) interface{} {
	node := self.graph.GetVertex(vertex).(*Node)

	return node.value
}

func (self *Tree) GetSuccessors(vertex int) ([]int, []int) {
	return self.graph.GetAdjacent(vertex), self.graph.GetEdges(vertex)
}

func (self *Tree) GetEdge(edge int) interface{} {
	return self.graph.GetEdge(edge)
}

func (self *Tree) HasSuccessors(vertex int) bool {
	return self.graph.HasEdges(vertex)
}

func (self *Tree) GetParent(vertex int) int {
	node := self.graph.GetVertex(vertex).(*Node)

	return node.parent
}

func (self *Tree) IsLeaf(vertex int) bool {
	return !self.HasSuccessors(vertex)
}

func (self *Tree) SetVertex(vertex int, value interface{}) {
	node := self.graph.GetVertex(vertex).(*Node)

	node.value = value
}

func pop(array []int) (int, []int) {
	return array[len(array)-1], array[:len(array)-1]
}

func (self *Tree) AllPaths() [][]int {
	paths := [][]int{}

	path := []int{}
	queue := []int{self.root}

	vertex := -1
	backward := false

	for len(queue) > 0 {
		vertex, queue = pop(queue)

		if backward {
			parent := self.GetParent(vertex)
			back := path[len(path)-1]
			for len(path) > 0 && parent != self.GetParent(back) {
				back, path = pop(path)
			}
			if len(path) > 1 {
				back, path = pop(path)
			}

			backward = false
		}

		path = append(path, vertex)

		if !self.HasSuccessors(vertex) {
			cp := make([]int, len(path))
			copy(cp, path)
			paths = append(paths, cp)
			backward = true
		}

		successors, _ := self.GetSuccessors(vertex)
		for _, suc := range successors {
			queue = append(queue, suc)
		}
	}

	return paths
}
