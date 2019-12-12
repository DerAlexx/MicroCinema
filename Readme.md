# CinemaService Pwn2Own

## Installation
- **Es wird die GO Version 1.12 benötigt**

## Ausführen der Tests
In den Unterverzeichnissen der einzelnen Services befindet sich jeweils eine Testfile, die Junit Tests ausführt.
Dabei werden alle Funktionen getestet die unabhängig von anderen Services sind.

cd cinemahall || go test -cover
cd movies || go test -cover
cd reservation || go test -cover
cd show || go test -cover
cd users || go test -cover

## Ausführen der Services
In den Unterverzeichnissen der einzelnen Services befindet sich jeweils eine ausführbare Datei, die den Service startet.

go run cinemahall/main.go
go run movies/main.go
go run reservation/main.go
go run show/main.go
go run users/main.go

Über go run client.go kann ein exemplarisches Testprogramm gestartet werden, dass die grundlegenden Funktionen der Services testet.

