# Adjazenzliste

Die Adjazenzliste ist eine Sammlung von Listen, wobei jede Liste die Nachbarn eines
Knotens enthält. D.h. technisch gesehen ist es ebenfalls eine zweidimensionale Liste,
allerdings enthält die "Zeile" für jeden Knoten nur tatsächlich vorhandenen Nachbarn.

## Vorteile

- Speichereffizient, weil nur tatsächlich vorhandene Kanten gespeichert werden.
- Effizient für die meisten Operationen, z.B. das Durchlaufen der Nachbarn eines Knotens.

## Nachteile

- Die gezielte Prüfung, ob ein Knoten mit einem anderen verbunden ist, kann ineffizient
  sein, da die Liste der Nachbarn durchsucht werden muss.
  Dies kommt v.A. zum Tragen, falls viele Knoten viele Nachbarn haben, d.h.
  wenn der Graph relativ dicht ist. In diesem Fall könnte die Adjazenzmatrix eine bessere
  Wahl sein.

## Implementierungs-Überlegungen

- Der einfachste Fall ist, die Knoten durch Indizes zu repräsentieren, z.B. 0, 1, 2, ...
- In diesem Fall kann die Adjazenzliste einfach ein Array von Listen sein,
  wobei die i-te Liste die Nachbarn des Knotens i enthält.
- Wenn die eigentlichen Knoten Daten enthalten (was i.d.R. der Fall ist), dann könnte
  der Graph zusätzlich zur Adjazenzliste eine separate Liste oder Map von Knoten-IDs
  zu Knotendaten enthalten.
- Alternativ könnte man statt einer Liste von Nachbar-IDs auch eine Liste von
  Knoten-Objekten oder von Zeigern auf Knoten verwenden.
- Auch die Liste selbst muss keine Liste sein, sondern könnte z.B. eine Map von
  Knoten-Objekten (oder -Pointern) auf die Listen ihrer Nachbarn sein.
- In der Implementierung in diesem Package wird die Adjazenzliste als
  Map von Knoten-IDs auf Listen von Nachbar-IDs implementiert.
  D.h. es wird der obige Optimierungs-Ansatz umgesetzt, um keine leeren Nachbar-Listen
  speichern zu müssen. Es werden aber keine Knoten-Objekte oder -Pointer verwendet.
