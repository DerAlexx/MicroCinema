package movies

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	proto "github.com/ob-vss-ws19/blatt-4-pwn2own/movies/proto"
)

const (
	maxmoviesid int32 = 98765
)

/*
Movie will represent a movie.
*/
type Movie struct {
	name string
}

/*
MovieHandlerService will be the representation of our service.
*/
type MovieHandlerService struct {
	movies       map[int32]*Movie
	dependencies []interface{}
	mutex        *sync.Mutex
}

/*
containsID will check whether generated ID is already set or not.
@id will be a int32 id to check for.
*/
func (m *MovieHandlerService) containsID(id int32) bool {
	_, inMap := (*m.getMoviesMap())[id]
	return inMap
}

/*
Function to produce a random moviesID.
@param length will be the length of the number.
@param seed will be a seed in order to produce "really" random numbers.
*/
func (m *MovieHandlerService) getRandomUserID(length int32) int32 {
	rand.Seed(time.Now().UnixNano())
	for {
		potantialID := rand.Int31n(length)
		if !m.containsID(int32(potantialID)) {
			return potantialID
		}
	}
}

/*
getMoviesMap will return a pointer to the current movie map in order to work in that.
*/
func (m *MovieHandlerService) getMoviesMap() *map[int32]*Movie {
	return &m.movies
}

/*
setMoviesMap will set the map of a userhandlerservice instance.
@param movies will be the map to set.
*/
func (m *MovieHandlerService) setMoviesMap(movies *map[int32]*Movie) {
	m.movies = *movies
}

/*
CreateNewMoviesHandlerInstance will return a new movies instance.
*/
func CreateNewMoviesHandlerInstance() *MovieHandlerService {
	return &MovieHandlerService{
		movies: make(map[int32]*Movie),
		mutex:  &sync.Mutex{},
	}
}

/*
appendMovie will add a movie in the datastructure.
*/
func (m *MovieHandlerService) appendMovie(id int32, movie *Movie) bool {
	if id > 0 && movie != nil {
		(*m.getMoviesMap())[id] = movie
		return true
	}
	return false
}

/*
CreateMovie will create a movie.
*/
func (m *MovieHandlerService) CreateMovie(context context.Context, in *proto.CreateMovieRequest, out *proto.CreatedMovieResponse) error {
	if in.GetName() != "" {
		m.mutex.Lock()
		mid := m.getRandomUserID(maxmoviesid)
		if m.appendMovie(mid, &Movie{name: in.GetName()}) {
			out.Movie.Id = mid
			out.Movie.Name = in.GetName()
			m.mutex.Unlock()
			return nil
		}
	}
	return fmt.Errorf("cannot create movie with an emtpy name")
}

/*
change will change an entry in the "database".
*/
func (m *MovieHandlerService) change(id int32, pname string) bool {
	if m.containsID(id) {
		m.mutex.Lock()
		(*m.getMoviesMap())[id] = &Movie{name: pname}
		m.mutex.Unlock()
		return true
	}
	return false
}

/*
ChangeMovie will change a movie.
*/
func (m *MovieHandlerService) ChangeMovie(ctx context.Context, in *proto.ChangeMovieRequest, out *proto.ChangeMovieResponse) error {
	if in.Movie.Id > 0 && in.Movie.Name != "" {
		if m.change(in.Movie.Id, in.Movie.Name) {
			out.Changed = true
			return nil
		}
	}
	return fmt.Errorf("cannot change the movie. The movieid or the name are not ok. See: %d %s", in.Movie.Id, in.Movie.Name)
}
