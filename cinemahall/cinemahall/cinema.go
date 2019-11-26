package cinemahall

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	proto "github.com/ob-vss-ws19/blatt-4-pwn2own/cinemahall/proto"
)

const (
	maxcinemaid = 432342
)

/*
CinemaPool contains all cinemas.
*/
type CinemaPool struct {
	cinemamap map[int64]*cinema
	mutex     *sync.Mutex
}

type seats struct {
	row    int
	column int
}

type cinema struct {
	name    string
	seatmap map[*seats]bool
}

/*
NewCinemaPool creates a new CinemaPool
*/
func NewCinemaPool() *CinemaPool {
	newcinema := make(map[int64]*cinema)

	return &CinemaPool{
		mutex:     &sync.Mutex{},
		cinemamap: newcinema,
	}
}

func (handler *CinemaPool) getRandomCinemaID() int64 {
	rand.Seed(time.Now().UnixNano())
	for {
		potentialID := int64(rand.Intn(maxcinemaid))
		if !handler.containscinema(potentialID) {
			return potentialID
		}
	}
}

func (handler *CinemaPool) containscinema(id int64) bool {
	_, containsinmap := handler.cinemamap[id]
	return containsinmap
}

func (handler *CinemaPool) containsandreturncinema(id int64) (*cinema, bool) {
	currentcinema, mapcontainscinema := handler.cinemamap[id]
	return currentcinema, mapcontainscinema
}

func (currentcinema *cinema) containsSeatMap(row int64, column int64) bool {
	for seat, occ := range currentcinema.seatmap {
		if int64(seat.row) == row && int64(seat.column) == column {
			return true
		}
	}
	return false
}

func (currentcinema *cinema) getSeat(row int64, column int64) *seats {
	for seat, occ := range currentcinema.seatmap {
		if int64(seat.row) == row && int64(seat.column) == column {
			return seat
		}
	}
	return nil
}

/*
Create creates a new cinema. It will be saved in the CinemaPool. The id to access the new cinema will be randomly generated.
If the creation was successfull the id will be returned
*/
func (handler *CinemaPool) Create(ctx context.Context, in *proto.CreateCinemaRequest, out *proto.CreateCinemaResponse) error {
	fmt.Printf("Cinema Service - Create | Creating cinema with name: %s, %d rows: and columns: %d \n", in.Name, in.Row, in.Column)

	if len(in.Name) == 0 || in.Row == 0 || in.Column == 0 {
		return errors.New("Cinema Service - Create | Cannot create a cinema with an empty name, zero rows or zero columns")
	}

	newseatmap := map[*seats]bool{}
	for i := int64(0); i < in.Row; i++ {
		for j := int64(0); j < in.Column; j++ {
			newseatmap[&seats{row: int(i), column: int(j)}] = false
		}
	}

	createid := handler.getRandomCinemaID()
	handler.mutex.Lock()
	handler.cinemamap[createid] = &cinema{name: in.Name, seatmap: newseatmap}
	handler.mutex.Unlock()
	out.Name = in.Name
	out.Id = createid

	fmt.Println("Cinema Service - Create | Successfully created cinema")

	return nil
}

/*
Delete will delete a cinema(id) from the CinemaPool.
*/
func (handler *CinemaPool) Delete(ctx context.Context, in *proto.DeleteCinemaRequest, out *proto.DeleteCinemaResponse) error {
	fmt.Printf("Cinema Service - Delete | Deleting cinema with id %d\n", in.Id)
	handler.mutex.Lock()
	mapcontainscinema := handler.containscinema(in.Id)
	if !mapcontainscinema {
		handler.mutex.Unlock()
		out.Answer = false
		return fmt.Errorf("Cinema Service - Delete | Cannot delete cinema with id: %d", in.Id)
	}
	delete(handler.cinemamap, in.Id)
	handler.mutex.Unlock()
	out.Answer = true

	fmt.Println("Cinema Service - Delete | Successfully deleted cinema")
	return nil
}

