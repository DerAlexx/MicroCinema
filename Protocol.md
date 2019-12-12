# CinemaService Pwn2Own

## Kommunikation zwischen den Services

### Grundlagen
Jeder Service kommuniziert über Proto Nachrichten mit den anderen Services.

### Abhängigkeiten
Damit die Daten der Services konsistent sind, existieren folgende Abhängigkeiten:

#### User und Reservation
- **User können nur gelöscht werden, wenn Sie keine Reservierungen haben**

#### Cinemahall und Reservation
- **Wird ein Kino gelöscht werden auch alle dazugehörigen Reservierungen automatisch gelöscht**

#### Cinemahall und Show
- **Wird ein Kino gelöscht werden auch alle dazugehörigen Veranstaltungen automatisch gelöscht**
