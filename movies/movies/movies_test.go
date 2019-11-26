package movies_test

import (
	"testing"

	"github.com/ob-vss-ws19/blatt-4-pwn2own/movies/movies"
	protoo "github.com/ob-vss-ws19/blatt-4-pwn2own/movies/proto"
)

/*
TestAddMovie will be a testcase for adding users into the service.
*/
func TestAddMovie(t *testing.T) {
	TestName := "Bibi und Tina"
	service := movies.CreateNewMoviesHandlerInstance()
	response := protoo.CreatedMovieResponse{Movie: &protoo.Movie{}}
	service.CreateMovie(nil, &protoo.CreateMovieRequest{Name: TestName}, &response)

	if response.Movie.Name != TestName {
		t.Errorf("Cannot create a movie with the name %s", TestName)
	} else if response.Movie.Id < 1 {
		t.Fatal("Cannot create a movie with a proper ID")
	} else {
		t.Log("Creating a User will work.")
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
	service.CreateMovie(nil, &protoo.CreateMovieRequest{Name: FirstName}, &responseInsert)
	service.CreateMovie(nil, &protoo.CreateMovieRequest{Name: SecondName}, &responseInsert2)

	all := protoo.StreamMovieResponse{}

	service.StreamMovie(nil, &protoo.StreamMovieRequest{}, &all)

	if len(all.Movies) != 2 {
		t.Errorf("The length does not match up. expected %d got %d", 2, len(all.Movies))
	}
}

/*
TestAddandDeleteAMovie will create a movie and later on delete him.
*/
func TestAddandDeleteAMovie(t *testing.T) {
	TestName := "Tim"
	service := movies.CreateNewMoviesHandlerInstance()
	response := protoo.CreatedMovieResponse{Movie: &protoo.Movie{}}
	service.CreateMovie(nil, &protoo.CreateMovieRequest{Name: TestName}, &response)
	id := response.Movie.Id
	deleteResponse := protoo.DeleteMovieResponse{}
	service.DeleteMovie(nil, &protoo.DeleteMovieRequest{Id: id}, &deleteResponse)

	responseFind := protoo.FindMovieResponse{Movie: &protoo.Movie{}}

	service.FindMovie(nil, &protoo.FindMovieRequest{Movie: &protoo.Movie{}}, &responseFind)

	if !deleteResponse.Deleted && responseFind.Movie.Id == -1 {
		t.Errorf("Movie was deleted. But returned the false state.")
	} else {
		t.Log("Create a user and change him later on is fine.")
	}
}