/*
Reservation will change the value in the seatmap of a cinema to true (for the given row and column) if the seat is still available
*/
func (handler *CinemaPool) Reservation(ctx context.Context, in *proto.ReservationRequest, out *proto.ReservationResponse) error {
	fmt.Printf("Cinema Service - Reservation | Reservating seats in row %d and column &d for cinema id %d\n", in.Seatreservation, in.Seatreservation, in.Id) //Fehler

	handler.mutex.Lock()
	currentcinema, mapcontainscinema := handler.containsandreturncinema(in.Id)

	if !mapcontainscinema {
		handler.mutex.Unlock()
		out.Answer = false
		return errors.New("Cinema Service - Reservation | Cannot execute reservation because cinema doesnt exist")
	}
	for curreservation := range in.Seatreservation {
		if currentcinema.containsSeatMap(*curreservation.Row, curreservation.column) {
			currentcinema.seatmap[currentcinema.getSeat(curreservation.row, curreservation.column)] = true
		}
	}
	handler.cinemamap[in.Id] = currentcinema
	handler.mutex.Unlock()

	fmt.Println("Cinema Service - Reservation | Reservation successfull")
	out.Answer = true
	return nil
}

/*
Storno will undo a reservation
*/
func (handler *CinemaPool) Storno(ctx context.Context, in *proto.StornoRequest, out *proto.StornoResponse) error {
	fmt.Println("Cinema Service - Storno | Deleting reservation")

	handler.mutex.Lock()
	currentcinema, mapcontainscinema := handler.containsandreturncinema(in.Id)

	if !mapcontainscinema {
		handler.mutex.Unlock()
		out.Answer = false
		return errors.New("Cinema Service - Storno | Cannot execute storno because cinema doesnt exist")
	}
	for curreservation := range in.Seatstorno {
		if currentcinema.containsSeatMap(curreservation.row, curreservation.column) {
			currentcinema.seatmap[currentcinema.getSeat(curreservation.row, curreservation.column)] = false
		}
	}
	handler.cinemamap[in.Id] = currentcinema
	handler.mutex.Unlock()

	out.Answer = true

	fmt.Println("Cinema Service - Storno | Successfully deleted reservation")
	return nil
}

/*
CheckSeats checks if the requested seats are available
*/
func (handler *CinemaPool) CheckSeats(ctx context.Context, in *proto.CheckSeatsRequest, out *proto.CheckSeatsResponse) error {
	fmt.Println("Cinema Service - CheckSeats | Deleting reservation")

	handler.mutex.Lock()
	currentcinema, mapcontainscinema := handler.containsandreturncinema(in.Id)

	if !mapcontainscinema {
		handler.mutex.Unlock()
		out.Available = false
		return errors.New("Cinema Service - CheckSeats | Cannot execute storno because cinema doesnt exist")
	}

	check := true

	for curseat := range in.Seatcheck {
		if currentcinema.containsSeatMap(curseat.row, curseat.column) {
			value := currentcinema.seatmap[getSeat(curseat.row, curseat.column)]
			if value {
				check := false
			}
		}
	}

	handler.mutex.Unlock()

	out.Available = check

	if check {
		log.Println("Cinema Service - CheckSeats | Seats are available")
		return nil
	}
	log.Println("Cinema Service - CheckSeats | Seats are already reserved")
	return nil
}

/*
FreeSeats returns all seats that are available
*/
func (handler *CinemaPool) FreeSeats(ctx context.Context, in *proto.FreeSeatsRequest, out *proto.FreeSeatsResponse) error {
	fmt.Println("Cinema Service - FreeSeats | Return all free seats")

	handler.mutex.Lock()
	currentcinema, mapcontainscinema := handler.containsandreturncinema(in.Id)

	if !mapcontainscinema {
		handler.mutex.Unlock()
		out.Freeseats = nil
		return errors.New("Cinema Service - FreeSeats | Cannot execute storno because cinema doesnt exist")
	}

	retseatmap := map[*seats]bool{}
	for curseat := range currentcinema.seatmap {
		if currentcinema.containsSeatMap(curseat.row, curseat.column) {
			value := currentcinema.seatmap[getSeat(curseat.row, curseat.column)]
			if value {
				retseatmap[currentcinema.seatmap[getSeat(curseat.row, curseat.column)]] = true
			}
		}
	}
	handler.mutex.Unlock()

	avseats := proto.FreeSeatsResponse{
		Freeseats: retseatmap,
	}
	out.freeseats = avseats

	fmt.Println("Cinema Service - FreeSeats | Successfully executed FreeSeats")
	return nil
}

/*
Reset resets the CinemaPool
*/
func (handler *CinemaPool) Reset(context.Context, *proto.ResetRequest, *proto.ResetResponse) error {
	fmt.Println("Cinema Service - Clear | Reset Service with id", in.id)

	handler.mux.Lock()
	handler.cinemamap = make(map[int64]*cinema)
	handler.mutex.Unlock()

	log.Println("Cinema Service - Clear | Reset successfull")
	out.answer = true
	return nil
}
