package main

import (
	"fmt"

	showproto "github.com/ob-vss-ws19/blatt-4-pwn2own/show/proto"

	cinemamicromain "github.com/micro/go-micro"
	"github.com/ob-vss-ws19/blatt-4-pwn2own/cinemahall/cinemahall"
	cinemaprotomain "github.com/ob-vss-ws19/blatt-4-pwn2own/cinemahall/proto"
)

const serviceName = "cinemahall"

func main() {
	service := cinemamicromain.NewService(
		cinemamicromain.Name(serviceName),
	)
	service.Init()
	newService := cinemahall.NewCinemaPool()
	newService.AddDependency(&cinemahall.CinemaDependency{
		ShowService: func() showproto.ShowService {
			return showproto.NewShowService("show", service.Client())
		},
	})
	err1 := cinemaprotomain.RegisterCinemaHandler(service.Server(), newService)
	if err1 == nil {
		if err := service.Run(); err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err1)
	}
}
