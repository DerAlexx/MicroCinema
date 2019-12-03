package main

import (
	"fmt"

	protooo "github.com/ob-vss-ws19/blatt-4-pwn2own/show/proto"
	"github.com/ob-vss-ws19/blatt-4-pwn2own/show/show"

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

	protooo.RegisterCinemaHandler(service.Server(), new(show.ShowPool))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
