package main

import (
	"fmt"

	"github.com/micro/go-micro"
	//proto "github.com/ob-vss-ws19/blatt-4-pwn2own/reservation/proto"
	res "github.com/ob-vss-ws19/blatt-4-pwn2own/reservation/reservation"
)

func main() {
	service := micro.NewService(micro.Name("users"))
	service.Init()
	proto.RegisterUsersHandler(service.Server(), new(res.ReservatServiceHandler))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
