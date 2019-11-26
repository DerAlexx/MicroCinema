package reservation

import (
	"context"

	proto "github.com/ob-vss-ws19/blatt-4-pwn2own/reservation/proto"
)

/*
Reservation will be the representation of a reservation.
*/
type Reservation struct {
}

/*
ReservatServiceHandler will handle all reservations.
*/
type ReservatServiceHandler struct {
	reservations map[int32]Reservation
}

/*
MakeReservation will receive a Reservsation request an store it temporally in the Database.
*/
func (r *ReservatServiceHandler) MakeReservation(ctx context.Context, in *proto.MakeReservationRequest, out *proto.MakeReservationResponse) error {
	return nil
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
