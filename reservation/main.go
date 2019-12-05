package main

import (
	"fmt"

	"github.com/micro/go-micro"
	proto "github.com/ob-vss-ws19/blatt-4-pwn2own/reservation/proto"
	res "github.com/ob-vss-ws19/blatt-4-pwn2own/reservation/reservation"
)

func main() {
	service := micro.NewService(micro.Name("registration"))
	service.Init()
	proto.RegisterReservationHandler(service.Server(), new(res.ReservatServiceHandler))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
