package prefixtree

import (
	"github.com/basp1/pocket/tree"
)

type PrefixTree struct {
	tree *tree.Tree
}

func New() *PrefixTree {
	t := &PrefixTree{}

	t.tree = tree.New()

	return t
}

func (self *PrefixTree) GetTree() *tree.Tree {
	return self.tree
}

func pop(array []PathPoint) (PathPoint, []PathPoint) {
	return array[len(array)-1], array[:len(array)-1]
}

type PathPoint struct {
	vertex int
	edge   int
}

func (self *PrefixTree) AllPaths() [][]interface{} {
	tree := self.tree
	paths := [][]interface{}{}

	path := []PathPoint{}
	queue := []PathPoint{PathPoint{tree.GetRoot(), -1}}

	point := PathPoint{}
	backward := false

	for len(queue) > 0 {
		point, queue = pop(queue)

		if backward {
			parent := tree.GetParent(point.vertex)
			back := path[len(path)-1]
			for len(path) > 0 && parent != tree.GetParent(back.vertex) {
				back, path = pop(path)
			}
			if len(path) > 1 {
				back, path = pop(path)
			}
			backward = false
		}

		path = append(path, point)

		if nil != tree.GetVertex(point.vertex) {
			edgePath := make([]interface{}, 0)
			for _, point := range path {
				if point.edge >= 0 {
					edgePath = append(edgePath, tree.GetEdge(point.edge))
				}
			}
			paths = append(paths, edgePath)

		}

		if !tree.HasSuccessors(point.vertex) {
			backward = true
		}

		successors, edges := tree.GetSuccessors(point.vertex)
		n := len(successors)
		for i := 0; i < n; i++ {
			queue = append(queue, PathPoint{successors[i], edges[i]})
		}
	}

	return paths
}

func (self *PrefixTree) Add(path []interface{}, value interface{}) {
	t := self.tree
	g := t.GetGraph()
	vertex := t.GetRoot()

	n := len(path)

	for i := 0; i < n; i++ {
		part := path[i]

		next := -1
		j := g.From[vertex]
		for j >= 0 {
			if g.Edges[j] == part {
				next = g.To[j]
				break
			}
			j = g.Next[j]
		}

		if next >= 0 {
			vertex = next
		} else {
			vertex = t.Add(vertex, nil, part)
		}
	}

	t.SetVertex(vertex, value)
}

func (self *PrefixTree) Find(path []interface{}) interface{} {
	t := self.tree
	g := t.GetGraph()
	vertex := t.GetRoot()

	n := len(path)

	for i := 0; i < n; i++ {
		part := path[i]

		next := -1
		j := g.From[vertex]
		for j >= 0 {
			if g.Edges[j] == part {
				next = g.To[j]
				break
			}
			j = g.Next[j]
		}

		if next < 0 {
			return nil
		}

		vertex = next
	}

	return t.GetVertex(vertex)
}
