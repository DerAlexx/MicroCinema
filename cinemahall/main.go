package main

import (
	"fmt"

	micro "github.com/micro/go-micro"
	"github.com/ob-vss-ws19/blatt-4-pwn2own/cinemahall/proto"
)

const serviceName = "cinema-service"

func main() {

	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name(serviceName),
	)

	// Init will parse the command line flags.
	service.Init()

	//RegisterHandler => fehlt
	proto.RegisterCinemaHandler(service.Server(), NewCinemaPool())

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}
