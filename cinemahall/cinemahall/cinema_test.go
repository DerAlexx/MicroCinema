package cinemahall_test

import (
	"testing"

	"github.com/ob-vss-ws19/blatt-4-pwn2own/cinemahall/cinemahall"
	protoo "github.com/ob-vss-ws19/blatt-4-pwn2own/cinemahall/proto"
)

/*
TestCreate will be a testcase for adding a cinema into the service.
*/
func TestCreate(t *testing.T) {
	TestName := "C1"
	TestId := 1
	service := cinemahall.NewCinemaPool()
	response := protoo.CreateCinemaResponse{}
	service.Create(nil, &protoo.CreateCinemaRequest{Name: TestName, Row: 5, Column: 5}, &response)

	if response.Name != "C1" {
		t.Errorf("Cannot create a cinema with the name %s", C1)
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
	TestId := 1
	service := cinemahall.NewCinemaPool()
	response := protoo.CreateCinemaResponse{}
	service.Create(nil, &protoo.CreateCinemaRequest{Name: TestName, Row: 5, Column: 5}, &response)
	responseDelete := protoo.DeleteCinemaResponse{}
	service.Delete(nil, &protoo.CreateCinemaRequest{Id: response.Id}, &responseDelete)

	if !responseDelete {
		t.Errorf("Cannot delete the cinema with the namide %d", TestId)
	} else {
		t.Log("Deleting a Cinema will work.")
	}
}

/*
TestReservation will be a testcase for doing a reservation.
*/
func TestReservation(t *testing.T) {
	TestName := "C1"
	TestId := 1
	service := cinemahall.NewCinemaPool()
	response := protoo.CreateCinemaResponse{}
	service.Create(nil, &protoo.CreateCinemaRequest{Name: TestName, Row: 5, Column: 5}, &response)
	responseReservation := protoo.ReservationResponse{}
	service.Reservation(nil, &protoo.ReservationRequest{Id: response.Id, &protoo.SeatMessage{Row: 1, Column: 1}}, &responseReservation)

	if !responseReservation {
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
	TestId := 1
	service := cinemahall.NewCinemaPool()
	response := protoo.CreateCinemaResponse{}
	service.Create(nil, &protoo.CreateCinemaRequest{Name: TestName, Row: 5, Column: 5}, &response)
	responseReservation := protoo.ReservationResponse{}
	service.Reservation(nil, &protoo.ReservationRequest{Id: response.Id, &protoo.SeatMessage{Row: 1, Column: 1}}, &responseReservation)
	responseStorno := protoo.StornoResponse{}
	service.Reservation(nil, &protoo.StornoRequest{Id: response.Id, &protoo.SeatMessage{Row: 1, Column: 1}}, &responseStorno)

	if !responseStorno {
		t.Error("Storno failed")
	} else {
		t.Log("Storno will work in a cinema.")
	}
}

/*
TestReset will be a testcase resetting a cinemapool.
*/
func TestReset(t *testing.T) {
	TestName := "C1"
	TestId := 1
	service := cinemahall.NewCinemaPool()
	response := protoo.CreateCinemaResponse{}
	service.Create(nil, &protoo.CreateCinemaRequest{Name: TestName, Row: 5, Column: 5}, &response)
	responseReset := protoo.ResetResponse{}
	service.Reset(nil, &protoo.ResetRequest{Id: response.Id}, &responseReset)

	if !responseReset {
		t.Error("Reset failed")
	} else {
		t.Log("Reset will work for a cinemapool.")
	}
}

/*
TestReset will be a testcase resetting a cinemapool.
*/
func TestReset(t *testing.T) {
	TestName := "C1"
	TestId := 1
	service := cinemahall.NewCinemaPool()
	response := protoo.CreateCinemaResponse{}
	service.Create(nil, &protoo.CreateCinemaRequest{Name: TestName, Row: 5, Column: 5}, &response)
	responseReset := protoo.ResetResponse{}
	service.Reset(nil, &protoo.ResetRequest{Id: response.Id}, &responseReset)

	if !responseReset {
		t.Error("Reset failed")
	} else {
		t.Log("Reset will work for a cinemapool.")
	}
}

/*
CheckSeats will be a testcase to check seats.
*/
func TestCheckSeats(t *testing.T) {
	TestName := "C1"
	TestId := 1
	service := cinemahall.NewCinemaPool()
	response := protoo.CreateCinemaResponse{}
	service.Create(nil, &protoo.CreateCinemaRequest{Name: TestName, Row: 5, Column: 5}, &response)
	responseReservation := protoo.ReservationResponse{}
	service.Reservation(nil, &protoo.ReservationRequest{Id: response.Id, &protoo.SeatMessage{Row: 1, Column: 1}}, &responseReservation)
	responseCheckSeats := protoo.CheckSeatsResponse{}
	service.Reservation(nil, &protoo.CheckSeatsRequest{Id: response.Id, &protoo.SeatMessage{Row: 1, Column: 1}}, &responseCheckSeats)

	if responseCheckSeats {
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
	TestId := 1
	service := cinemahall.NewCinemaPool()
	response := protoo.CreateCinemaResponse{}
	service.Create(nil, &protoo.CreateCinemaRequest{Name: TestName, Row: 2, Column: 2}, &response)
	responseReservation := protoo.ReservationResponse{}
	service.Reservation(nil, &protoo.ReservationRequest{Id: response.Id, &protoo.SeatMessage{Row: 1, Column: 1}}, &responseReservation)
	responseFreeSeats := protoo.FreeSeatsResponse{}
	service.Reservation(nil, &protoo.CheckSeatsRequest{Id: response.Id, &responseFreeSeats)

	if response.Freeseats.Row != 2 || response.Freeseats.Column != 2{
		t.Error("FreeSeats failed")
	} else {
		t.Log("FreeSeats will work in a cinema .")
	}
}
