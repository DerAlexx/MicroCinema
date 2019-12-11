package movies_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/ob-vss-ws19/blatt-4-pwn2own/movies/movies"
	protoo "github.com/ob-vss-ws19/blatt-4-pwn2own/movies/proto"
)

/*
TestAddMovie will be a testcase for adding movies into the service.
*/
func TestAddMovie(t *testing.T) {
	TestName := "Bibi und Tina"
	service := movies.CreateNewMoviesHandlerInstance()
	response := protoo.CreatedMovieResponse{Movie: &protoo.Movie{}}
	err := service.CreateMovie(context.TODO(), &protoo.CreateMovieRequest{Name: TestName}, &response)

	if err == nil {
		switch {
		case response.Movie.Name != TestName:
			t.Errorf("Cannot create a movie with the name %s", TestName)
		case response.Movie.Id < 1:
			t.Fatal("Cannot create a movie with a proper ID")
		default:
			t.Log("Creating a movie will work.")
		}
	} else {
		fmt.Println(err)
	}
}

/*
TestAddMultipleMoviesAndReadAllOfThem will create bunch of movies and read them later on all from the service.
*/
func TestAddMultipleMoviesAndReadAllOfThem(t *testing.T) {
	FirstName := "Jim Knopf"
	SecondName := "Bibi Blocksberg"
	service := movies.CreateNewMoviesHandlerInstance()
	responseInsert := protoo.CreatedMovieResponse{Movie: &protoo.Movie{}}
	responseInsert2 := protoo.CreatedMovieResponse{Movie: &protoo.Movie{}}
	err := service.CreateMovie(context.TODO(), &protoo.CreateMovieRequest{Name: FirstName}, &responseInsert)
	err1 := service.CreateMovie(context.TODO(), &protoo.CreateMovieRequest{Name: SecondName}, &responseInsert2)

	all := protoo.StreamMovieResponse{}

	err2 := service.StreamMovie(context.TODO(), &protoo.StreamMovieRequest{}, &all)

	if err == nil && err1 == nil && err2 == nil {
		if len(all.Movies) != 2 {
			t.Errorf("The length does not match up. expected %d got %d", 2, len(all.Movies))
		}
	} else {
		fmt.Println(err)
		fmt.Println(err1)
		fmt.Println(err2)
	}
}

/*
TestAddandDeleteAMovie will create a movie and later on delete him.
*/
func TestAddandDeleteAMovie(t *testing.T) {
	TestName := "Tim"
	service := movies.CreateNewMoviesHandlerInstance()
	response := protoo.CreatedMovieResponse{Movie: &protoo.Movie{}}
	err := service.CreateMovie(context.TODO(), &protoo.CreateMovieRequest{Name: TestName}, &response)
	id := response.Movie.Id
	deleteResponse := protoo.DeleteMovieResponse{}
	err1 := service.DeleteMovie(context.TODO(), &protoo.DeleteMovieRequest{Id: id}, &deleteResponse)

	responseFind := protoo.FindMovieResponse{Movie: &protoo.Movie{}}

	err2 := service.FindMovie(context.TODO(), &protoo.FindMovieRequest{Movie: &protoo.Movie{}}, &responseFind)

	if err == nil && err1 == nil && err2 == nil {
		if !deleteResponse.Deleted && responseFind.Movie.Id == -1 {
			t.Errorf("Movie was deleted. But returned the false state.")
		} else {
			t.Log("Create a movie and change him later on is fine.")
		}
	} else {
		fmt.Println(err)
		fmt.Println(err1)
		fmt.Println(err2)
	}
}

/*
TestAddMovieAndFindIt will be a testcase for adding movies into the service and try to find him by his name.
*/
func TestAddMovieAndFindIt(t *testing.T) {
	TestName := "Mission Impossible 6"
	service := movies.CreateNewMoviesHandlerInstance()
	response := protoo.CreatedMovieResponse{Movie: &protoo.Movie{}}
	err := service.CreateMovie(context.TODO(), &protoo.CreateMovieRequest{Name: TestName}, &response)
	id := response.Movie.Id

	responseFind := protoo.FindMovieResponse{Movie: &protoo.Movie{}}

	err1 := service.FindMovie(context.TODO(), &protoo.FindMovieRequest{Movie: &protoo.Movie{Name: TestName}}, &responseFind)

	if err == nil && err1 == nil {
		switch {
		case responseFind.Movie.Id != id:
			t.Errorf("Cannot find or create a movie with the name %s --> ID does not match given %d want %d", TestName, responseFind.Movie.Id, id)
		case responseFind.Movie.Id < 1 || responseFind.Movie.Id != id:
			t.Errorf("Cannot find a movie with given ID --> Does not match up given with expected ID given %d, wanted %d", id, responseFind.Movie.Id)
		case responseFind.Movie.Name == "" || responseFind.Movie.Name != TestName:
			t.Errorf("Cannot find a movie with given Name --> Missing match given %s, wanted %s", responseFind.Movie.Name, TestName)
		default:
			t.Log("Can create a movie and get him by his name is fine.")
		}
	} else {
		fmt.Println(err)
		fmt.Println(err1)
	}
}

/*
TestAddChange will create a user change him an later on call getinformationfrommap in order to see whether or not something
has changed.
*/
func TestAddChange(t *testing.T) {
	TestName := "Die Abenteuer des Paulanius"
	NewName := "Die Abenteuer des Paulanius 2"
	service := movies.CreateNewMoviesHandlerInstance()
	response := protoo.CreatedMovieResponse{Movie: &protoo.Movie{}}
	err := service.CreateMovie(context.TODO(), &protoo.CreateMovieRequest{Name: TestName}, &response)
	id := response.Movie.Id
	chresponse := protoo.ChangeMovieResponse{}
	beforeChange := service.Find(response.Movie.Id).(string)
	err1 := service.ChangeMovie(context.TODO(), &protoo.ChangeMovieRequest{Movie: &protoo.Movie{Id: id, Name: NewName}}, &chresponse)
	AfterChange := service.Find(response.Movie.Id).(string)

	if err == nil && err1 == nil {
		switch {
		case beforeChange != TestName:
			t.Errorf("Beforename is wrong got: %s wanted %s", beforeChange, TestName)
		case id <= 0:
			t.Errorf("Got a wrong movieid %d wanted a value bigger than 0", id)
		case AfterChange != NewName:
			t.Errorf("Aftername is wrong got: %s wanted %s", AfterChange, NewName)
		case AfterChange == NewName && !chresponse.Changed:
			t.Errorf("Name was changed. Found by getinformationfrommap but did not send the correct response.")
		default:
			t.Log("Create a movie and change it later on is fine.")
		}
	} else {
		fmt.Println(err)
		fmt.Println(err1)
	}
}
