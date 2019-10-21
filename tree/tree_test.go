package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func countSuccessors(t *Tree, vertex int) int {
	successors, _ := t.GetSuccessors(vertex)
	return len(successors)
}

func TestTree(T *testing.T) {
	t := New()

	a := t.Add(t.root, 'a', 'a')
	b := t.Add(t.root, 'b', 'b')
	c := t.Add(a, 'c', 'c')
	d := t.Add(a, 'd', 'd')
	e := t.Add(b, 'e', 'e')
	f := t.Add(c, 'f', 'f')
	g := t.Add(d, 'g', 'g')
	h := t.Add(d, 'h', 'h')
	i := t.Add(d, 'i', 'i')
	j := t.Add(e, 'j', 'j')

	assert.True(T, t.HasSuccessors(t.root))
	assert.Equal(T, 2, countSuccessors(t, t.root))
	assert.Equal(T, 2, countSuccessors(t, a))
	assert.Equal(T, 1, countSuccessors(t, b))
	assert.Equal(T, 1, countSuccessors(t, c))
	assert.Equal(T, 3, countSuccessors(t, d))
	assert.Equal(T, 1, countSuccessors(t, e))
	assert.False(T, t.HasSuccessors(f))
	assert.Equal(T, 0, countSuccessors(t, f))
	assert.False(T, t.HasSuccessors(g))
	assert.Equal(T, 0, countSuccessors(t, g))
	assert.False(T, t.HasSuccessors(h))
	assert.Equal(T, 0, countSuccessors(t, h))
	assert.False(T, t.HasSuccessors(i))
	assert.Equal(T, 0, countSuccessors(t, i))
	assert.False(T, t.HasSuccessors(j))
	assert.Equal(T, 0, countSuccessors(t, j))

	assert.Equal(T, 'a', t.GetVertex(a))
	assert.Equal(T, 'b', t.GetVertex(b))
	assert.Equal(T, 'c', t.GetVertex(c))
	assert.Equal(T, 'd', t.GetVertex(d))
	assert.Equal(T, 'e', t.GetVertex(e))
	assert.Equal(T, 'f', t.GetVertex(f))
	assert.Equal(T, 'g', t.GetVertex(g))
	assert.Equal(T, 'h', t.GetVertex(h))
	assert.Equal(T, 'i', t.GetVertex(i))
	assert.Equal(T, 'j', t.GetVertex(j))

	assert.Equal(T, t.root, t.GetParent(a))
	assert.Equal(T, t.root, t.GetParent(b))
	assert.Equal(T, a, t.GetParent(c))
	assert.Equal(T, a, t.GetParent(d))
	assert.Equal(T, b, t.GetParent(e))
	assert.Equal(T, c, t.GetParent(f))
	assert.Equal(T, d, t.GetParent(g))
	assert.Equal(T, d, t.GetParent(h))
	assert.Equal(T, d, t.GetParent(i))
	assert.Equal(T, e, t.GetParent(j))

	assert.Equal(T, 5, len(t.AllPaths()))
}

func TestAllPaths(T *testing.T) {
	t := New()

	a := t.Add(t.root, 'a', 'a')
	b := t.Add(t.root, 'b', 'b')
	ab := t.Add(a, 'b', 'b')
	abb := t.Add(ab, 'b', 'b')
	abbc := t.Add(abb, 'c', 'c')
	abbd := t.Add(abb, 'd', 'd')

	paths := t.AllPaths()
	assert.Equal(T, 3, len(paths))

	assert.Equal(T, 0, paths[0][0])
	assert.Equal(T, a, paths[0][1])
	assert.Equal(T, ab, paths[0][2])
	assert.Equal(T, abb, paths[0][3])
	assert.Equal(T, abbc, paths[0][4])

	assert.Equal(T, 0, paths[1][0])
	assert.Equal(T, a, paths[1][1])
	assert.Equal(T, ab, paths[1][2])
	assert.Equal(T, abb, paths[1][3])
	assert.Equal(T, abbd, paths[1][4])

	assert.Equal(T, 0, paths[2][0])
	assert.Equal(T, b, paths[2][1])
}
