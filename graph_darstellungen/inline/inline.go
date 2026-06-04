package inline

import "slices"

// InlineGraph ist eine Implementierung des Graph-Interfaces, die
// die eine Liste von Knoten enthält, wobei jeder Knoten seine Nachbarn
// und die Gewichte der Kanten zu diesen Nachbarn direkt speichert.
type InlineGraph struct {
	// Nodes speichert die Knoten des Graphen. Jeder Eintrag enthält die ID des Knotens,
	// eine Liste der Nachbarn und die entsprechenden Gewichte der Kanten zu diesen Nachbarn.
	Nodes []struct {
		ID        int
		Neighbors []int
		Weights   []int
	}
}

// EmptyInlineGraph erstellt einen neuen leeren InlineGraph.
func EmptyInlineGraph() *InlineGraph {
	return &InlineGraph{
		Nodes: []struct {
			ID        int
			Neighbors []int
			Weights   []int
		}{},
	}
}

// GetNodes gibt die Liste aller Knoten-IDs im Graphen zurück.
func (g *InlineGraph) GetNodes() []int {
	// Hinweis:
	// - Durchsuchen Sie die Liste `g.Nodes`, um die IDs der Knoten zu sammeln.
	// - Extrahieren Sie die ID jedes Knotens und geben Sie sie zurück.

	var nodes []int
	// begin:solution
	for _, node := range g.Nodes {
		nodes = append(nodes, node.ID)
	}
	// end:solution
	return nodes
}

// GetEdges gibt die Liste aller Kanten im Graphen zurück.
// Jeder Eintrag enthält Startknoten, Zielknoten und Gewicht der Kante.
func (g *InlineGraph) GetEdges() [][3]int {
	// Hinweis:
	// - Durchsuchen Sie die Liste `g.Nodes`, um alle Knoten und ihre Nachbarn zu erhalten.
	// - Für jeden Knoten und seine Nachbarn erstellen Sie Einträge der Form [3]int{start, end, weight}.

	var edges [][3]int
	// begin:solution
	for _, node := range g.Nodes {
		for i, neighbor := range node.Neighbors {
			edges = append(edges, [3]int{node.ID, neighbor, node.Weights[i]})
		}
	}
	// end:solution
	return edges
}

// AddEdge fügt eine gerichtete Kante von start zu end mit dem angegebenen Gewicht hinzu.
func (g *InlineGraph) AddEdge(start, end, weight int) {
	// Hinweis:
	// - Finden Sie den Startknoten in der Liste `g.Nodes`.
	// - Wenn er nicht existiert, fügen Sie ihn hinzu.
	// - Fügen Sie die Nachbar- und Gewichtsinformationen zum Startknoten hinzu.

	// begin:solution
	startIndex := slices.IndexFunc(g.Nodes, func(node struct {
		ID        int
		Neighbors []int
		Weights   []int
	}) bool {
		return node.ID == start
	})

	if startIndex == -1 {
		g.Nodes = append(g.Nodes, struct {
			ID        int
			Neighbors []int
			Weights   []int
		}{
			ID:        start,
			Neighbors: []int{},
			Weights:   []int{},
		})
		startIndex = len(g.Nodes) - 1
	}

	startNode := &g.Nodes[startIndex]
	startNode.Neighbors = append(startNode.Neighbors, end)
	startNode.Weights = append(startNode.Weights, weight)
	// end:solution
}

// RemoveEdge entfernt die gerichtete Kante von start zu end, falls sie existiert.
func (g *InlineGraph) RemoveEdge(start, end int) {
	// Hinweis:
	// - Finden Sie den Startknoten in der Liste `g.Nodes`.
	// - Wenn er existiert, durchsuchen Sie seine Nachbarn, um die Kante zu `end` zu finden.
	// - Entfernen Sie die Nachbar- und Gewichtsinformationen für diese Kante.

	// begin:solution
	for i := range g.Nodes {
		if g.Nodes[i].ID == start {
			for j := range g.Nodes[i].Neighbors {
				if g.Nodes[i].Neighbors[j] == end {
					g.Nodes[i].Neighbors = append(g.Nodes[i].Neighbors[:j], g.Nodes[i].Neighbors[j+1:]...)
					g.Nodes[i].Weights = append(g.Nodes[i].Weights[:j], g.Nodes[i].Weights[j+1:]...)
					return
				}
			}
		}
	}
	// end:solution
}

// GetNeighbors gibt die Nachbarn eines Knotens zurück, d.h. alle Knoten,
// zu denen eine gerichtete Kante von diesem Knoten aus existiert.
func (g *InlineGraph) GetNeighbors(node_id int) []int {
	// Hinweis:
	// - Durchsuchen Sie die Liste `g.Nodes`, um den Knoten mit der ID `node_id` zu finden.
	// - Wenn der Knoten gefunden wird, geben Sie seine Nachbarn zurück.
	// - Wenn der Knoten nicht gefunden wird, geben Sie eine leere Liste zurück.

	// begin:solution
	for _, node := range g.Nodes {
		if node.ID == node_id {
			return node.Neighbors
		}
	}
	// end:solution
	return []int{}
}
