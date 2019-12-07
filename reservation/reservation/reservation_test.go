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

/*
Test to add a potantialreservation accepted it and send the same data again.
Result: First send, stored and accepted second not stored.
*/
func TestAddAcceptReservation(t *testing.T) {
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
	ina := &proto.AcceptReservationRequest{TmpID: out.TmpID, Want: true}
	outa := &proto.AcceptReservationResponse{}
	re.AcceptReservation(nil, ina, outa)

	if out.TmpID < 0 || !out.Works {
		t.Errorf("cannot add a potentialreservation into the map got ID: %d and Works: %t", out.TmpID, out.Works)
	} else if outa.FinalID < 1 && !outa.Taken {
		t.Errorf("Accepted responsed with the wrong answer: %d and Works: %t", outa.FinalID, outa.Taken)
	} else {
		t.Log("test add the same reservation 2 times worked fine.")
	}

}

/*
Test Add and accept a Reservation.
Result Potantial good final bad.
*/
func TestAddTroughAwayReservation(t *testing.T) {
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
	ina := &proto.AcceptReservationRequest{TmpID: out.TmpID, Want: false}
	outa := &proto.AcceptReservationResponse{}
	re.AcceptReservation(nil, ina, outa)

	if out.TmpID < 0 || !out.Works {
		t.Errorf("cannot add a potentialreservation into the map got ID: %d and Works: %t", out.TmpID, out.Works)
	} else if outa.FinalID > 1 && outa.Taken {
		t.Errorf("Accepted responsed with the wrong answer: %d and Works: %t", outa.FinalID, outa.Taken)
	} else {
		t.Log("test add the same reservation 2 times worked fine.")
	}

}

/*
Test Add and try to make an error a Reservation.
Result Potantial good final bad.
*/
func TestAddErrorReservation(t *testing.T) {
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
	ina := &proto.AcceptReservationRequest{TmpID: -1, Want: true}
	outa := &proto.AcceptReservationResponse{}
	re.AcceptReservation(nil, ina, outa)

	if out.TmpID < 0 || !out.Works {
		t.Errorf("cannot add a potentialreservation into the map got ID: %d and Works: %t", out.TmpID, out.Works)
	} else if outa.FinalID > 1 && outa.Taken {
		t.Errorf("Accepted responsed with the wrong answer: %d and Works: %t", outa.FinalID, outa.Taken)
	} else {
		t.Log("test add the same reservation 2 times worked fine.")
	}

}

/*
Test to add a potantialreservation accepted it and send the same data again.
Result: First send, stored and accepted second not stored.
*/
func TestAddAcceptAddAgainReservation(t *testing.T) {
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
	ina := &proto.AcceptReservationRequest{TmpID: out.TmpID, Want: true}
	outa := &proto.AcceptReservationResponse{}
	re.AcceptReservation(nil, ina, outa)
	in2 := &proto.MakeReservationRequest{
		Res: &proto.Reservation{
			User:  23,
			Show:  34,
			Seats: []*proto.Seat{&proto.Seat{Seat: 23}, &proto.Seat{Seat: 34}},
		},
	}
	out2 := &proto.MakeReservationResponse{}
	re.MakeReservation(nil, in2, out2)
	if out.TmpID < 0 || !out.Works {
		t.Errorf("cannot add a potentialreservation into the map got ID: %d and Works: %t", out.TmpID, out.Works)
	} else if out2.TmpID != -1 || out2.Works {
		t.Errorf("second time to send a reservation which should not be executed returned the wrong response: %d and Works: %t", out.TmpID, out.Works)
	} else if outa.FinalID < 1 && !outa.Taken {
		t.Errorf("Accepted responsed with the wrong answer: %d and Works: %t", outa.FinalID, outa.Taken)
	} else {
		t.Log("test add the same reservation 2 times worked fine.")
	}

}

/*
Test to add a reservation accepted it and delete it.
Result not reservation in the service anymore.
*/
func TestAddAcceptDeleteReservation(t *testing.T) {
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
	ina := &proto.AcceptReservationRequest{TmpID: out.TmpID, Want: true}
	outa := &proto.AcceptReservationResponse{}
	re.AcceptReservation(nil, ina, outa)
	din := &proto.DeleteReservationRequest{Id: outa.FinalID}
	dout := &proto.DeleteReservationResponse{}

	re.DeleteReservation(nil, din, dout)

	if out.TmpID < 0 || !out.Works {
		t.Errorf("cannot add a potentialreservation into the map got ID: %d and Works: %t", out.TmpID, out.Works)
	} else if outa.FinalID < 1 && !outa.Taken {
		t.Errorf("Accepted responsed with the wrong answer: %d and Works: %t", outa.FinalID, outa.Taken)
	} else if !dout.Deleted {
		t.Errorf("Cannot delete the reservation with the ID:%d, Taken: %t, Deleted responed: %t", outa.FinalID, outa.Taken, dout.Deleted)
	} else {
		t.Log("test add the same reservation 2 times worked fine.")
	}
}

