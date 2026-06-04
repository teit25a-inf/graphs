# Inline-Darstellung

Mit _Inline-Darstellung_ ist gemeint, dass es Structs für die Knoten gibt, die direkt
Zeiger auf die Nachbarn enthalten.
D.h. Knoten enthalten neben ihren Daten auch direkt die Informationen über ihre
Nachbarn, z.B. in Form von einer Liste von (Zeigern auf die) Nachbarknoten.

## Vor- und Nachteile

- Die Vor- und Nachteile sind i.W. die gleichen wie bei der Adjazenzliste, weil die
  Idee bei dieser Darstellung die gleiche ist.
- Die Implementierung ist etwas weniger flexibel und potentiell unübersichtlicher
  als bei den anderen Darstellungen.

## Implementierungs-Überlegungen

- Dadurch dass Daten und Struktur des Graphen im selben Struct enthalten sind,
  können Änderungen am Implementierungskonzept umständlicher sein und der Code
  könnte schwerer zu verstehen sein, weil mehrere Aspekte vermischt werden.
- Bei einfachen Graphen/Anforderungen kann die Inline-Darstellung aber durchaus eine
  gute Wahl sein, weil sie direkt und einfach ist und wenig Boilerplate-Code erfordert,
  um z.B. nach Such-Operationen and die Daten der Knoten zu kommen.
