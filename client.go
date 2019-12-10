package main

import (
	"context"
	"fmt"
	"strconv"

	micro "github.com/micro/go-micro"
	cinemaprot "github.com/ob-vss-ws19/blatt-4-pwn2own/cinemahall/proto"
	moviesprot "github.com/ob-vss-ws19/blatt-4-pwn2own/movies/proto"
	reservationprot "github.com/ob-vss-ws19/blatt-4-pwn2own/reservation/proto"
	showprot "github.com/ob-vss-ws19/blatt-4-pwn2own/show/proto"
	usersprot "github.com/ob-vss-ws19/blatt-4-pwn2own/users/proto"
)

func main() {
	fmt.Println("Start Test Skript")
	clientService := micro.NewService(micro.Name("Client"))
	clientService.Init()

	fmt.Println("Creating 5 Movies")
	_, moviearray := createTestMovies(clientService)

	fmt.Println("Creating 3 Cinemas")
	cinemaService, cinemaarray := createTestCinemas(clientService)

	fmt.Println("Creating 3 Shows")
	showService, _ := createTestShows(clientService, moviearray, cinemaarray)

	fmt.Println("Creating 6 Users")
	createTestUsers(clientService)

	//fmt.Println("Creating 5 Reservation")
	//reservationService, _ := createTestReservations(clientService)

	fmt.Println("Start Scenario 1")

	fmt.Printf("Delete Cinema with id: %d\n", cinemaarray[2])
	response, err := cinemaService.Delete(context.TODO(), &cinemaprot.DeleteCinemaRequest{Id: cinemaarray[2]})
	if err != nil {
		fmt.Println(err)
	}
	if response.Answer {
		fmt.Println("Deleting the cinema was successfull")
	} else {
		fmt.Println("Deleting the cinema failed")
	}

	response1, err1 := showService.ListShow(context.TODO(), &showprot.ListShowRequest{})
	if err1 != nil {
		fmt.Println(err1)
	}
	//list all shows
	for k := range response1.ShowId {
		println("ShowID: " + strconv.Itoa(int(response1.ShowId[k])) + " CinemaID: " + strconv.Itoa(int(response1.AllShowsData[k].CinemaId)) + " MovieID: " + strconv.Itoa(int(response1.AllShowsData[k].MovieId)))
	}

	/*
		response2, err2 := reservationService.StreamUsersReservations(context.TODO(), &reservationprot.StreamUsersReservationsRequest{})
		if err2 != nil {
			fmt.Println(err1)
		}
		//list all reservations
		for k := range response2.Reservations {
			println("ReservationID: " + strconv.Itoa(int(response2.Reservations[k].ResId)) + " Show " + strconv.Itoa(int(response2.Reservations[k].Show)) + " User " + strconv.Itoa(int(response2.Reservations[k].User)))
			for i := range response2.Reservations[k].Seats {
				println("Seat: " + strconv.Itoa(int(response2.Reservations[k].Seats[i].Seat)))
			}
		}
	*/
	fmt.Println("Start Scenario 2")

}

func createTestMovies(service micro.Service) (moviesprot.MoviesService, []int32) {
	movieService := moviesprot.NewMoviesService("movies", service.Client())
	arr := make([]int32, 5)

	for i := 1; i < 6; i++ {
		response, err := movieService.CreateMovie(context.TODO(), &moviesprot.CreateMovieRequest{Name: "Movie" + strconv.Itoa(i)})
		if err != nil {
			fmt.Println(err)
		}
		if response != nil {
			arr[i-1] = response.Movie.Id
			fmt.Printf("Adding Movie succeeded; id: %d, name: %s\n", response.Movie.Id, response.Movie.Name)
		} else {
			fmt.Println("Error - repsonse is nil")
		}
	}
	return movieService, arr
}

func createTestCinemas(service micro.Service) (cinemaprot.CinemaService, []int32) {
	cinemaService := cinemaprot.NewCinemaService("cinema-service", service.Client())
	arr := make([]int32, 4)

	for i := 1; i < 4; i++ {
		response, err := cinemaService.Create(context.TODO(), &cinemaprot.CreateCinemaRequest{Name: "Cinema" + strconv.Itoa(i), Row: int32(5 * i), Column: int32(5 * i)})
		if err != nil {
			fmt.Println(err)
		}
		if response != nil {
			arr[i-1] = response.Id
			fmt.Printf("Adding Cinema succeeded; id: %d, name: %s\n", response.Id, response.Name)
		} else {
			fmt.Println("Error - repsonse is nil")
		}
	}
	return cinemaService, arr
}

func createTestShows(service micro.Service, moviearr []int32, cinemaarr []int32) (showprot.ShowService, []int32) {
	showService := showprot.NewShowService("show-service", service.Client())
	arr := make([]int32, 3)

	for i := 1; i < 4; i++ {
		response, err := showService.CreateShow(context.TODO(), &showprot.CreateShowRequest{CreateData: &showprot.ShowMessage{CinemaId: cinemaarr[i-1], MovieId: moviearr[i-1]}})
		if err != nil {
			fmt.Println(err)
		}
		if response != nil {
			arr[i-1] = response.CreateShowId
			fmt.Printf("Adding Show succeeded; id: %d\n", response.CreateShowId)
		} else {
			fmt.Println("Error - repsonse is nil")
		}
	}
	return showService, arr
}

func createTestUsers(service micro.Service) (usersprot.UsersService, []int32) {
	userService := usersprot.NewUsersService("users", service.Client())
	arr := make([]int32, 6)

	for i := 1; i < 7; i++ {
		response, err := userService.CreateUser(context.TODO(), &usersprot.CreateUserRequest{Name: "User" + strconv.Itoa(i)})
		if err != nil {
			fmt.Println(err)
		}
		if response != nil {
			arr[i-1] = response.User.Userid
			fmt.Printf("Adding User succeeded; id: %d, name: %s\n", response.User.Userid, response.User.Name)
		} else {
			fmt.Println("Error - repsonse is nil")
		}
	}
	return userService, arr
}

func createTestReservations(service micro.Service) (reservationprot.ReservationService, []int32) {
	reservationService := reservationprot.NewReservationService("registration", service.Client())
	arr := []int32{}

	/*response, err := reservationService.

	if err != nil {
		fmt.Println(err)
	}
	*/

	return reservationService, arr
}
