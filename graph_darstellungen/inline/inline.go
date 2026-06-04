package inline

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
	// TODO
	return nodes
}

// GetEdges gibt die Liste aller Kanten im Graphen zurück.
// Jeder Eintrag enthält Startknoten, Zielknoten und Gewicht der Kante.
func (g *InlineGraph) GetEdges() [][3]int {
	// Hinweis:
	// - Durchsuchen Sie die Liste `g.Nodes`, um alle Knoten und ihre Nachbarn zu erhalten.
	// - Für jeden Knoten und seine Nachbarn erstellen Sie Einträge der Form [3]int{start, end, weight}.

	var edges [][3]int
	// TODO
	return edges
}

// AddEdge fügt eine gerichtete Kante von start zu end mit dem angegebenen Gewicht hinzu.
func (g *InlineGraph) AddEdge(start, end, weight int) {
	// Hinweis:
	// - Finden Sie den Startknoten in der Liste `g.Nodes`.
	// - Wenn er nicht existiert, fügen Sie ihn hinzu.
	// - Fügen Sie die Nachbar- und Gewichtsinformationen zum Startknoten hinzu.

	// TODO
}

// RemoveEdge entfernt die gerichtete Kante von start zu end, falls sie existiert.
func (g *InlineGraph) RemoveEdge(start, end int) {
	// Hinweis:
	// - Finden Sie den Startknoten in der Liste `g.Nodes`.
	// - Wenn er existiert, durchsuchen Sie seine Nachbarn, um die Kante zu `end` zu finden.
	// - Entfernen Sie die Nachbar- und Gewichtsinformationen für diese Kante.

	// TODO
}

// GetNeighbors gibt die Nachbarn eines Knotens zurück, d.h. alle Knoten,
// zu denen eine gerichtete Kante von diesem Knoten aus existiert.
func (g *InlineGraph) GetNeighbors(node_id int) []int {
	// Hinweis:
	// - Durchsuchen Sie die Liste `g.Nodes`, um den Knoten mit der ID `node_id` zu finden.
	// - Wenn der Knoten gefunden wird, geben Sie seine Nachbarn zurück.
	// - Wenn der Knoten nicht gefunden wird, geben Sie eine leere Liste zurück.

	// TODO
	return []int{}
}
