package main

import (
	"context"
	"fmt"

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
	if clientService == nil {
		fmt.Println("fuck")
	}

	fmt.Println("Creating 5 Movies")
	//_, moviearray := createTestMovies(clientService)

	fmt.Println("Creating 3 Cinemas")
	///cinemaService, cinemaarray := createTestCinemas(clientService)

	fmt.Println("Creating 3 Shows")
	//showService, _ := createTestShows(clientService, moviearray, cinemaarray)

	fmt.Println("Creating 6 Users")
	createTestUsers(clientService)
	/*
		fmt.Println("Creating 5 Reservation")
		reservationService, _ := createTestReservations(clientService)

		fmt.Println("Start Scenario 1")

		fmt.Printf("Delete Cinema with id: %d", cinemaarray[2])
		response, err := cinemaService.Delete(context.TODO(), &cinemaprot.DeleteCinemaRequest{Id: cinemaarray[2]})
		if err != nil {
			fmt.Println(err)
		}
		if response.Answer {
			fmt.Println("Deleting the cinema was successfull")
		}
		fmt.Println("Deleting the cinema failed")

		response1, err1 := showService.ListShow(context.TODO(), &showprot.ListShowRequest{})
		if err1 != nil {
			fmt.Println(err1)
		}
		//list all shows
		for k := range response1.ShowId {
			println("ShowID: " + string(response1.ShowId[k]))
			println("CinemaID: " + string(response1.AllShowsData[k].CinemaId) + " MovieID: " + string(response1.AllShowsData[k].MovieId))
		}

		response2, err2 := reservationService.StreamUsersReservations(context.TODO(), &reservationprot.StreamUsersReservationsRequest{})
		if err2 != nil {
			fmt.Println(err1)
		}
		//list all reservations
		for k := range response2.Reservations {
			println("ReservationID: " + string(response2.Reservations[k].ResId) + "Show " + string(response2.Reservations[k].Show) + "User " + string(response2.Reservations[k].User))
			for i := range response2.Reservations[k].Seats {
				println("Seat: " + string(response2.Reservations[k].Seats[i].Seat))
			}
		}

		fmt.Println("Start Scenario 2")
	*/
}

func createTestMovies(service micro.Service) (moviesprot.MoviesService, []int32) {
	movieService := moviesprot.NewMoviesService("movies", service.Client())
	println("hallo1")

	arr := []int32{}

	//for i := 1; i < 5; i++ {
	response, err := movieService.CreateMovie(context.TODO(), &moviesprot.CreateMovieRequest{Name: "Movie1"})
	if err != nil {
		fmt.Println(err)
	}
	if response == nil {
		fmt.Println(string(response.Movie.Id))
	}
	//arr[1] = response.Movie.Id
	//fmt.Printf("Adding Movie succeeded; id: %d, name: %s", response.Movie.Id, response.Movie.Name)
	println("hallo2")
	//}
	return movieService, arr
}

func createTestCinemas(service micro.Service) (cinemaprot.CinemaService, []int32) {
	cinemaService := cinemaprot.NewCinemaService("cinema-service", service.Client())
	arr := []int32{}

	for i := 1; i < 3; i++ {
		response, err := cinemaService.Create(context.TODO(), &cinemaprot.CreateCinemaRequest{Name: "Cinema" + string(i), Row: int32(5 * i), Column: int32(5 * i)})
		if err != nil {
			fmt.Println(err)
		}
		arr[i] = response.Id
		fmt.Printf("Adding Cinema succeeded; id: %d, name: %s", response.Id, response.Name)
	}
	return cinemaService, arr
}

func createTestShows(service micro.Service, moviearr []int32, cinemaarr []int32) (showprot.ShowService, []int32) {
	showService := showprot.NewShowService("show-service", service.Client())
	arr := []int32{}

	for i := 1; i < 3; i++ {
		response, err := showService.CreateShow(context.TODO(), &showprot.CreateShowRequest{CreateData: &showprot.ShowMessage{CinemaId: cinemaarr[len(cinemaarr)-i], MovieId: moviearr[len(moviearr)-i]}})
		if err != nil {
			fmt.Println(err)
		}
		arr[i] = response.CreateShowId
		fmt.Printf("Adding Show succeeded; id: %d", response.CreateShowId)
	}
	return showService, arr
}

func createTestUsers(service micro.Service) (usersprot.UsersService, []int32) {
	userService := usersprot.NewUsersService("users", service.Client())
	arr := []int32{}

	for i := 1; i < 6; i++ {
		response, err := userService.CreateUser(context.TODO(), &usersprot.CreateUserRequest{Name: ""})
		if err != nil {
			fmt.Println(err)
		}
		if response == nil {
			fmt.Println(string(response.User.Name))
		}
		//arr[i] = response.User.Userid
		//fmt.Printf("Adding User succeeded; id: %d, name: %s", response.User.Userid, response.User.Name)
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
