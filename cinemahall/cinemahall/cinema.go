package cinemahall

import (
	"context"
	proto "github/ob-vss-ws19/blatt-4-pwn2own/cinemahall/proto"
	"math/rand"
	"time"
)

const (
	maxcinemaid = 432342
)

/*
CinemaPool comment
*/
type CinemaPool struct {
	cinemamap map[int64]*cinema
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
NewCinemaPool comment
*/
func NewCinemaPool() *CinemaServiceHandler {
	newcinema := make(map[int64]*cinema)

	return &CinemaServiceHandler{
		cinemamap: newcinemas,
	}
}

func getRandomCinemaID() int64 {
	rand.Seed(time.Now().UnixNano())
	for {
		potentialID := rand.Intn(maxcinemaid)
		if !containscinema(potentialID) {
			return potentialID
		}
	}
}

func (handler *CinemaPool) containscinema(id int64) bool {
	_, containsinmap := handler.cinemamap[id]
	return containsinmap
}

func (handler *CinemaPool) containsandreturncinema(id int64) (*cinema, bool) {
	currentcinema, mapcontainscinema := handler.cinemamap[in.id]
	return currentcinema, mapcontainscinema
}

/*
Create comment
*/
func (handler *CinemaPool) Create(ctx context.Context, in *proto.CreateCinemaRequest, out *proto.CreateCinemaResponse) error {
	log.Printf("Cinema Service - Create | Creating cinema with name: %s, %d rows: and columns: %d \n", in.name, in.row, in.column)

	if len(in.name) == 0 || in.row == 0 || in.column == 0 {
		err := fmt.Errorf("Cannot create a cinema with an empty name, zero rows or zero columns")
		log.Errorf("Cinema Service - Create | %s\n", err.Error())
		return err
	}

	newseatmap := [*seats]bool{}
	for i := int64(0); i < in.row; i++ {
		for j := int64(0); j < in.column; j++ {
			newseatmap[&seats{row: i, column: j}] = false
		}
	}

	createid := getRandomCinemaID()
	handler.mux.Lock()
	handler.cinemamap[createid] = newseatmap
	defer handler.mux.Unlock()
	out.name = in.name
	out.id = createid

	log.Println("Cinema Service - Create | Successfully created cinema")

	return nil

}

/*
Delete comment
*/
func (handler *CinemaPool) Delete(ctx context.Context, in *proto.DeleteCinemaRequest, out *proto.DeleteCinemaResponse) error {
	log.Printf("Cinema Service - Delete | Deleting cinema with id %d\n", in.id)
	handler.mux.Lock()
	mapcontainscinema := handler.containscinema(in.id)
	if !mapcontainscinema {
		out.answer = false
		handler.mux.Unlock()

		err := fmt.Errorf("Cannot delete cinema with id: %d", in.id)
		log.Printf("Cinema Service - Delete | %s\n", err.Error())
		return err
	}
	delete(handler.cinemamap, in.id)
	handler.mux.Unlock()
	out.answer = true

	log.Println("Cinema Service - Delete | Successfully deleted cinema")
	return nil
}

/*
Reservation comment
*/
func (handler *CinemaPool) Reservation(ctx context.Context, in *proto.ReservationRequest, out *proto.ReservationResponse) error {
	log.Printf("Cinema Service - Reservation | Reservating seats %v for cinema id %d\n", in.column, in.id)

	handler.mux.Lock()
	currentcinema, mapcontainscinema := containsandreturncinema(in.id)

	if !mapcontainscinema {
		defer handler.mux.Unlock()
		err := fmt.Error("Cannot execute reservation because cinema doesnt exist")
		log.Printf("Cinema Service - Reservation | %s\n", err.Error())
		out.answer = false
		return err
	}
	for curreservation := range in.seatreservation {
		if currentcinema.containsSeatMap(curreservation.row, curreservation.column) {
			currentcinema.seatmap[getSeat(curreservation.row, curreservation.column)] == true
		}
	}
	handler.cinemamap[in.id] = currentcinema
	defer handler.mux.Unlock()

	log.Println("Cinema Service - Reservation | Reservation successfull")
	out.answer = true
	return nil

}

func (currentcinema *cinema) containsSeatMap(row int64, column int64) bool {
	for seat, occ := range currentcinema.seatmap {
		if seat.row == row && seat.column == column {
			return occ
		}
	}
}

func (currentcinema *cinema) getSeat(row int64, column int64) *seats {
	for seat, occ := range currentcinema.seatmap {
		if seat.row == row && seat.column == column {
			return seat
		}
	}
}

/*
Storno comment
*/
func (handler *CinemaPool) Storno(ctx context.Context, in *proto.StornoRequest, out *proto.StornoResponse) error {
	log.Println("Cinema Service - Storno | Deleting reservation")

	handler.mux.Lock()
	currentcinema, mapcontainscinema := containsandreturncinema(in.id)

	if !mapcontainscinema {
		defer handler.mux.Unlock()
		err := fmt.Errorf("Cannot execute storno because cinema doesnt exist")
		log.Printf("Cinema Service - Storno | %s\n", err.Error())
		out.answer = false
		return err
	}
	for curreservation := range in.seatstorno {
		if !currentcinema.containsSeatMap(curreservation.row, curreservation.column) {
			currentcinema.seatmap[getSeat(curreservation.row, curreservation.column)] == false
		}
	}
	handler.cinemamap[in.id] = currentcinema

	out.answer = true

	log.Println("Cinema Service - Storno | Successfully deleted reservation")
	return nil

}

/*
AvailableSeats comment
*/
func (handler *CinemaPool) AvailableSeats(ctx context.Context, in *proto.AvailableSeatsRequest, out *proto.AvailableSeatsResponse) error {

}