/*
Test to add, accept and find a reservation.
Result: Return reservation.
*/
func TestAddAcceptFindReservation(t *testing.T) {
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
	ina := &proto.AcceptReservationRequest{TmpID: out.TmpID, Want: true}
	outa := &proto.AcceptReservationResponse{}
	re.AcceptReservation(nil, ina, outa)
	din := &proto.ShowReservationsRequest{Id: outa.FinalID}
	dout := &proto.ShowReservationsResponse{}

	re.ShowReservations(nil, din, dout)

	if out.TmpID < 0 || !out.Works {
		t.Errorf("cannot add a potentialreservation into the map got ID: %d and Works: %t", out.TmpID, out.Works)
	} else if outa.FinalID < 1 && !outa.Taken {
		t.Errorf("Accepted responsed with the wrong answer: %d and Works: %t", outa.FinalID, outa.Taken)
	} else if dout.Res.ResId < 1 || dout.Res.ResId != outa.FinalID || dout.Res.Show != 34 || dout.Res.Show < 1 || dout.Res.User != 23 || dout.Res.User < 1 || len(dout.Res.Seats) < 1 || len(dout.Res.Seats) != 2 {
		t.Errorf("got the wrong answer back wanted-RES-ID: %d, got: %d", dout.Res.ResId, outa.FinalID)
		t.Errorf("got the wrong answer back wanted-UserID: %d, got: %d", dout.Res.User, 23)
		t.Errorf("got the wrong answer back wanted-ShowID: %d, got: %d", dout.Res.Show, 34)
		t.Errorf("got the wrong answer back wanted-Seats: %d, got: %d", len(dout.Res.Seats), 2)
	} else {
		t.Log("test add the same reservation 2 times worked fine.")
	}
}

/*
Test to add a potantialreservation and check if a user has a reservation.
*/
func TestAddCheckHasReservationReservation(t *testing.T) {
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

	fin := &proto.HasReservationsRequest{Res: &proto.Reservation{User: 23}}
	fout := &proto.HasReservationsResponse{}

	re.HasReservations(nil, fin, fout)

	if out.TmpID < 0 || !out.Works {
		t.Errorf("cannot add a potentialreservation into the map got ID: %d and Works: %t", out.TmpID, out.Works)
	} else if fout.Amount < 1 || !fout.Has {
		t.Errorf("got the wrong answer there is a reservation but answer was wrong Amount: %d and Has: %t", fout.Amount, fout.Has)
	} else {
		t.Log("worked")
	}

}

/*
Test to add a potantialreservation and try to check whether a user has a reservation on an empty service.
*/
func TestAddCheckHasReservationOnEmptyServiceReservation(t *testing.T) {
	re := res.CreateNewReservationHandlerInstance()
	fin := &proto.HasReservationsRequest{Res: &proto.Reservation{User: 23}}
	fout := &proto.HasReservationsResponse{}

	re.HasReservations(nil, fin, fout)

	if fout.Amount > 1 || fout.Has {
		t.Errorf("got the wrong answer there is a reservation but answer was wrong Amount: %d and Has: %t", fout.Amount, fout.Has)
	} else {
		t.Log("worked")
	}

}

/*
Test to add a reservation and accepted it and stream it later on.
*/
func TestAddCheckStreamReservation(t *testing.T) {
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
	ina := &proto.AcceptReservationRequest{TmpID: out.TmpID, Want: true}
	outa := &proto.AcceptReservationResponse{}
	re.AcceptReservation(nil, ina, outa)

	sin := &proto.StreamUsersReservationsRequest{}
	sout := &proto.StreamUsersReservationsResponse{}

	re.StreamUsersReservations(nil, sin, sout)

	if out.TmpID < 0 || !out.Works {
		t.Errorf("cannot add a potentialreservation into the map got ID: %d and Works: %t", out.TmpID, out.Works)
	} else if outa.FinalID < 1 && !outa.Taken {
		t.Errorf("Accepted responsed with the wrong answer: %d and Works: %t", outa.FinalID, outa.Taken)
	} else if len(sout.Reservations) < 1 {
		t.Errorf("The length of the answer and the expectation does not match up: %d and Works: %d", len(sout.Reservations), 1)
	} else if (*sout.Reservations[0]).User != 23 || (*sout.Reservations[0]).Show != 34 || (*sout.Reservations[0]).Seats[0].Seat != 23 || (*sout.Reservations[0]).Seats[1].Seat != 34 {
		t.Errorf("The user got does not match up with the expected one: got %d wanted: %d", (*sout.Reservations[0]).User, 23)
		t.Errorf("The show got does not match up with the expected one: got %d wanted: %d", (*sout.Reservations[0]).Show, 34)
		t.Errorf("The seat 1 got does not match up with the expected one: got %d wanted: %d", (*sout.Reservations[0]).Seats[0].Seat, 23)
		t.Errorf("The seat 2 got does not match up with the expected one: got %d wanted: %d", (*sout.Reservations[0]).Seats[1].Seat, 34)
	} else {
		t.Log("test add the same reservation 2 times worked fine.")
	}
}

/*
Test to add a potantialreservation accepted it and change it.
Result: First send, stored and accepted and changed.
*/
func TestAddAcceptChangeReservation(t *testing.T) {
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
	ina := &proto.AcceptReservationRequest{TmpID: out.TmpID, Want: true}
	outa := &proto.AcceptReservationResponse{}
	re.AcceptReservation(nil, ina, outa)

	cin := &proto.ChangeReservationRequest{Res: &proto.Reservation{
		ResId: outa.FinalID,
		User:  23,
		Show:  33,
		Seats: []*proto.Seat{&proto.Seat{Seat: 23}, &proto.Seat{Seat: 34}},
	}}
	cout := &proto.ChangeReservationResponse{}

	re.ChangeReservation(nil, cin, cout)

	if out.TmpID < 0 || !out.Works {
		t.Errorf("cannot add a potentialreservation into the map got ID: %d and Works: %t", out.TmpID, out.Works)
	} else if outa.FinalID < 1 && !outa.Taken {
		t.Errorf("Accepted responsed with the wrong answer: %d and Works: %t", outa.FinalID, outa.Taken)
	} else if !cout.Changed {
		t.Errorf("The value did not changed --> Got %t", cout.Changed)
	} else {
		t.Log("test add the same reservation 2 times worked fine.")
	}

}
