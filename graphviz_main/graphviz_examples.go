package main

import (
	"fmt"
	"graphs/graph"
	"graphs/graph_darstellungen/adjazenzliste"
)

func main() {
	g1 := adjazenzliste.EmptyAdjacencyListGraph()
	g1.AddEdge(0, 1, 5)
	g1.AddEdge(0, 2, 3)
	g1.AddEdge(1, 2, 1)

	fmt.Println(graph.DotString(g1))
}

// Die Ausgabe dieses Programms kann mit dem Programm "dot" aus dem Graphviz-Paket visualisiert werden.
// Dies ist frei verfügbar unter https://graphviz.org/download/.
// Alternativ kann die Ausgabe auch auf https://dreampuf.github.io/GraphvizOnline/ visualisiert werden.
