package main

import (
	"fmt"

	"github.com/ob-vss-ws19/blatt-4-pwn2own/cinemahall/cinemahall"

	protooo "github.com/ob-vss-ws19/blatt-4-pwn2own/cinemahall/proto"

	micro "github.com/micro/go-micro"
)

const serviceName = "cinema-service"

func main() {

	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name(serviceName),
	)

	// Init will parse the command line flags.
	service.Init()

	protooo.RegisterCinemaHandler(service.Server(), new(cinemahall.CinemaServiceHandler))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}
