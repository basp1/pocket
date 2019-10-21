package prefixtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func P(str string) []interface{} {
	arr := []interface{}{}

	for i := 0; i < len(str); i++ {
		arr = append(arr, str[i])
	}

	return arr
}

func TestPrefixTreeAdd(T *testing.T) {
	p := New()

	p.Add(P("a"), "A")
	p.Add(P("a"), "!")
	p.Add(P("b"), "B")
	p.Add(P("ab"), "AB")
	p.Add(P("abbc"), "ABBC")
	p.Add(P("abbd"), "ABBD")

	paths := p.AllPaths()
	assert.Equal(T, 5, len(paths))

	assert.Equal(T, 1, len(paths[0]))
	assert.Equal(T, 2, len(paths[1]))
	assert.Equal(T, 4, len(paths[2]))
	assert.Equal(T, 4, len(paths[3]))
	assert.Equal(T, 1, len(paths[4]))
}

func TestPrefixTreeFind(T *testing.T) {
	p := New()

	p.Add(P("a"), "A")
	p.Add(P("a"), "!")
	p.Add(P("b"), "B")
	p.Add(P("ab"), "AB")
	p.Add(P("abbc"), "ABBC")
	p.Add(P("abbd"), "ABBD")

	assert.Equal(T, "!", p.Find(P("a")))
	assert.Equal(T, "B", p.Find(P("b")))
	assert.Equal(T, "AB", p.Find(P("ab")))
	assert.Equal(T, "ABBC", p.Find(P("abbc")))
	assert.Equal(T, "ABBD", p.Find(P("abbd")))
}

func TestPrefixTreeAllPaths(T *testing.T) {
	p := New()

	p.Add(P("a"), "A")
	p.Add(P("a"), "!")
	p.Add(P("b"), "B")
	p.Add(P("ab"), "AB")
	p.Add(P("abbc"), "ABBC")
	p.Add(P("abbd"), "ABBD")

	paths := p.AllPaths()

	assert.Equal(T, "!", p.Find(paths[0]))
	assert.Equal(T, "AB", p.Find(paths[1]))
	assert.Equal(T, "ABBC", p.Find(paths[2]))
	assert.Equal(T, "ABBD", p.Find(paths[3]))
	assert.Equal(T, "B", p.Find(paths[4]))
}
