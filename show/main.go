package main

import (
	"fmt"

	micro "github.com/micro/go-micro"
	showproto "github.com/ob-vss-ws19/blatt-4-pwn2own/show/proto"
	"github.com/ob-vss-ws19/blatt-4-pwn2own/show/show"
)

const serviceName = "show-service"

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name(serviceName),
	)

	// Init will parse the command line flags.
	service.Init()

	err1 := showproto.RegisterShowHandler(service.Server(), show.NewShowPool())

	// Run the server
	if err1 == nil {
		if err := service.Run(); err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err1)
	}

}
