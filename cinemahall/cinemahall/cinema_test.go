package cinemahall_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/ob-vss-ws19/blatt-4-pwn2own/cinemahall/cinemahall"
	cinemaprototest "github.com/ob-vss-ws19/blatt-4-pwn2own/cinemahall/proto"
)

/*
TestCreate will be a testcase for adding a cinema into the service.
*/
func TestCreate(t *testing.T) {
	TestName := "C1"
	service := cinemahall.NewCinemaPool()
	response := cinemaprototest.CreateCinemaResponse{}
	err := service.Create(context.TODO(), &cinemaprototest.CreateCinemaRequest{Name: TestName, Row: 5, Column: 5}, &response)
	if err == nil {
		switch {
		case response.Name != "C1":
			t.Errorf("Cannot create a cinema with the name %s", TestName)
		case response.Id < 0:
			t.Fatal("Cannot create a cinema with a proper ID")
		default:
			t.Log("Creating a Cinema will work.")
		}
	} else {
		fmt.Println(err)
	}
}

/*
TestDelete will be a testcase for deleting a cinema from the service.
*/
func TestDelete(t *testing.T) {
	TestName := "C1"
	service := cinemahall.NewCinemaPool()
	response := cinemaprototest.CreateCinemaResponse{}
	err := service.Create(context.TODO(), &cinemaprototest.CreateCinemaRequest{Name: TestName, Row: 5, Column: 5}, &response)
	responseDelete := cinemaprototest.DeleteCinemaResponse{}
	err1 := service.Delete(context.TODO(), &cinemaprototest.DeleteCinemaRequest{Id: response.Id}, &responseDelete)
	if err == nil && err1 == nil {
		if !responseDelete.Answer {
			t.Errorf("Cannot delete the cinema with the namide %d", 1)
		} else {
			t.Log("Deleting a Cinema will work.")
		}
	} else {
		fmt.Println(err)
		fmt.Println(err1)
	}
}

/*
TestReservation will be a testcase for doing a reservation.
*/
func TestReservation(t *testing.T) {
	TestName := "C1"
	service := cinemahall.NewCinemaPool()
	response := cinemaprototest.CreateCinemaResponse{}
	err := service.Create(context.TODO(), &cinemaprototest.CreateCinemaRequest{Name: TestName, Row: 5, Column: 5}, &response)
	responseReservation := cinemaprototest.ReservationResponse{}
	x := []*cinemaprototest.SeatMessage{}
	x = append(x, &cinemaprototest.SeatMessage{Row: 1, Column: 1})
	err1 := service.Reservation(context.TODO(), &cinemaprototest.ReservationRequest{Id: response.Id, Seatreservation: x}, &responseReservation)

	if err == nil && err1 == nil {
		if !responseReservation.Answer {
			t.Error("Reservation failed")
		} else {
			t.Log("Reservation for a seat will work in a cinema .")
		}
	} else {
		fmt.Println(err)
		fmt.Println(err1)
	}
}

/*
TestStorno will be a testcase for doing a storno.
*/
func TestStorno(t *testing.T) {
	TestName := "C1"
	service := cinemahall.NewCinemaPool()
	response := cinemaprototest.CreateCinemaResponse{}
	err := service.Create(context.TODO(), &cinemaprototest.CreateCinemaRequest{Name: TestName, Row: 5, Column: 5}, &response)
	responseReservation := cinemaprototest.ReservationResponse{}
	x := []*cinemaprototest.SeatMessage{}
	x = append(x, &cinemaprototest.SeatMessage{Row: 1, Column: 1})
	err1 := service.Reservation(context.TODO(), &cinemaprototest.ReservationRequest{Id: response.Id, Seatreservation: x}, &responseReservation)
	responseStorno := cinemaprototest.StornoResponse{}
	err2 := service.Storno(context.TODO(), &cinemaprototest.StornoRequest{Id: response.Id, Seatstorno: x}, &responseStorno)

	if err == nil && err1 == nil && err2 == nil {
		if !responseStorno.Answer {
			t.Error("Storno failed")
		} else {
			t.Log("Storno will work in a cinema.")
		}
	} else {
		fmt.Println(err)
		fmt.Println(err1)
		fmt.Println(err2)
	}
}

/*
CheckSeats will be a testcase to CheckSeats.
*/
func TestCheckSeats(t *testing.T) {
	TestName := "C1"
	service := cinemahall.NewCinemaPool()
	response := cinemaprototest.CreateCinemaResponse{}
	err := service.Create(context.TODO(), &cinemaprototest.CreateCinemaRequest{Name: TestName, Row: 5, Column: 5}, &response)
	responseReservation := cinemaprototest.ReservationResponse{}
	x := []*cinemaprototest.SeatMessage{}
	x = append(x, &cinemaprototest.SeatMessage{Row: 1, Column: 1})
	err1 := service.Reservation(context.TODO(), &cinemaprototest.ReservationRequest{Id: response.Id, Seatreservation: x}, &responseReservation)
	responseCheckSeats := cinemaprototest.CheckSeatsResponse{}
	err2 := service.CheckSeats(context.TODO(), &cinemaprototest.CheckSeatsRequest{Id: response.Id, Seatcheck: x}, &responseCheckSeats)

	if err == nil && err1 == nil && err2 == nil {
		if responseCheckSeats.Answer {
			t.Error("CheckSeats failed")
		} else {
			t.Log("CheckSeats will work in a cinema .")
		}
	} else {
		fmt.Println(err)
		fmt.Println(err1)
		fmt.Println(err2)
	}
}

/*
FreeSeats will be a testcase to check FreeSeats.
*/
func TestFreeSeats(t *testing.T) {
	TestName := "C1"
	service := cinemahall.NewCinemaPool()
	response := cinemaprototest.CreateCinemaResponse{}
	err := service.Create(context.TODO(), &cinemaprototest.CreateCinemaRequest{Name: TestName, Row: 2, Column: 2}, &response)
	responseReservation := cinemaprototest.ReservationResponse{}
	x := []*cinemaprototest.SeatMessage{}
	x = append(x, &cinemaprototest.SeatMessage{Row: 1, Column: 1})
	err1 := service.Reservation(context.TODO(), &cinemaprototest.ReservationRequest{Id: response.Id, Seatreservation: x}, &responseReservation)
	responseFreeSeats := cinemaprototest.FreeSeatsResponse{}
	err2 := service.FreeSeats(context.TODO(), &cinemaprototest.FreeSeatsRequest{Id: response.Id}, &responseFreeSeats)
	if err == nil && err1 == nil && err2 == nil {
		if len(responseFreeSeats.Freeseats) != 3 {
			t.Errorf("FreeSeats failed; len: %d", len(responseFreeSeats.Freeseats))
		} else {
			t.Log("FreeSeats will work in a cinema .")
		}
	} else {
		fmt.Println(err)
		fmt.Println(err1)
		fmt.Println(err2)
	}
}
