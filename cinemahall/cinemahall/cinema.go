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
CinemaPool contains all cinemas.
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
NewCinemaPool creates a new CinemaPool
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
Create creates a new cinema. It will be saved in the CinemaPool. The id to access the new cinema will be randomly generated.
If the creation was successfull the id will be returned
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
	handler.mutex.Lock()
	handler.cinemamap[createid] = newseatmap
	handler.mutex.Unlock()
	out.name = in.name
	out.id = createid

	log.Println("Cinema Service - Create | Successfully created cinema")

	return nil

}

/*
Delete will delete a cinema(id) from the CinemaPool.
*/
func (handler *CinemaPool) Delete(ctx context.Context, in *proto.DeleteCinemaRequest, out *proto.DeleteCinemaResponse) error {
	log.Printf("Cinema Service - Delete | Deleting cinema with id %d\n", in.id)
	handler.mutex.Lock()
	mapcontainscinema := handler.containscinema(in.id)
	if !mapcontainscinema {
		handler.mutex.Unlock()
		out.answer = false
		err := fmt.Errorf("Cannot delete cinema with id: %d", in.id)
		log.Printf("Cinema Service - Delete | %s\n", err.Error())
		return err
	}
	delete(handler.cinemamap, in.id)
	handler.mutex.Unlock()
	out.answer = true

	log.Println("Cinema Service - Delete | Successfully deleted cinema")
	return nil
}

/*
Reservation will change the value in the seatmap of a cinema to true (for the given row and column) if the seat is still available
*/
func (handler *CinemaPool) Reservation(ctx context.Context, in *proto.ReservationRequest, out *proto.ReservationResponse) error {
	log.Printf("Cinema Service - Reservation | Reservating seats %v for cinema id %d\n", in.column, in.id)

	handler.mutex.Lock()
	currentcinema, mapcontainscinema := containsandreturncinema(in.id)

	if !mapcontainscinema {
		handler.mutex.Unlock()
		err := fmt.Error("Cannot execute reservation because cinema doesnt exist")
		log.Printf("Cinema Service - Reservation | %s\n", err.Error())
		out.answer = false
		return err
	}
	for curreservation := range in.seatreservation {
		if currentcinema.containsSeatMap(curreservation.row, curreservation.column) {
			currentcinema.seatmap[getSeat(curreservation.row, curreservation.column)] = true
		}
	}
	handler.cinemamap[in.id] = currentcinema
	handler.mutex.Unlock()

	log.Println("Cinema Service - Reservation | Reservation successfull")
	out.answer = true
	return nil
}

/*
Storno will undo a reservation
*/
func (handler *CinemaPool) Storno(ctx context.Context, in *proto.StornoRequest, out *proto.StornoResponse) error {
	log.Println("Cinema Service - Storno | Deleting reservation")

	handler.mutex.Lock()
	currentcinema, mapcontainscinema := containsandreturncinema(in.id)

	if !mapcontainscinema {
		handler.mutex.Unlock()
		err := fmt.Errorf("Cannot execute storno because cinema doesnt exist")
		log.Printf("Cinema Service - Storno | %s\n", err.Error())
		out.answer = false
		return err
	}
	for curreservation := range in.seatstorno {
		if currentcinema.containsSeatMap(curreservation.row, curreservation.column) {
			currentcinema.seatmap[getSeat(curreservation.row, curreservation.column)] = false
		}
	}
	handler.cinemamap[in.id] = currentcinema
	handler.mutex.Unlock()

	out.answer = true

	log.Println("Cinema Service - Storno | Successfully deleted reservation")
	return nil
}

/*
CheckSeats checks if the requested seats are available
*/
func (handler *CinemaPool) CheckSeats(ctx context.Context, in *proto.CheckSeatsRequest, out *proto.CheckSeatsResponse) error {
	log.Println("Cinema Service - CheckSeats | Deleting reservation")

	handler.mutex.Lock()
	currentcinema, mapcontainscinema := containsandreturncinema(in.id)

	if !mapcontainscinema {
		handler.mutex.Unlock()
		err := fmt.Errorf("Cannot execute checkseats because cinema doesnt exist")
		log.Printf("Cinema Service - CheckSeats | %s\n", err.Error())
		out.answer = false
		return err
	}

	check := true

	for curseat := range in.seatcheck {
		if currentcinema.containsSeatMap(curseat.row, curseat.column) {
			value := currentcinema.seatmap[getSeat(curseat.row, curseat.column)]
			if value {
				check := false
			}
		}
	}
	handler.mutex.Unlock()

	out.answer = check

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
	log.Println("Cinema Service - FreeSeats | Return all free seats")

	handler.mutex.Lock()
	currentcinema, mapcontainscinema := containsandreturncinema(in.id)

	if !mapcontainscinema {
		handler.mutex.Unlock()
		err := fmt.Errorf("Cannot execute freeseats because cinema doesnt exist")
		log.Printf("Cinema Service - FreeSeats | %s\n", err.Error())
		out.answer = false
		return err
	}

	retseatmap := [*seats]bool{}
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
		freeseats: retseatmap,
	}
	out.freeseats = avseats

	log.Println("Cinema Service - FreeSeats | Successfully executed FreeSeats")
	return nil
}

/*
Reset resets the CinemaPool
*/
func (handler *CinemaServiceHandler) Reset(context.Context, *proto.ResetRequest, *proto.ResetResponse) error {
	log.Printf("Cinema Service - Clear | Reset Service with id\n", in.id)

	handler.mux.Lock()
	handler.cinemamap = make(map[int64]*cinema)
	handler.mutex.Unlock()

	log.Println("Cinema Service - Clear | Reset successfull")
	out.answer = true
	return nil
}
