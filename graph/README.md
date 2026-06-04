# Graph-Interface

Dieses Modul definiert ein Interface `GraphRepr` für Graphen,
das von den verschiedenen Darstellungen implementiert werden muss.

## Interface

Das Interface definiert die grundlegenden Operationen,
die auf einem Graphen ausgeführt werden können.
Es ist eine Liste von Methoden, die von den konkreten Implementierungen
bereitgestellt werden müssen.
Die Structs in den Packages in `graph_darstellungen` implementieren dieses Interface.

## Funktionen, die das Interface vewenden

In der Datei [`graph_functions.go`](./graph_functions.go) befinden sich Funktionen,
die das Interface verwenden.
Diese Funktionen können mit jeder Implementierung des Interfaces verwendet werden,
was die Flexibilität und Wiederverwendbarkeit des Codes erhöht.

Um die Aufgaben in `graph_functions.go` zu lösen, müssen Sie zunächst eine der
Implementierungen des Interfaces fertigstellen, damit Sie die Funktionen testen können.
Die Tests für die Funktionen in `graph_functions.go` befinden sich mit bei den Tests
der Implementierungen des Interfaces, da sie von diesen abhängig sind.

## Tests für `GraphRepr`-Instanzen

In den Dateien `test_functions.go` und `test_helpers.go` befinden sich Hilfsfunktionen
für Instanzen des `GraphRepr`-Interfaces, die in den Tests der Implementierungen
verwendet werden.
Die eigentlichen Tests sind bei den Implementierungen, rufen aber die Funktionen aus
`test_functions.go` auf, um die Funktionalität der Implementierungen zu überprüfen.
