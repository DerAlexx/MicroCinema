package show

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	showproto "github.com/ob-vss-ws19/blatt-4-pwn2own/show/proto"
)

const (
	maxshowid = 432342
)

/*
ShowPool contains all cinemas.
*/
type ShowPool struct {
	showmap map[int32]*show
	mutex   *sync.Mutex
}

type show struct {
	cinemaID int
	movieID  int
}

/*
NewShowPool creates a new CinemaPool
*/
func NewShowPool() *ShowPool {
	newshowmap := make(map[int32]*show)

	return &ShowPool{
		mutex:   &sync.Mutex{},
		showmap: newshowmap,
	}
}

func (handler *ShowPool) containsshow(id int32) bool {
	_, containsinmap := handler.showmap[id]
	return containsinmap
}

func (handler *ShowPool) getRandomShowID() int32 {
	rand.Seed(time.Now().UnixNano())
	for {
		potentialID := int32(rand.Intn(maxshowid))
		if !handler.containsshow(potentialID) {
			return potentialID
		}
	}
}

/*
CreateShow will create a show, combining a cinema with a movie.
*/
func (handler *ShowPool) CreateShow(ctx context.Context, request *showproto.CreateShowRequest, response *showproto.CreateShowResponse) error {
	if request.CreateData.CinemaId > 0 && request.CreateData.MovieId > 0 {
		createid := handler.getRandomShowID()
		handler.mutex.Lock()
		handler.showmap[createid] = &show{
			cinemaID: int(request.CreateData.CinemaId),
			movieID:  int(request.CreateData.MovieId),
		}
		handler.mutex.Unlock()
		fmt.Println(string(createid))
		response.CreateShowId = createid
		return nil
	}
	return fmt.Errorf("cannot create show with cinemaID: %d and movieId: %d", request.CreateData.CinemaId, request.CreateData.MovieId)
}

/*
DeleteShow will delete one show(id) from the cinemapool.
*/
func (handler *ShowPool) DeleteShow(ctx context.Context, request *showproto.DeleteShowRequest, response *showproto.DeleteShowResponse) error {
	if request.DeleteShowId > 0 && handler.containsshow(request.DeleteShowId) {
		handler.mutex.Lock()
		delete(handler.showmap, request.DeleteShowId)
		handler.mutex.Unlock()
		response.Answer = true
		return nil
	}
	return fmt.Errorf("cannot delete show with id: %d", request.DeleteShowId)
}

/*
DeleteShowConnectedCinema will delete all shows connected to a cinema.
*/
func (handler *ShowPool) DeleteShowConnectedCinema(ctx context.Context, request *showproto.DeleteShowConnectedCinemaRequest, response *showproto.DeleteShowConnectedCinemaResponse) error {
	if request.CinemaId > 0 {
		handler.mutex.Lock()
		success := false
		for deleteID, deleteShow := range handler.showmap {
			if deleteShow.cinemaID == int(request.CinemaId) {
				delete(handler.showmap, deleteID)
				success = true
			}
		}
		handler.mutex.Unlock()
		if success {
			response.Answer = true
			return nil
		}
	}
	return fmt.Errorf("cannot delete show with id: %d", request.CinemaId)
}

/*
DeleteShowConnectedMovie will delete all shows connected to a movie.
*/
func (handler *ShowPool) DeleteShowConnectedMovie(ctx context.Context, request *showproto.DeleteShowConnectedMovieRequest, response *showproto.DeleteShowConnectedMovieResponse) error {
	if request.MovieId > 0 {
		handler.mutex.Lock()
		success := false
		for deleteID, deleteShow := range handler.showmap {
			if deleteShow.movieID == int(request.MovieId) {
				delete(handler.showmap, deleteID)
				success = true
			}
		}
		handler.mutex.Unlock()
		if success {
			response.Answer = true
			return nil
		}
	}
	return fmt.Errorf("cannot delete show with id: %d", request.MovieId)
}

/*
ListShow will show all shows currently available in cinemapool.
*/
func (handler *ShowPool) ListShow(ctx context.Context, request *showproto.ListShowRequest, response *showproto.ListShowResponse) error {
	responseID := []int32{}
	responseData := []*showproto.ShowMessage{}
	handler.mutex.Lock()
	for listShowID, listShow := range handler.showmap {
		responseID = append(responseID, listShowID)
		responseData = append(responseData, &showproto.ShowMessage{CinemaId: int32(listShow.cinemaID), MovieId: int32(listShow.movieID)})
	}
	handler.mutex.Unlock()
	response.ShowId = responseID
	response.AllShowsData = responseData
	return nil
}

/*
FindShowConnectedCinema will show all shows connected to a cinema.
*/
func (handler *ShowPool) FindShowConnectedCinema(ctx context.Context, request *showproto.FindShowConnectedCinemaRequest, response *showproto.FindShowConnectedCinemaResponse) error {
	if request.CinemaId > 0 {
		ids := []int32{}
		responseData := []*showproto.ShowMessage{}
		handler.mutex.Lock()
		for showid, findCinShow := range handler.showmap {
			if findCinShow.cinemaID == int(request.CinemaId) {
				ids = append(ids, showid)
				responseData = append(responseData, &showproto.ShowMessage{CinemaId: int32(findCinShow.cinemaID), MovieId: int32(findCinShow.movieID)})
			}
		}
		handler.mutex.Unlock()
		response.Ids = ids
		response.CinemaData = responseData
		return nil
	}
	return fmt.Errorf("cannot find show with cinemaID: %d", request.CinemaId)
}

/*
FindShowConnectedMovie will show all shows connected to a movie. ++
*/
func (handler *ShowPool) FindShowConnectedMovie(ctx context.Context, request *showproto.FindShowConnectedMovieRequest, response *showproto.FindShowConnectedMovieResponse) error {
	if request.MovieId > 0 {
		responseData := []*showproto.ShowMessage{}
		handler.mutex.Lock()
		for _, findMovShow := range handler.showmap {
			if findMovShow.movieID == int(request.MovieId) {
				responseData = append(responseData, &showproto.ShowMessage{CinemaId: int32(findMovShow.cinemaID), MovieId: int32(findMovShow.movieID)})
			}
		}
		handler.mutex.Unlock()
		response.MovieData = responseData
		return nil
	}
	return fmt.Errorf("cannot find show with movieid: %d", request.MovieId)
}
