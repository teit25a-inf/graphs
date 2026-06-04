package adjazenzliste

// AdjacencyListGraph repräsentiert einen gerichteten Graphen als Adjazenzliste.
// Jeder Eintrag in der Liste enthält die Nachbarn eines Knotens
// und die Gewichteder Kanten.
type AdjacencyListGraph struct {
	AdjList map[int][][2]int // Knoten-ID -> Liste von (Nachbar-Knoten-ID, Gewicht)
}

// EmptyAdjacencyListGraph erstellt einen neuen leeren AdjacencyListGraph.
func EmptyAdjacencyListGraph() *AdjacencyListGraph {
	return &AdjacencyListGraph{
		AdjList: make(map[int][][2]int),
	}
}

// GetNodes gibt die Liste aller Knoten-IDs im Graphen zurück.
func (g *AdjacencyListGraph) GetNodes() []int {
	// Hinweis:
	// - Durchsuchen Sie die Schlüssel in `g.AdjList`, um die Knoten-IDs zu sammeln.
	// - Die Schlüssel können Sie mit einer range-Schleife über die Map erhalten.

	nodes := make([]int, 0, len(g.AdjList))
	// begin:solution
	for node := range g.AdjList {
		nodes = append(nodes, node)
	}
	// end:solution
	return nodes
}

// GetEdges gibt die Liste aller Kanten im Graphen zurück.
// Jeder Eintrag enthält Startknoten, Zielknoten und Gewicht der Kante.
func (g *AdjacencyListGraph) GetEdges() [][3]int {
	// Hinweis:
	// - Durchsuchen Sie die Map `g.AdjList`, um alle Knoten und ihre Nachbarn zu erhalten.
	// - Für jeden Knoten und seine Nachbarn erstellen Sie Einträge der Form [3]int{start, end, weight}.

	edges := make([][3]int, 0)
	// begin:solution
	for start, neighbors := range g.AdjList {
		for _, neighbor := range neighbors {
			end := neighbor[0]
			weight := neighbor[1]
			edges = append(edges, [3]int{start, end, weight})
		}
	}
	// end:solution
	return edges
}

// AddEdge fügt eine gerichtete Kante von start zu end mit dem angegebenen Gewicht hinzu.
func (g *AdjacencyListGraph) AddEdge(start, end, weight int) {
	// Hinweis:
	// - Fügen Sie die Kante in der Map `g.AdjList` hinzu, indem Sie den Eintrag für `start` aktualisieren.
	// - Wenn `start` noch nicht als Schlüssel in der Map existiert, wird er automatisch hinzugefügt.

	// begin:solution
	g.AdjList[start] = append(g.AdjList[start], [2]int{end, weight})
	// end:solution
}

// RemoveEdge entfernt die gerichtete Kante von start zu end, falls sie existiert.
func (g *AdjacencyListGraph) RemoveEdge(start, end int) {
	// Hinweis:
	// - Bestimmen Sie die Nachbarn von `start` in der Map `g.AdjList`.
	// - Durchsuchen Sie die Nachbarn, um die Kante zu `end` zu finden und entfernen Sie sie aus der Liste.

	neighbors := g.AdjList[start]
	// begin:solution
	for i, neighbor := range neighbors {
		if neighbor[0] == end {
			g.AdjList[start] = append(neighbors[:i], neighbors[i+1:]...)
			return
		}
	}
	// end:solution
}

// GetNeighbors gibt die Nachbarn eines Knotens zurück, d.h. alle Knoten,
// zu denen eine gerichtete Kante von diesem Knoten aus existiert.
func (g *AdjacencyListGraph) GetNeighbors(node_id int) []int {
	// Hinweis:
	// - Bestimmen Sie die Nachbarn von `node_id` in der Map `g.AdjList`.
	// - Extrahieren Sie die Zielknoten aus den Nachbarinformationen und geben Sie sie zurück.

	neighbors := make([]int, 0)
	// begin:solution
	for _, neighbor := range g.AdjList[node_id] {
		neighbors = append(neighbors, neighbor[0])
	}
	// end:solution
	return neighbors
}
