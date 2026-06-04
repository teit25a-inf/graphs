# Graph-Darstellungen

Dieses Verzeichnis enthält verschiedene Beispiele, wie ein Graph im Computer
repräsentiert werden kann. Jede der Varianten hat unterschiedliche Vor- und Nachteile.

- [Kantenliste](kantenliste/README.md):
  - Einfach Repräsentation, aber ineffizient für die meisten Operationen.
  - Sehr nahe an der mathematischen Definition eines Graphen.

- [Adjazenzmatrix](adjazenzmatrix/README.md)
  - Effizient für die meisten Operationen, aber speicherintensiv.
  - Gut geeignet für dichte Graphen.

- [Adjazenzliste](adjazenzliste/README.md)
  - Effizient für die meisten Operationen und speichereffizient.
  - Gut geeignet für dünne Graphen.

- [Inline-Darstellung](inline/README.md)
  - Direkte Repräsentation der Nachbarn durch Zeiger.
  - Effizient für die meisten Operationen, aber weniger flexibel in der Implementierung.
