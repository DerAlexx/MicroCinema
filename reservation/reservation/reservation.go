package reservation

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	proto "github.com/ob-vss-ws19/blatt-4-pwn2own/reservation/proto"
)

const (
	maxreservationsid int32 = 987654321
)

/*
containsID will check whether generated ID is already set or not.
@id will be a int32 id to check for.
*/
func (r *ReservatServiceHandler) containsID(id int32) bool {
	_, inMap := (*r.getReservationsMap())[id]
	return inMap
}

/*
Function to produce a random reservationsID.
@param length will be the length of the number.
@param seed will be a seed in order to produce "really" random numbers.
*/
func (r *ReservatServiceHandler) getRandomReservationsID(length int32) int32 {
	rand.Seed(time.Now().UnixNano())
	for {
		potantialID := rand.Int31n(length)
		if !r.containsID(int32(potantialID)) {
			return potantialID
		}
	}
}

/*
getReservationsMap will return a pointer to the current reservations map in order to work in that.
*/
func (r *ReservatServiceHandler) getReservationsMap() *map[int32]*Reservation {
	return &r.reservations
}

/*
Reservation will be the representation of a reservation.
*/
type Reservation struct {
	UserID int32
	ShowID int32
}

/*
ReservatServiceHandler will handle all reservations.
*/
type ReservatServiceHandler struct {
	reservations map[int32]*Reservation
	dependencies []interface{}
	mutex        *sync.Mutex
}

/*
CreateNewReservationHandlerInstance will create a new service for the reservations management.
*/
func CreateNewReservationHandlerInstance() *ReservatServiceHandler {
	return &ReservatServiceHandler{
		reservations: make(map[int32]*Reservation),
		mutex:        &sync.Mutex{},
	}
}

/*
MakeReservation will receive a Reservsation request an store it temporally in the Database.
*/
func (r *ReservatServiceHandler) MakeReservation(ctx context.Context, in *proto.MakeReservationRequest, out *proto.MakeReservationResponse) error {
	if len(in.Res) > 0 {
		r.mutex.Lock()

		r.mutex.Unlock()
		return nil
	}
	return fmt.Errorf("cannot create a reservation with an list of size: %d", len(in.Res))
}

/*
AcceptReservation will accept a reservation of a temporally stored reservation request.
*/
func (r *ReservatServiceHandler) AcceptReservation(ctx context.Context, in *proto.AcceptReservationRequest, out *proto.AcceptReservationResponse) error {
	return nil
}

/*
DeleteReservation will delete a final reservation from the database.
*/
func (r *ReservatServiceHandler) DeleteReservation(ctx context.Context, in *proto.DeleteReservationRequest, out *proto.DeleteReservationResponse) error {
	return nil
}

/*
ChangeReservation will change a reservation for example the place of the reservation.
*/
func (r *ReservatServiceHandler) ChangeReservation(ctx context.Context, in *proto.ChangeReservationRequest, out *proto.ChangeReservationResponse) error {
	return nil
}

/*
ShowReservations will send a reservation to the client
*/
func (r *ReservatServiceHandler) ShowReservations(ctx context.Context, in *proto.ShowReservationsRequest, out *proto.ShowReservationsResponse) error {
	return nil
}

/*
StreamUsersReservations will send all reservations.
*/
func (r *ReservatServiceHandler) StreamUsersReservations(ctx context.Context, in *proto.StreamUsersReservationsRequest, out *proto.StreamUsersReservationsResponse) error {
	return nil
}

/*
HasReservations will send a bool to indicate.
*/
func (r *ReservatServiceHandler) HasReservations(ctx context.Context, in *proto.HasReservationsRequest, out *proto.HasReservationsResponse) error {
	return nil
}
