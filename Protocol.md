# CinemaService Pwn2Own

## Kommunikation zwischen den Services

### Grundlagen
Jeder Service kommuniziert über Proto Nachrichten mit den anderen Services.

Beispielhafter Aufbau der Proto Files
![](https://github.com/ob-vss-ws19/blatt-4-pwn2own/blob/develop/images/Proto.PNG)

Beispielhaftes Kreieren eines Services
![](https://github.com/ob-vss-ws19/blatt-4-pwn2own/blob/develop/images/MovieService.PNG)

Beispielhaftes Ausführen einer Funktion eines Services
![](https://github.com/ob-vss-ws19/blatt-4-pwn2own/blob/develop/images/MovieCreate.PNG)

### Abhängigkeiten
Damit die Daten der Services konsistent sind, existieren folgende Abhängigkeiten:

#### User und Reservation
- **User können nur gelöscht werden, wenn Sie keine Reservierungen haben**

#### Cinemahall und Reservation
- **Wird ein Kino gelöscht werden auch alle dazugehörigen Reservierungen automatisch gelöscht**

#### Cinemahall und Show
- **Wird ein Kino gelöscht werden auch alle dazugehörigen Veranstaltungen automatisch gelöscht**

Einbau einer Abhängigkeit in die main.go
![](https://github.com/ob-vss-ws19/blatt-4-pwn2own/blob/develop/images/Depend.PNG)

Einbau der Abhängigkeit in die jeweilige Funktion.
Die Services müssen hier miteinander kommunizieren. Daher ist es zwingend notwendig, dass beim Aufruf der Funktion beide Services laufen.
![](https://github.com/ob-vss-ws19/blatt-4-pwn2own/blob/develop/images/DependUser.PNG)
