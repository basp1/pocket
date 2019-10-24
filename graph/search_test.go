package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraphDepthSearchA(T *testing.T) {
	graph := New()

	graph.AddVertex(0)
	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)
	graph.AddVertex(4)
	graph.AddVertex(5)
	graph.AddVertex(6)
	graph.AddVertex(7)

	graph.AddEdge(0, 3, nil)
	graph.AddEdge(3, 0, nil)
	graph.AddEdge(3, 6, nil)
	graph.AddEdge(6, 3, nil)
	graph.AddEdge(0, 4, nil)
	graph.AddEdge(4, 0, nil)
	graph.AddEdge(4, 7, nil)
	graph.AddEdge(7, 4, nil)
	graph.AddEdge(4, 5, nil)
	graph.AddEdge(5, 4, nil)
	graph.AddEdge(5, 7, nil)
	graph.AddEdge(7, 5, nil)
	graph.AddEdge(0, 2, nil)
	graph.AddEdge(2, 0, nil)
	graph.AddEdge(0, 1, nil)
	graph.AddEdge(1, 0, nil)
	graph.AddEdge(1, 2, nil)
	graph.AddEdge(2, 1, nil)
	graph.AddEdge(1, 5, nil)
	graph.AddEdge(5, 1, nil)
	graph.AddEdge(2, 5, nil)
	graph.AddEdge(5, 2, nil)

	assert.Equal(T, 8, graph.VertexCount)
	assert.Equal(T, 22, graph.EdgeCount)

	v := graph.DepthSearch(0, func(u int) bool {
		return 1 == graph.GetDegree(u)
	})
	assert.Equal(T, 6, v)

	v = graph.DepthSearch(7, func(u int) bool {
		return 1 == graph.GetDegree(u)
	})
	assert.Equal(T, 6, v)

	v = graph.DepthSearch(2, func(u int) bool {
		return 1 == graph.GetDegree(u)
	})
	assert.Equal(T, 6, v)

	v = graph.DepthSearch(0, func(u int) bool {
		return 7 == u
	})
	assert.Equal(T, 7, v)

	v = graph.DepthSearch(0, func(u int) bool {
		return 4 == u
	})
	assert.Equal(T, 4, v)

	v = graph.DepthSearch(6, func(u int) bool {
		return 0 == u
	})
	assert.Equal(T, 0, v)
}

func TestGraphDepthSearchB(T *testing.T) {
	graph := New()

	graph.AddVertex(0)
	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)
	graph.AddVertex(4)

	graph.AddEdge(0, 1, nil)
	graph.AddEdge(1, 0, nil)
	graph.AddEdge(0, 2, nil)
	graph.AddEdge(2, 0, nil)
	graph.AddEdge(1, 3, nil)
	graph.AddEdge(3, 1, nil)
	graph.AddEdge(2, 4, nil)
	graph.AddEdge(4, 2, nil)

	traversed := []int{}
	graph.DepthSearch(0, func(u int) bool {
		traversed = append(traversed, u)
		return 3 == u
	})

	assert.Equal(T, 3, len(traversed))

	assert.Equal(T, 0, traversed[0])
	assert.Equal(T, 1, traversed[1])
	assert.Equal(T, 3, traversed[2])
}

func TestGraphBreadthSearchA(T *testing.T) {
	graph := New()

	graph.AddVertex(0)
	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)
	graph.AddVertex(4)
	graph.AddVertex(5)
	graph.AddVertex(6)
	graph.AddVertex(7)

	graph.AddEdge(0, 3, nil)
	graph.AddEdge(3, 0, nil)
	graph.AddEdge(3, 6, nil)
	graph.AddEdge(6, 3, nil)
	graph.AddEdge(0, 4, nil)
	graph.AddEdge(4, 0, nil)
	graph.AddEdge(4, 7, nil)
	graph.AddEdge(7, 4, nil)
	graph.AddEdge(4, 5, nil)
	graph.AddEdge(5, 4, nil)
	graph.AddEdge(5, 7, nil)
	graph.AddEdge(7, 5, nil)
	graph.AddEdge(0, 2, nil)
	graph.AddEdge(2, 0, nil)
	graph.AddEdge(0, 1, nil)
	graph.AddEdge(1, 0, nil)
	graph.AddEdge(1, 2, nil)
	graph.AddEdge(2, 1, nil)
	graph.AddEdge(1, 5, nil)
	graph.AddEdge(5, 1, nil)
	graph.AddEdge(2, 5, nil)
	graph.AddEdge(5, 2, nil)

	assert.Equal(T, 8, graph.VertexCount)
	assert.Equal(T, 22, graph.EdgeCount)

	v := graph.BreadthSearch(0, func(u int) bool {
		return 1 == graph.GetDegree(u)
	})
	assert.Equal(T, 6, v)

	v = graph.BreadthSearch(7, func(u int) bool {
		return 1 == graph.GetDegree(u)
	})
	assert.Equal(T, 6, v)

	v = graph.BreadthSearch(2, func(u int) bool {
		return 1 == graph.GetDegree(u)
	})
	assert.Equal(T, 6, v)

	v = graph.BreadthSearch(0, func(u int) bool {
		return 7 == u
	})
	assert.Equal(T, 7, v)

	v = graph.BreadthSearch(0, func(u int) bool {
		return 4 == u
	})
	assert.Equal(T, 4, v)

	v = graph.BreadthSearch(6, func(u int) bool {
		return 0 == u
	})
	assert.Equal(T, 0, v)
}

func TestGraphBreadthSearchB(T *testing.T) {
	graph := New()

	graph.AddVertex(0)
	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)
	graph.AddVertex(4)

	graph.AddEdge(0, 1, nil)
	graph.AddEdge(1, 0, nil)
	graph.AddEdge(0, 2, nil)
	graph.AddEdge(2, 0, nil)
	graph.AddEdge(1, 3, nil)
	graph.AddEdge(3, 1, nil)
	graph.AddEdge(2, 4, nil)
	graph.AddEdge(4, 2, nil)

	traversed := []int{}
	graph.BreadthSearch(0, func(u int) bool {
		traversed = append(traversed, u)
		return 3 == u
	})

	assert.Equal(T, 5, len(traversed))

	assert.Equal(T, 0, traversed[0])
	assert.Equal(T, 2, traversed[1])
	assert.Equal(T, 1, traversed[2])
	assert.Equal(T, 4, traversed[3])
	assert.Equal(T, 3, traversed[4])
}
