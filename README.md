# Übungsaufgaben zu Graphen

In diesem Ordner befinden sich die Aufgaben zu Graphen.

## Überblick

Im Verzeichnis `graph` ist ein Interface `GraphRepr` definiert,
das in allgemeiner Form die grundlegenden Operationen für Graphen beschreibt.

Im Verzeichnis `graph_darstellungen` befinden sich mehrere Module,
die verschiedene Darstellungen von Graphen implementieren.
Jede dieser Darstellungen is funktional gleichwertig und implementiert das Interface
`GraphRepr`.
D.h. für die generelle Graph-Funktionalität spielt es keine Rolle, welche Darstellung
verwendet wird und die Module sind voneinander unabhängig.
Die Wahl der Darstellung ist nur relevant für die Implementierung von Algorithmen,
da manche Darstellungen für bestimmte Algorithmen effizienter sein können als andere.

## Aufgaben

- Implementieren Sie mindestens eine der Darstellungen von Graphen in `graph_darstellungen`.
  - Eine Implementierung genügt, um die folgende Aufgabe zu lösen.
  - Sie sollten dennoch (gleich oder hinterher) mehrere Darstellungen implementieren,
    um die Unterschiede zwischen den Darstellungen zu verstehen.
- Implementieren Sie die Funktionen in `graph_functions.go` im Verzeichnis `graph`.
  - Diese Funktionen verwenden das Interface `GraphRepr` und können daher mit jeder
    Implementierung des Interfaces verwendet werden.
  - Die zugehörigen Tests befinden sich bei den Implementierungen der Darstellungen,
    verwenden aber die Hilfsfunktionen aus `test_functions.go` im Verzeichnis `graph`.
- Im Modul `graphviz_main` befindet sich ein Beispiel, wie die Funktion `DotString`
  aus `graph_functions.go` verwendet werden kann.
  Dies ist ein `main-Programm, das einen Graphen erstellt, ihn in das DOT-Format umwandelt
  und die Ausgabe anzeigt.
  Verwenden Sie dieses Modul, um Ihre Implementierung zu testen und die Ausgabe zu überprüfen.
