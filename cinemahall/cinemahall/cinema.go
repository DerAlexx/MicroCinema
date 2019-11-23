package cinemahall

import (
	proto "github/ob-vss-ws19/blatt-4-pwn2own/cinemahall/proto"
)

type CinemaPool struct {
	cinemas map[int64]*proto.CinemaData
}

func NewCinemaPool *CinemaServiceHandler {
	cinemas := make(map[int64]*proto.CinemaData)

	return &CinemaServiceHandler{
		cinemas:      cinemas,
	}
}

func (handler *CinemaPool) Create (ctx context.Context, in *proto.CreateCinemaRequest, out *proto.CreateCinemaResponse) error {
	log.Printf("Cinema Service - Create | Creating cinema with name: %s, %d rows: and columns: %d \n", in.name, in.row, in.column)

	if len(in.name) == 0 || in.row == 0 || in.column == 0 {
		err := fmt.Errorf("Cannot create a cinema with an empty name, zero rows or zero columns")
		log.Errorf("Cinema Service - Create | %s\n", err.Error())
		return err
	} else {
		seatmap := make([]*proto.SeatMessage, in.row * in.column)
		for i := int64(0); i < in.row; i++ {
			for j := int64(0); j < in.column; j++ {
				seatmap = append(seatmap, &proto.SeatMessage {
					row: i + 1, 
					column: j + 1, 
					occupied: false,
				})
			}
		}

		createid := rand.Intn(10000)
	
		data := proto.CinemaMessage{
			name:      in.name,
			id:        createid,
			seats:     seatmap,
			rownumber:  in.row,
			columnnumber: in.column,
		}
	
		handler.mux.Lock()
		handler.cinemas[data.id] = &data
		out.Data = handler.cinemas[data.id]
		defer handler.mux.Unlock()
	
		log.Println("Cinema Service - Create | Successfully created cinema")
	
		return nil
	}
}

func (handler *CinemaPool) Delete(ctx context.Context, in *proto.DeleteCinemaRequest, out *proto.DeleteCinemaResponse) error {
	log.Printf("Cinema Service - Delete | Deleting cinema with id %d\n", in.id)
	handler.mux.Lock()
	_, mapcontainscinema := handler.cinemas[in.id]
	if !mapcontainscinema {
		out.Success = false
		handler.mux.Unlock()

		err := fmt.Errorf("Cannot delete cinema with id: %d", in.id)
		log.Printf("Cinema Service - Delete | %s\n", err.Error())
		return err
	}else {
		delete(handler.cinemas, in.id)
		handler.mux.Unlock()
		out.Success = true

		log.Println("Cinema Service - Delete | Successfully deleted cinema")
		return nil
	}
}

func (handler *CinemaPool) Reservation(ctx context.Context, in *proto.ReservationRequest, out *proto.ReservationResponse) error {
	log.Printf("Cinema Service - Reservation | Reservating seats %v for cinema id %d\n", in.column, in.id)

	handler.mux.Lock()
	currentcinema, mapcontainscinema := handler.cinemas[in.id]
	
	if !mapcontainscinema {
		defer handler.mux.Unlock()
		err := fmt.Error("Cannot execute reservation because cinema doesnt exist")
		log.Printf("Cinema Service - Reservation | %s\n", err.Error())
		return err
	} else {
		//fehlt
		defer handler.mux.Unlock()
		
		log.Println("Cinema Service - Reservation | Reservation successfull")
		return nil
	}
}

func (handler *CinemaPool) Storno(ctx context.Context, in *proto.StornoRequest, out *proto.StornoResponse) error {
	log.Println("Cinema Service - Storno | Deleting reservation")

	handler.mux.Lock()
	currentcinema, mapcontainscinema := handler.cinemas[in.Id]
	
	if !mapcontainscinema {
		defer handler.mux.Unlock()
		err := fmt.Errorf("Cannot execute storno because cinema doesnt exist")
		log.Printf("Cinema Service - Storno | %s\n", err.Error())
		return err
	} else {
		for _, seat := range in.seats {
			currentcinema.seats[((seat.row-1)*cinema.column)+seat.seat-1].reserved = false
		}
		handler.cinemas[in.id] = currentcinema
	
		out.seats = cinema.seats
	
		log.Println("Cinema Service - Storno | Successfully deleted reservation")
		return nil
	}
}

func (handler *CinemaPool) AvailableSeats(ctx context.Context, in *proto.AvailableSeatsRequest, out *proto.AvailableSeatsResponse) error {

}