package graph

// Degree gibt die Anzahl der ausgehenden Kanten eines Knotens zurück.
func Degree(g GraphRepr, node_id int) int {
	// Hinweis:
	// - Verwenden Sie die Methode `GetNeighbors`, um die Nachbarn von `node_id` zu erhalten.
	// - Die Anzahl der Nachbarn entspricht der Anzahl der ausgehenden Kanten von `node_id`.

	// TODO
	return 0
}

// HasEdge prüft, ob eine gerichtete Kante von start zu end existiert.
func HasEdge(g GraphRepr, start, end int) bool {
	// Hinweis:
	// - Verwenden Sie die Methode `GetNeighbors`, um die Nachbarn von `start` zu erhalten.
	// - Überprüfen Sie, ob `end` in der Liste der Nachbarn von `start` enthalten ist.

	// TODO
	return false
}

// GetEdgeWeight gibt das Gewicht der gerichteten Kante von start zu end zurück,
// oder 0, wenn die Kante nicht existiert.
func GetEdgeWeight(g GraphRepr, start, end int) int {
	// Hinweis:
	// - Durchsuchen Sie die Liste der Kanten, die Sie mit `GetEdges` erhalten, um die Kante von `start` zu `end` zu finden.
	// - Wenn die Kante gefunden wird, geben Sie ihr Gewicht zurück.
	// - Wenn die Kante nicht gefunden wird, geben Sie -1 zurück.

	// TODO
	return 0
}

// DotString gibt eine String-Darstellung des Graphen im DOT-Format zurück,
// die zur Visualisierung mit Graphviz verwendet werden kann.
func DotString(g GraphRepr) string {
	// Hinweis:
	// - Durchsuchen Sie die Liste der Kanten, die Sie mit `GetEdges` erhalten, und fügen Sie für jede Kante eine Zeile mittels `fmt.Sprintf` hinzu.

	var result string
	// TODO
	return result
}
