package show

import "sync"

/*
ShowPool contains all cinemas.
*/
type ShowPool struct {
	showmap map[int64]*show
	mutex   *sync.Mutex
}

type show struct {
	showId  int
	movieId int
}

/*
NewCinemaPool creates a new CinemaPool
*/
func NewShowPool() *ShowPool {
	newshowmap := make(map[int64]*show)

	return &ShowPool{
		mutex:   &sync.Mutex{},
		showmap: newshowmap,
	}
}
