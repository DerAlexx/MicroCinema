package main

import (
	"fmt"

	"github.com/ob-vss-ws19/blatt-4-pwn2own/cinemahall/cinemahall"

	cinemaprotomain "github.com/ob-vss-ws19/blatt-4-pwn2own/cinemahall/proto"

	cinemamicromain "github.com/micro/go-micro"
)

const serviceName = "cinema-service"

func main() {

	// Create a new service. Optionally include some options here.
	service := cinemamicromain.NewService(
		cinemamicromain.Name(serviceName),
	)

	// Init will parse the command line flags.
	service.Init()

	cinemaprotomain.RegisterCinemaHandler(service.Server(), new(cinemahall.CinemaPool))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
