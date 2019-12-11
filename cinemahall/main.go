package main

import (
	"fmt"

	cinemamicromain "github.com/micro/go-micro"
	"github.com/ob-vss-ws19/blatt-4-pwn2own/cinemahall/cinemahall"
	cinemaprotomain "github.com/ob-vss-ws19/blatt-4-pwn2own/cinemahall/proto"
)

const serviceName = "cinema-service"

func main() {
	// Create a new service. Optionally include some options here.
	service := cinemamicromain.NewService(
		cinemamicromain.Name(serviceName),
	)

	// Init will parse the command line flags.
	service.Init()

	err1 := cinemaprotomain.RegisterCinemaHandler(service.Server(), cinemahall.NewCinemaPool())
	if err1 == nil {
		// Run the server
		if err := service.Run(); err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}
