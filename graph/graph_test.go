package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraphA(T *testing.T) {
	g := New()

	a := g.AddVertex('a')
	b := g.AddVertex('b')
	c := g.AddVertex('c')

	g.AddEdge(a, a, '-')
	g.AddEdge(b, a, '-')

	assert.True(T, g.HasEdge(a, a))
	assert.True(T, g.HasEdge(b, a))

	assert.False(T, g.HasEdge(a, b))
	assert.False(T, g.HasEdge(b, b))
	assert.False(T, g.HasEdge(c, a))
	assert.False(T, g.HasEdge(c, b))
}

func TestGraphB(T *testing.T) {
	g := New()

	a := g.AddVertex('a')
	b := g.AddVertex('b')
	c := g.AddVertex('c')

	g.AddEdge(a, a, '-')
	g.AddEdge(a, b, '-')
	g.AddEdge(b, a, '-')

	assert.Equal(T, 3, g.EdgeCount)

	g.AddEdge(b, b, '-')
	assert.Equal(T, 4, g.EdgeCount)

	assert.True(T, g.HasEdge(a, b))

	g.RemoveEdge(a, b)

	assert.False(T, g.HasEdge(a, b))

	g.AddEdge(c, a, '-')
	g.AddEdge(c, b, '-')

	assert.True(T, g.HasEdge(a, a))
	assert.True(T, g.HasEdge(b, a))
	assert.True(T, g.HasEdge(b, b))
	assert.True(T, g.HasEdge(c, a))
	assert.True(T, g.HasEdge(c, b))

	assert.False(T, g.HasEdge(a, b))
}

func TestGraphC(T *testing.T) {
	g := New()

	a := g.AddVertex('a')
	b := g.AddVertex('b')
	c := g.AddVertex('c')

	g.AddEdge(a, a, '-')
	g.AddEdge(a, b, '-')
	g.AddEdge(b, a, '-')
	g.AddEdge(b, b, '-')
	g.AddEdge(c, a, '-')
	g.AddEdge(c, b, '-')

	g.RemoveEdge(a, b)
	g.RemoveEdge(b, b)
	g.RemoveEdge(c, a)
	g.RemoveEdge(c, b)

	g.RemoveEdge(a, a)
	g.AddEdge(a, a, '-')

	e := New()
	a = e.AddVertex('a')
	b = e.AddVertex('b')
	c = e.AddVertex('c')

	e.AddEdge(a, a, '-')
	e.AddEdge(b, a, '-')

	assert.True(T, e.Equal(g))
}

func TestGraphD(T *testing.T) {
	g := New()

	a := g.AddVertex('a')
	b := g.AddVertex('b')
	c := g.AddVertex('c')

	g.AddEdge(a, a, '-')
	g.AddEdge(a, b, '-')
	g.AddEdge(b, a, '-')
	g.AddEdge(b, b, '-')
	g.AddEdge(c, a, '-')
	g.AddEdge(c, b, '-')

	assert.Equal(T, 6, g.EdgeCount)

	e := g.Copy()

	g.RemoveEdge(a, a)
	g.RemoveEdge(a, b)
	g.RemoveEdge(b, a)
	g.RemoveEdge(b, b)
	g.RemoveEdge(c, a)
	g.RemoveEdge(c, b)

	assert.Equal(T, 0, g.EdgeCount)

	g.AddEdge(a, a, '-')
	g.AddEdge(a, b, '-')
	g.AddEdge(b, a, '-')
	g.AddEdge(b, b, '-')
	g.AddEdge(c, a, '-')
	g.AddEdge(c, b, '-')

	assert.True(T, e.Equal(g))
}

func TestGraphE(T *testing.T) {
	g := New()

	a := g.AddVertex('a')
	b := g.AddVertex('b')
	c := g.AddVertex('c')

	g.AddEdge(a, a, '-')
	g.AddEdge(a, b, '-')
	g.AddEdge(b, a, '-')
	g.AddEdge(b, b, '-')
	g.AddEdge(c, a, '-')
	g.AddEdge(c, b, '-')

	assert.Equal(T, 6, g.EdgeCount)

	e := g.Copy()

	g.RemoveEdges(a)
	g.RemoveEdges(b)
	g.RemoveEdges(c)

	assert.Equal(T, 0, g.EdgeCount)

	g.AddEdge(a, a, '-')
	g.AddEdge(a, b, '-')
	g.AddEdge(b, a, '-')
	g.AddEdge(b, b, '-')
	g.AddEdge(c, a, '-')
	g.AddEdge(c, b, '-')

	assert.True(T, e.Equal(g))
}
