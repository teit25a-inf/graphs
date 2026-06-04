# Adjazenzmatrix

Die Adjazenzmatrix ist eine quadratische Matrix, in der die Zeilen und Spalten
die Knoten des Graphen repräsentieren. Ein Eintrag in der Matrix gibt an, ob es eine
zwischen den entsprechenden Knoten gibt, bzw. er enthält das Gewicht der Kante.

## Vorteile

- Effizient für die meisten Operationen, weil jede Kante direkt über die Indizes der
  Knoten gefunden werden kann.

## Nachteile

- Speicherintensiv, weil auch nicht vorhandene Kanten gespeichert werden müssen.
  Dies kommt v.A. zum Tragen, wenn der Graph relativ dünn ist, d.h. wenn es viele
  Knoten gibt, aber nur wenige Kanten.

## Implementierungs-Überlegungen

- Der einfachste Fall ist, die Knoten durch Indizes zu repräsentieren, z.B. 0, 1, 2, ...
- In diesem Fall kann die Adjazenzmatrix einfach ein Array von Arrays sein,
  wobei das Element an Position (i, j) den Wert der Kante zwischen den Knoten i und j enthält.
- Wenn die eigentlichen Knoten Daten enthalten (was i.d.R. der Fall ist), dann könnte
  der Graph zusätzlich zur Adjazenzmatrix eine separate Liste oder Map von Knoten-IDs
  zu Knotendaten enthalten.
