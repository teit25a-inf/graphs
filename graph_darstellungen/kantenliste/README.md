# Kantenliste

Dies ist die einfachste Art, einen Graphen zu repräsentieren.
Es wird einfach eine Liste aller Kanten geführt, wobei jede Kante durch die beiden
beteiligten Knoten und ggf. das Gewicht der Kante beschrieben wird. Es ist also eine
einfache Liste von Tupeln oder Tripeln (Knoten1, Knoten2 und ggf. Gewicht).

## Vorteile

- Sehr nahe an der mathematischen Definition eines Graphen, daruch einfach zu verstehen.
- Einfach zu implementieren.
- V.a. bei kleinen Graphen kann die Einfachheit der Kantenliste ein Vorteil sein.

## Nachteile

- Ineffizient für die meisten Operationen, z.B. das Durchlaufen der Nachbarn eines
  Knotens oder die gezielte Prüfung, ob eine Kante zwischen zwei Knoten existiert.
- All diese Operationen erfordern das Durchsuchen der gesamten (unsortierten) Liste
  der Kanten, was bei großen Graphen sehr langsam sein kann.

## Implementierungs-Überlegungen

- Die Knoten können durch Indizes oder durch tatsächliche Daten repräsentiert werden.
- Die Kanten können als Tupel oder Tripel gespeichert werden, z.B. (Knoten1, Knoten2)
  oder (Knoten1, Knoten2, Gewicht).
- Um die Effizienz zu verbessern, könnte man die Kantenliste sortieren oder in einer
  Datenstruktur speichern, die schnellere Suchoperationen ermöglicht, z.B. in Hash-Sets
  oder in einem Baum. Allerdings würde dies die Einfachheit der Kantenliste
  beeinträchtigen und könnte den Speicherbedarf erhöhen.
  Besser wäre dann wahrscheinlich eine andere Darstellungsform,
  z.B. die Adjazenzliste oder die Adjazenzmatrix.
