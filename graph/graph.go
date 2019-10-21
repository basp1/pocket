package graph

const NIL = -1

type Graph struct {
	VertexCount int
	EdgeCount   int
	Free        int

	From []int
	Next []int
	To   []int

	Vertices []interface{}
	Edges    []interface{}
}

func New() *Graph {
	self := &Graph{}

	self.VertexCount = 0
	self.EdgeCount = 0
	self.Free = NIL

	return self
}

func (self *Graph) Clear() {
	self.VertexCount = 0
	self.EdgeCount = 0
	self.Free = NIL

	self.From = self.From[:0]
	self.Next = self.Next[:0]
	self.To = self.To[:0]

	self.Vertices = []interface{}{}
	self.Edges = []interface{}{}
}

func (self *Graph) AddVertex(vertexValue interface{}) int {
	self.VertexCount += 1

	self.From = append(self.From, NIL)
	self.Vertices = append(self.Vertices, vertexValue)

	return self.VertexCount - 1
}

func (self *Graph) GetVertex(vertex int) interface{} {
	if vertex < 0 || vertex >= self.VertexCount {
		panic("vertex not in [0; VertexCount)")
	}

	return self.Vertices[vertex]
}

func (self *Graph) SetVertex(vertex int, vertexValue interface{}) {
	if vertex < 0 || vertex >= self.VertexCount {
		panic("vertex not in [0; VertexCount)")
	}

	self.Vertices[vertex] = vertexValue
}

func (self *Graph) HasEdges(vertex int) bool {
	if vertex < 0 || vertex >= self.VertexCount {
		panic("vertex not in [0; VertexCount)")
	}

	if NIL == self.From[vertex] {
		return false
	} else {
		return true
	}
}

func (self *Graph) HasEdge(From int, To int) bool {
	if From < 0 || From >= self.VertexCount {
		panic("vertex not in [0; VertexCount)")
	}
	if To < 0 {
		panic("vertex not in [0; VertexCount)")
	}

	if !self.HasEdges(From) {
		return false
	}

	j := self.From[From]
	for NIL != j {
		if To == self.To[j] {
			return true
		}

		j = self.Next[j]
	}

	return false
}

func (self *Graph) AddEdge(From int, To int, edgeValue interface{}) {
	if From < 0 || From >= self.VertexCount {
		panic("vertex not in [0; VertexCount)")
	}
	if To < 0 {
		panic("vertex not in [0; VertexCount)")
	}

	p := 0
	if self.Free >= 0 {
		p = self.Free

		self.To[self.Free] = To
		self.Edges[self.Free] = edgeValue
		self.Free = self.Next[self.Free]
	} else {
		p = self.EdgeCount

		self.Next = append(self.Next, NIL)
		self.To = append(self.To, To)
		self.Edges = append(self.Edges, edgeValue)
	}

	self.Next[p] = self.From[From]
	self.From[From] = p

	self.EdgeCount += 1
}

func (self *Graph) RemoveEdge(From int, To int) {
	if From < 0 || From >= self.VertexCount {
		panic("vertex not in [0; VertexCount)")
	}
	if To < 0 {
		panic("vertex not in [0; VertexCount)")
	}

	if !self.HasEdges(From) {
		return
	}

	k := NIL
	p := NIL
	j := self.From[From]

	for NIL != j {
		if To == self.To[j] {
			k = j
			break
		}
		p = j
		j = self.Next[j]
	}

	if NIL == k {
		return
	}

	if k == self.From[From] {
		self.From[From] = self.Next[k]
		self.Next[k] = self.Free
		self.Free = k
	} else {
		self.Next[p] = self.Next[k]
		self.Next[k] = self.Free
		self.Free = k
	}

	self.EdgeCount -= 1
}

func (self *Graph) RemoveEdges(vertex int) {
	if vertex < 0 || vertex >= self.VertexCount {
		panic("vertex not in [0; VertexCount)")
	}

	if !self.HasEdges(vertex) {
		return
	}

	VertexCount := 1
	p := self.From[vertex]

	for NIL != self.Next[p] {
		p = self.Next[p]
		VertexCount += 1
	}

	self.Next[p] = self.Free
	self.Free = self.From[vertex]
	self.From[vertex] = NIL

	self.EdgeCount -= VertexCount
}

func (self *Graph) IsLeaf(vertex int) bool {
	if vertex < 0 || vertex >= self.VertexCount {
		panic("vertex not in [0; VertexCount)")
	}

	if !self.HasEdges(vertex) {
		return true
	}

	first := vertex

	i := self.From[vertex]
	for NIL != i {
		if vertex != self.To[i] {
			first = self.To[i]
			break
		}
		i = self.Next[i]
	}

	if vertex == first {
		return true
	}

	i = self.From[vertex]
	for NIL != i {
		if first != self.To[i] && vertex != self.To[i] {
			return false
		}
		i = self.Next[i]
	}

	return true
}

func (self *Graph) Copy() *Graph {
	h := New()

	h.From = make([]int, len(self.From))
	h.Next = make([]int, len(self.Next))
	h.To = make([]int, len(self.To))
	copy(h.From, self.From)
	copy(h.Next, self.Next)
	copy(h.To, self.To)

	h.Vertices = make([]interface{}, len(self.Vertices))
	h.Edges = make([]interface{}, len(self.Edges))
	copy(h.Vertices, self.Vertices)
	copy(h.Edges, self.Edges)

	h.VertexCount = self.VertexCount
	h.EdgeCount = self.EdgeCount
	h.Free = self.Free

	return h
}

func (self *Graph) Equal(h *Graph) bool {
	if nil == h {
		panic("null pointer")
	}

	if self.VertexCount != h.VertexCount {
		return false
	}
	if self.EdgeCount != h.EdgeCount {
		return false
	}

	for i := 0; i < self.VertexCount; i++ {
		j := self.From[i]
		k := h.From[i]

		if self.Vertices[i] != h.Vertices[i] {
			return false
		}

		for NIL != j && NIL != k {
			if self.To[j] != h.To[k] {
				return false
			}
			if self.Edges[j] != h.Edges[k] {
				return false
			}

			j = self.Next[j]
			k = h.Next[k]
		}

		if NIL != j || NIL != k {
			return false
		}
	}

	return true
}

func (self *Graph) GetEdge(edge int) interface{} {
	return self.Edges[edge]
}

func (self *Graph) GetAdjacent(vertex int) []int {
	if vertex < 0 || vertex >= self.VertexCount {
		panic("vertex not in [0; VertexCount)")
	}

	adjacent := []int{}

	j := self.From[vertex]

	for NIL != j {
		adjacent = append(adjacent, self.To[j])
		j = self.Next[j]
	}

	return adjacent
}

func (self *Graph) GetEdges(vertex int) []int {
	if vertex < 0 || vertex >= self.VertexCount {
		panic("vertex not in [0; VertexCount)")
	}

	edges := []int{}

	j := self.From[vertex]

	for NIL != j {
		edges = append(edges, j)
		j = self.Next[j]
	}

	return edges
}
