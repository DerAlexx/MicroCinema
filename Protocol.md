# CinemaService Pwn2Own

## Kommunikation zwischen den Services

### Grundlagen
Jeder Service kommuniziert über Proto Nachrichten mit den anderen Services.

![](https://github.com/ob-vss-ws19/blatt-4-pwn2own/blob/develop/images/MovieService.PNG)

![](https://github.com/ob-vss-ws19/blatt-4-pwn2own/blob/develop/images/MovieCreate.PNG)

![](https://github.com/ob-vss-ws19/blatt-4-pwn2own/blob/develop/images/Proto.PNG)



### Abhängigkeiten
Damit die Daten der Services konsistent sind, existieren folgende Abhängigkeiten:

#### User und Reservation
- **User können nur gelöscht werden, wenn Sie keine Reservierungen haben**

#### Cinemahall und Reservation
- **Wird ein Kino gelöscht werden auch alle dazugehörigen Reservierungen automatisch gelöscht**

#### Cinemahall und Show
- **Wird ein Kino gelöscht werden auch alle dazugehörigen Veranstaltungen automatisch gelöscht**

![](https://github.com/ob-vss-ws19/blatt-4-pwn2own/blob/develop/images/Depend.PNG)

![](https://github.com/ob-vss-ws19/blatt-4-pwn2own/blob/develop/images/DependUser.PNG)
