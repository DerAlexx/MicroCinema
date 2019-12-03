package cinemahall_test

import (
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
	service.Create(nil, &cinemaprototest.CreateCinemaRequest{Name: TestName, Row: 5, Column: 5}, &response)

	if response.Name != "C1" {
		t.Errorf("Cannot create a cinema with the name %s", TestName)
	} else if response.Id < 0 {
		t.Fatal("Cannot create a cinema with a proper ID")
	} else {
		t.Log("Creating a Cinema will work.")
	}
}

/*
TestDelete will be a testcase for deleting a cinema from the service.
*/
func TestDelete(t *testing.T) {
	TestName := "C1"
	service := cinemahall.NewCinemaPool()
	response := cinemaprototest.CreateCinemaResponse{}
	service.Create(nil, &cinemaprototest.CreateCinemaRequest{Name: TestName, Row: 5, Column: 5}, &response)
	responseDelete := cinemaprototest.DeleteCinemaResponse{}
	service.Delete(nil, &cinemaprototest.DeleteCinemaRequest{Id: response.Id}, &responseDelete)

	if !responseDelete.Answer {
		t.Errorf("Cannot delete the cinema with the namide %d", 1)
	} else {
		t.Log("Deleting a Cinema will work.")
	}
}

/*
TestReservation will be a testcase for doing a reservation.
*/
func TestReservation(t *testing.T) {
	TestName := "C1"
	service := cinemahall.NewCinemaPool()
	response := cinemaprototest.CreateCinemaResponse{}
	service.Create(nil, &cinemaprototest.CreateCinemaRequest{Name: TestName, Row: 5, Column: 5}, &response)
	responseReservation := cinemaprototest.ReservationResponse{}
	x := []*cinemaprototest.SeatMessage{}
	x = append(x, &cinemaprototest.SeatMessage{Row: 1, Column: 1})
	service.Reservation(nil, &cinemaprototest.ReservationRequest{Id: response.Id, Seatreservation: x}, &responseReservation)

	if !responseReservation.Answer {
		t.Error("Reservation failed")
	} else {
		t.Log("Reservation for a seat will work in a cinema .")
	}
}

/*
TestStorno will be a testcase for doing a storno.
*/
func TestStorno(t *testing.T) {
	TestName := "C1"
	service := cinemahall.NewCinemaPool()
	response := cinemaprototest.CreateCinemaResponse{}
	service.Create(nil, &cinemaprototest.CreateCinemaRequest{Name: TestName, Row: 5, Column: 5}, &response)
	responseReservation := cinemaprototest.ReservationResponse{}
	x := []*cinemaprototest.SeatMessage{}
	x = append(x, &cinemaprototest.SeatMessage{Row: 1, Column: 1})
	service.Reservation(nil, &cinemaprototest.ReservationRequest{Id: response.Id, Seatreservation: x}, &responseReservation)
	responseStorno := cinemaprototest.StornoResponse{}
	service.Storno(nil, &cinemaprototest.StornoRequest{Id: response.Id, Seatstorno: x}, &responseStorno)

	if !responseStorno.Answer {
		t.Error("Storno failed")
	} else {
		t.Log("Storno will work in a cinema.")
	}
}

/*
CheckSeats will be a testcase to CheckSeats.
*/
func TestCheckSeats(t *testing.T) {
	TestName := "C1"
	service := cinemahall.NewCinemaPool()
	response := cinemaprototest.CreateCinemaResponse{}
	service.Create(nil, &cinemaprototest.CreateCinemaRequest{Name: TestName, Row: 5, Column: 5}, &response)
	responseReservation := cinemaprototest.ReservationResponse{}
	x := []*cinemaprototest.SeatMessage{}
	x = append(x, &cinemaprototest.SeatMessage{Row: 1, Column: 1})
	service.Reservation(nil, &cinemaprototest.ReservationRequest{Id: response.Id, Seatreservation: x}, &responseReservation)
	responseCheckSeats := cinemaprototest.CheckSeatsResponse{}
	service.CheckSeats(nil, &cinemaprototest.CheckSeatsRequest{Id: response.Id, Seatcheck: x}, &responseCheckSeats)

	if responseCheckSeats.Answer {
		t.Error("CheckSeats failed")
	} else {
		t.Log("CheckSeats will work in a cinema .")
	}
}

/*
FreeSeats will be a testcase to check FreeSeats.
*/
func TestFreeSeats(t *testing.T) {
	TestName := "C1"
	service := cinemahall.NewCinemaPool()
	response := cinemaprototest.CreateCinemaResponse{}
	service.Create(nil, &cinemaprototest.CreateCinemaRequest{Name: TestName, Row: 2, Column: 2}, &response)
	responseReservation := cinemaprototest.ReservationResponse{}
	x := []*cinemaprototest.SeatMessage{}
	x = append(x, &cinemaprototest.SeatMessage{Row: 1, Column: 1})
	service.Reservation(nil, &cinemaprototest.ReservationRequest{Id: response.Id, Seatreservation: x}, &responseReservation)
	responseFreeSeats := cinemaprototest.FreeSeatsResponse{}
	service.FreeSeats(nil, &cinemaprototest.FreeSeatsRequest{Id: response.Id}, &responseFreeSeats)

	if len(responseFreeSeats.Freeseats) != 3 {
		t.Errorf("FreeSeats failed; len: %d", len(responseFreeSeats.Freeseats))
	} else {
		t.Log("FreeSeats will work in a cinema .")
	}
}
