package kantenliste

// EdgeListGraph repräsentiert einen gerichteten Graphen als Liste von Kanten.
// Jede Kante wird als Tripel (Startknoten, Zielknoten, Gewicht) dargestellt.
// Die Knoten sind dabei numerisch indiziert, z.B. 0, 1, 2, ...
type EdgeListGraph struct {
	Edges [][3]int
}

// EmptyEdgeListGraph erstellt einen neuen leeren EdgeListGraph.
func EmptyEdgeListGraph() *EdgeListGraph {
	return &EdgeListGraph{
		Edges: make([][3]int, 0),
	}
}

// GetNodes gibt die Liste aller Knoten-IDs im Graphen zurück.
func (g *EdgeListGraph) GetNodes() []int {
	// Hinweis:
	// - Durchsuchen Sie die Liste `g.Edges`, um alle Knoten-IDs zu sammeln.
	// - Da Knoten sowohl als Start- als auch als Zielknoten in den Kanten auftreten können, sollten Sie eine Menge verwenden, um Duplikate zu vermeiden.
	// - Eine Menge können Sie in Go mit einer Map realisieren, z.B. `map[int]struct{}` oder `map[int]bool`.

	nodeSet := make(map[int]struct{})
	for _, edge := range g.Edges {
		nodeSet[edge[0]] = struct{}{}
		nodeSet[edge[1]] = struct{}{}
	}

	nodes := make([]int, 0, len(nodeSet))
	for node := range nodeSet {
		nodes = append(nodes, node)
	}

	return nodes
}

// GetEdges gibt die Liste aller Kanten im Graphen zurück.
// Jeder Eintrag enthält Startknoten, Zielknoten und Gewicht der Kante.
func (g *EdgeListGraph) GetEdges() [][3]int {
	// Hinweis:
	// - Die Kanten sind bereits in der Liste `g.Edges` gespeichert, daher können Sie diese direkt zurückgeben.

	// begin:solution
	return g.Edges
	// end:solution
}

// AddEdge fügt eine gerichtete Kante von start zu end mit dem angegebenen Gewicht hinzu.
func (g *EdgeListGraph) AddEdge(start, end, weight int) {
	// Hinweis:
	// - Fügen Sie die neue Kante als Tripel (start, end, weight) zur Liste `g.Edges` hinzu.

	// begin:solution
	g.Edges = append(g.Edges, [3]int{start, end, weight})
	// end:solution
}

// RemoveEdge entfernt die gerichtete Kante von start zu end, falls sie existiert.
func (g *EdgeListGraph) RemoveEdge(start, end int) {
	// Hinweis:
	// - Durchsuchen Sie die Liste `g.Edges`, um die Kante von `start` zu `end` zu finden.
	// - Wenn die Kante gefunden wird, entfernen Sie sie aus der Liste.

	// begin:solution
	for i, edge := range g.Edges {
		if edge[0] == start && edge[1] == end {
			g.Edges = append(g.Edges[:i], g.Edges[i+1:]...)
			return
		}
	}
	// end:solution
}

// GetNeighbors gibt die Nachbarn eines Knotens zurück, d.h. alle Knoten,
// zu denen eine gerichtete Kante von diesem Knoten aus existiert.
func (g *EdgeListGraph) GetNeighbors(node_id int) []int {
	// Hinweis:
	// - Durchsuchen Sie die Liste `g.Edges`, um alle Kanten zu finden, die von `node_id` ausgehen.
	// - Sammeln Sie die Zielknoten dieser Kanten und geben Sie sie zurück.

	neighbors := make([]int, 0)
	// begin:solution
	for _, edge := range g.Edges {
		if edge[0] == node_id {
			neighbors = append(neighbors, edge[1])
		}
	}
	// end:solution
	return neighbors
}
