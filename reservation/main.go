package main

import (
	"fmt"

	"github.com/micro/go-micro"
	cinproto "github.com/ob-vss-ws19/blatt-4-pwn2own/cinemahall/proto"
	movieproto "github.com/ob-vss-ws19/blatt-4-pwn2own/movies/proto"
	proto "github.com/ob-vss-ws19/blatt-4-pwn2own/reservation/proto"
	res "github.com/ob-vss-ws19/blatt-4-pwn2own/reservation/reservation"
)

func main() {
	service := micro.NewService(micro.Name("reservation"))
	service.Init()
	newResService := res.CreateNewReservationHandlerInstance()
	newResService.AddDependencyRes(&res.ReservationsDependency{
		Movies: func() movieproto.MoviesService {
			return movieproto.NewMoviesService("movies", service.Client())
		},
		Cinemahall: func() cinproto.CinemaService {
			return cinproto.NewCinemaService("cinemahall", service.Client())
		},
	})
	proto.RegisterReservationHandler(service.Server(), newResService)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
