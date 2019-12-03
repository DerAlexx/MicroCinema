package show_test

import (
	"testing"

	showtestproto "github.com/ob-vss-ws19/blatt-4-pwn2own/show/proto"
	"github.com/ob-vss-ws19/blatt-4-pwn2own/show/show"
)

/*
TestCreateShow will be a testcase for adding show into the service.
*/
func TestCreateShow(t *testing.T) {
	service := show.NewShowPool()
	response := showtestproto.CreateShowResponse{}
	service.CreateShow(nil, &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 2}}, &response)

	if response.CreateShowId < 0 {
		t.Errorf("Cannot create a show with the id %d", response.CreateShowId)
	} else {
		t.Log("Creating a Show will work.")
	}
}

/*
TestDeleteShow will be a testcase for deleting a show from the service.
*/
func TestDeleteShow(t *testing.T) {
	service := show.NewShowPool()
	response := showtestproto.CreateShowResponse{}
	service.CreateShow(nil, &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 2}}, &response)
	responseDelete := showtestproto.DeleteShowResponse{}
	service.DeleteShow(nil, &showtestproto.DeleteShowRequest{DeleteShowId: response.CreateShowId}, &responseDelete)

	if !responseDelete.Answer {
		t.Errorf("Cannot delete the show with the id %d", response.CreateShowId)
	} else {
		t.Log("Deleting a Show will work.")
	}
}

/*
TestDeleteShowConnectedCinema will be a testcase for deleting all shows with a specific cinemaId.
*/
func TestDeleteShowConnectedCinema(t *testing.T) {
	service := show.NewShowPool()
	service.CreateShow(nil, &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 1}}, &showtestproto.CreateShowResponse{})
	service.CreateShow(nil, &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 2}}, &showtestproto.CreateShowResponse{})
	service.CreateShow(nil, &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 3}}, &showtestproto.CreateShowResponse{})
	service.CreateShow(nil, &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 2, MovieId: 2}}, &showtestproto.CreateShowResponse{})
	responseDeleteCinema := showtestproto.DeleteShowConnectedCinemaResponse{}
	service.DeleteShowConnectedCinema(nil, &showtestproto.DeleteShowConnectedCinemaRequest{CinemaId: 1}, &responseDeleteCinema)

	if !responseDeleteCinema.Answer {
		t.Errorf("Cannot delete shows with the cinemaid %d", 1)
	} else {
		t.Log("Deleting a show with the cinemaid will work.")
	}
}

/*
TestDeleteShowConnectedMove will be a testcase for deleting will be a testcase for deleting all shows with a specific movieId.
*/
func TestDeleteShowConnectedMovie(t *testing.T) {
	service := show.NewShowPool()
	service.CreateShow(nil, &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 1}}, &showtestproto.CreateShowResponse{})
	service.CreateShow(nil, &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 2}}, &showtestproto.CreateShowResponse{})
	service.CreateShow(nil, &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 3}}, &showtestproto.CreateShowResponse{})
	service.CreateShow(nil, &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 2, MovieId: 2}}, &showtestproto.CreateShowResponse{})
	responseDeleteMovie := showtestproto.DeleteShowConnectedMovieResponse{}
	service.DeleteShowConnectedMovie(nil, &showtestproto.DeleteShowConnectedMovieRequest{MovieId: 2}, &responseDeleteMovie)

	if !responseDeleteMovie.Answer {
		t.Errorf("Cannot delete shows with the moveid %d", 2)
	} else {
		t.Log("Deleting a show with the movieid will work.")
	}
}

/*
TestListShow will be a testcase for listing all shows.
*/
func TestListShow(t *testing.T) {
	service := show.NewShowPool()
	service.CreateShow(nil, &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 1}}, &showtestproto.CreateShowResponse{})
	service.CreateShow(nil, &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 2}}, &showtestproto.CreateShowResponse{})
	service.CreateShow(nil, &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 3}}, &showtestproto.CreateShowResponse{})
	service.CreateShow(nil, &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 2, MovieId: 2}}, &showtestproto.CreateShowResponse{})
	service.DeleteShowConnectedMovie(nil, &showtestproto.DeleteShowConnectedMovieRequest{MovieId: 2}, &showtestproto.DeleteShowConnectedMovieResponse{})
	responseList := showtestproto.ListShowResponse{}
	service.ListShow(nil, &showtestproto.ListShowRequest{}, &responseList)
	if len(responseList.AllShowsData) != 2 {
		t.Errorf("Cannot list all shows")
	} else {
		t.Log("listing all will work.")
	}
}

/*
TestFindShowConnectedCinema will be a testcase for finding all shows with a specific cinemaId.
*/
func TestFindShowConnectedCinema(t *testing.T) {
	service := show.NewShowPool()
	service.CreateShow(nil, &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 1}}, &showtestproto.CreateShowResponse{})
	service.CreateShow(nil, &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 2}}, &showtestproto.CreateShowResponse{})
	service.CreateShow(nil, &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 3}}, &showtestproto.CreateShowResponse{})
	service.CreateShow(nil, &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 2, MovieId: 2}}, &showtestproto.CreateShowResponse{})
	responseFindCinema := showtestproto.FindShowConnectedCinemaResponse{}
	service.FindShowConnectedCinema(nil, &showtestproto.FindShowConnectedCinemaRequest{CinemaId: 1}, &responseFindCinema)
	if len(responseFindCinema.CinemaData) != 3 {
		t.Errorf("Cannot find all shows with cinemaid: %d", 1)
	} else {
		t.Log("Finding all shows with a cinemaid will work.")
	}
}

/*
TestFindShowConnectedMovie will be a testcase for finding all shows with a specific movieId.
*/
func TestFindShowConnectedMovie(t *testing.T) {
	service := show.NewShowPool()
	service.CreateShow(nil, &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 1}}, &showtestproto.CreateShowResponse{})
	service.CreateShow(nil, &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 2}}, &showtestproto.CreateShowResponse{})
	service.CreateShow(nil, &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 3}}, &showtestproto.CreateShowResponse{})
	service.CreateShow(nil, &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 2, MovieId: 2}}, &showtestproto.CreateShowResponse{})
	responseFindMovie := showtestproto.FindShowConnectedMovieResponse{}
	service.FindShowConnectedMovie(nil, &showtestproto.FindShowConnectedMovieRequest{MovieId: 2}, &responseFindMovie)
	if len(responseFindMovie.MovieData) != 2 {
		t.Errorf("Cannot find all shows with movieid: %d", 2)
	} else {
		t.Log("Finding all shows with a movieid will work.")
	}
}
