package graphutils

type Graph struct {
	Nodes []Node
	Edges []Edge
}

func NewGraph() *Graph {
	return &Graph{}
}

func (g *Graph) HasNode(id string) bool {
	for _, node := range g.Nodes {
		if node.ID == id {
			return true
		}
	}
	return false
}

func (g *Graph) FindEdges(id string) (result []Edge) {
	for _, edge := range g.Edges {
		if edge.From == id {
			result = append(result, edge)
		}
	}
	return result
}

func (g *Graph) AddNode(node Node) {
	g.Nodes = append(g.Nodes, node)
}

func (g *Graph) AddEdge(edge Edge) {
	if !g.HasNode(edge.From) {
		node := Node{ID: edge.From}
		g.AddNode(node)
	}
	if !g.HasNode(edge.To) {
		node := Node{ID: edge.To}
		g.AddNode(node)
	}

	// find edge & update if it exists
	found1 := false
	found2 := false
	for i, e := range g.Edges {
		if e.From == edge.From && e.To == edge.To {
			g.Edges[i].Weight = edge.Weight
			found1 = true
		}
		if e.To == edge.From && e.From == edge.To {
			g.Edges[i].Weight = edge.Weight
			found2 = true
		}
	}

	if !found1 {
		g.Edges = append(g.Edges, edge)
	}
	if !found2 {
		e := Edge{From: edge.To, To: edge.From, Weight: edge.Weight}
		g.Edges = append(g.Edges, e)
	}
}
