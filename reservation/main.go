package main

import (
	"fmt"

	"github.com/micro/go-micro"
	cinproto "github.com/ob-vss-ws19/blatt-4-pwn2own/cinemahall/proto"
	proto "github.com/ob-vss-ws19/blatt-4-pwn2own/reservation/proto"
	res "github.com/ob-vss-ws19/blatt-4-pwn2own/reservation/reservation"
	showproto "github.com/ob-vss-ws19/blatt-4-pwn2own/show/proto"
)

func main() {
	service := micro.NewService(micro.Name("reservation"))
	service.Init()
	newResService := res.CreateNewReservationHandlerInstance()
	newResService.AddDependencyRes(&res.ReservationsDependency{
		Show: func() showproto.ShowService {
			return showproto.NewShowService("show", service.Client())
		},
		Cinemahall: func() cinproto.CinemaService {
			return cinproto.NewCinemaService("cinemahall", service.Client())
		},
	})
	err1 := proto.RegisterReservationHandler(service.Server(), newResService)

	if err1 == nil {
		if err := service.Run(); err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err1)
	}
}
