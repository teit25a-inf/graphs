package adjazenzmatrix

// AdjacencyMatrixGraph ist eine Implementierung des Graph-Interfaces, die eine Adjazenzmatrix verwendet, um die Kanten und ihre Gewichte zu speichern.
type AdjacencyMatrixGraph struct {
	// Edges speichert die Gewichte der Kanten zwischen den Knoten.
	// matrix[i][j] enthält das Gewicht der Kante von Knoten i zu Knoten j, oder 0, wenn keine Kante existiert.
	Edges [][]int
}

// EmptyAdjacencyMatrixGraph erstellt einen neuen leeren AdjacencyMatrixGraph mit einer vorgegebenen Anzahl von Knoten.
func EmptyAdjacencyMatrixGraph(numNodes int) *AdjacencyMatrixGraph {
	edges := make([][]int, numNodes)
	for i := range edges {
		edges[i] = make([]int, numNodes)
	}
	return &AdjacencyMatrixGraph{
		Edges: edges,
	}
}

// GetNodes gibt die Liste aller Knoten-IDs im Graphen zurück.
func (g *AdjacencyMatrixGraph) GetNodes() []int {
	// Hinweis:
	// - Iterieren Sie über die Zeilen der Adjazenzmatrix `g.Edges`, um die Knoten-IDs zu sammeln.
	// - Die Knoten-IDs entsprechen den Zeilennummern in der Matrix.

	nodes := make([]int, len(g.Edges))
	// TODO
	return nodes
}

// GetEdges gibt die Liste aller Kanten im Graphen zurück.
// Jeder Eintrag enthält Startknoten, Zielknoten und Gewicht der Kante.
func (g *AdjacencyMatrixGraph) GetEdges() [][3]int {
	// Hinweis:
	// - Iterieren Sie über die Adjazenzmatrix `g.Edges`, um alle Kanten zu identifizieren.
	// - Für jede Zelle, die ein Gewicht ungleich 0 enthält, erstellen Sie einen Eintrag der Form [3]int{start, end, weight}.

	edges := make([][3]int, 0)
	// TODO
	return edges
}

// AddEdge fügt eine gerichtete Kante von start zu end mit dem angegebenen Gewicht hinzu.
func (g *AdjacencyMatrixGraph) AddEdge(start, end, weight int) {
	// Hinweis:
	// - Aktualisieren Sie die Zelle in der Adjazenzmatrix `g.Edges`, die der Kante von `start` zu `end` entspricht, mit dem angegebenen Gewicht.

	// TODO
}

// RemoveEdge entfernt die gerichtete Kante von start zu end, falls sie existiert.
func (g *AdjacencyMatrixGraph) RemoveEdge(start, end int) {
	// Hinweis:
	// - Setzen Sie die Zelle in der Adjazenzmatrix `g.Edges`, die der Kante von `start` zu `end` entspricht, auf 0, um die Kante zu entfernen.

	// TODO
}

// GetNeighbors gibt die Nachbarn eines Knotens zurück, d.h. alle Knoten,
// zu denen eine gerichtete Kante von diesem Knoten aus existiert.
func (g *AdjacencyMatrixGraph) GetNeighbors(node_id int) []int {
	// Hinweis:
	// - Bestimmen Sie die Nachbarn von `node_id` in der Adjazenzmatrix `g.Edges`.
	// - Alle Spalten, die ein Gewicht ungleich 0 enthalten, entsprechen Nachbarn von `node_id`.

	neighbors := make([]int, 0)
	// TODO
	return neighbors
}
