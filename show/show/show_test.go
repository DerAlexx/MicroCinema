package show_test

import (
	"context"
	"fmt"
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
	err := service.CreateShow(context.TODO(), &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 2}}, &response)

	if err == nil {
		if response.CreateShowId < 0 {
			t.Errorf("Cannot create a show with the id %d", response.CreateShowId)
		} else {
			t.Log("Creating a Show will work.")
		}
	} else {
		fmt.Println(err)
	}
}

/*
TestDeleteShow will be a testcase for deleting a show from the service.
*/
func TestDeleteShow(t *testing.T) {
	service := show.NewShowPool()
	response := showtestproto.CreateShowResponse{}
	err := service.CreateShow(context.TODO(), &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 2}}, &response)
	responseDelete := showtestproto.DeleteShowResponse{}
	err1 := service.DeleteShow(context.TODO(), &showtestproto.DeleteShowRequest{DeleteShowId: response.CreateShowId}, &responseDelete)
	if err == nil && err1 == nil {
		if !responseDelete.Answer {
			t.Errorf("cannot delete the show with the id %d", response.CreateShowId)
		} else {
			t.Log("Deleting a Show will work.")
		}
	} else {
		fmt.Println(err)
		fmt.Println(err1)
	}
}

/*
TestDeleteShowConnectedMove will be a testcase for deleting will be a testcase for deleting all shows with a specific movieId and cinemaid.
*/
func TestDeleteShowConnectedMovie(t *testing.T) {
	service := show.NewShowPool()
	err := service.CreateShow(context.TODO(), &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 1}}, &showtestproto.CreateShowResponse{})
	err1 := service.CreateShow(context.TODO(), &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 2}}, &showtestproto.CreateShowResponse{})
	err2 := service.CreateShow(context.TODO(), &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 3}}, &showtestproto.CreateShowResponse{})
	err3 := service.CreateShow(context.TODO(), &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 2, MovieId: 2}}, &showtestproto.CreateShowResponse{})
	responseDeleteMovie := showtestproto.DeleteShowConnectedMovieResponse{}
	err4 := service.DeleteShowConnectedMovie(context.TODO(), &showtestproto.DeleteShowConnectedMovieRequest{MovieId: 2}, &responseDeleteMovie)
	responseDeleteCinema := showtestproto.DeleteShowConnectedCinemaResponse{}
	err5 := service.DeleteShowConnectedCinema(context.TODO(), &showtestproto.DeleteShowConnectedCinemaRequest{CinemaId: 1}, &responseDeleteCinema)

	if err == nil && err1 == nil && err2 == nil && err3 == nil && err4 == nil && err5 == nil {
		switch {
		case !responseDeleteMovie.Answer:
			t.Errorf("cannot delete shows with the moveid %d", 2)
		case !responseDeleteCinema.Answer:
			t.Errorf("cannot delete shows with the cinemaid %d", 1)
		default:
			t.Log("Deleting a show with the cinemaid and movieid will work.")
		}
	} else {
		fmt.Println(err)
		fmt.Println(err1)
		fmt.Println(err2)
		fmt.Println(err3)
		fmt.Println(err4)
		fmt.Println(err5)
	}
}

/*
TestListShow will be a testcase for listing all shows.
*/
func TestListShow(t *testing.T) {
	service := show.NewShowPool()
	err := service.CreateShow(context.TODO(), &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 1}}, &showtestproto.CreateShowResponse{})
	err1 := service.CreateShow(context.TODO(), &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 2}}, &showtestproto.CreateShowResponse{})
	err2 := service.CreateShow(context.TODO(), &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 3}}, &showtestproto.CreateShowResponse{})
	err3 := service.CreateShow(context.TODO(), &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 2, MovieId: 2}}, &showtestproto.CreateShowResponse{})
	err4 := service.DeleteShowConnectedMovie(context.TODO(), &showtestproto.DeleteShowConnectedMovieRequest{MovieId: 2}, &showtestproto.DeleteShowConnectedMovieResponse{})
	responseList := showtestproto.ListShowResponse{}
	err5 := service.ListShow(context.TODO(), &showtestproto.ListShowRequest{}, &responseList)
	if err == nil && err1 == nil && err2 == nil && err3 == nil && err4 == nil && err5 == nil {
		if len(responseList.AllShowsData) != 2 {
			t.Errorf("cannot list all shows")
		} else {
			t.Log("listing all will work.")
		}
	} else {
		fmt.Println(err)
		fmt.Println(err1)
		fmt.Println(err2)
		fmt.Println(err3)
		fmt.Println(err4)
		fmt.Println(err5)
	}
}

/*
TestFindShowConnectedMovie will be a testcase for finding all shows with a specific movieId and cinemaid.
*/
func TestFindShowConnectedMovie(t *testing.T) {
	service := show.NewShowPool()
	err := service.CreateShow(context.TODO(), &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 1}}, &showtestproto.CreateShowResponse{})
	err1 := service.CreateShow(context.TODO(), &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 2}}, &showtestproto.CreateShowResponse{})
	err2 := service.CreateShow(context.TODO(), &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 1, MovieId: 3}}, &showtestproto.CreateShowResponse{})
	err3 := service.CreateShow(context.TODO(), &showtestproto.CreateShowRequest{CreateData: &showtestproto.ShowMessage{CinemaId: 2, MovieId: 2}}, &showtestproto.CreateShowResponse{})
	responseFindMovie := showtestproto.FindShowConnectedMovieResponse{}
	err4 := service.FindShowConnectedMovie(context.TODO(), &showtestproto.FindShowConnectedMovieRequest{MovieId: 2}, &responseFindMovie)
	responseFindCinema := showtestproto.FindShowConnectedCinemaResponse{}
	err5 := service.FindShowConnectedCinema(context.TODO(), &showtestproto.FindShowConnectedCinemaRequest{CinemaId: 1}, &responseFindCinema)

	if err == nil && err1 == nil && err2 == nil && err3 == nil && err4 == nil && err5 == nil {
		switch {
		case len(responseFindMovie.MovieData) != 2:
			t.Errorf("cannot find all shows with movieid: %d", 2)
		case len(responseFindCinema.CinemaData) != 3:
			t.Errorf("cannot find all shows with cinemaid: %d", 1)
		default:
			t.Log("Finding all shows with a movieid and cinemaid will work.")
		}
	} else {
		fmt.Println(err)
		fmt.Println(err1)
		fmt.Println(err2)
		fmt.Println(err3)
		fmt.Println(err4)
		fmt.Println(err5)
	}
}
