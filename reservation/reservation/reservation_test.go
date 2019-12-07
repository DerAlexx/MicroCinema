package reservation_test

import (
	"testing"

	proto "github.com/ob-vss-ws19/blatt-4-pwn2own/reservation/proto"
	res "github.com/ob-vss-ws19/blatt-4-pwn2own/reservation/reservation"
)

/*
Test to add a potantialreservation.
*/
func TestAddReservation(t *testing.T) {
	re := res.CreateNewReservationHandlerInstance()
	in := &proto.MakeReservationRequest{
		Res: &proto.Reservation{
			User:  23,
			Show:  34,
			Seats: []*proto.Seat{&proto.Seat{Seat: 23}, &proto.Seat{Seat: 34}},
		},
	}
	out := &proto.MakeReservationResponse{}
	re.MakeReservation(nil, in, out)
	if out.TmpID < 0 || !out.Works {
		t.Errorf("cannot add a potentialreservation into the map got ID: %d and Works: %t", out.TmpID, out.Works)
	}

}
