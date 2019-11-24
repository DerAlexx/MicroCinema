package main

import (
	"fmt"

	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(micro.Name("movies"))
	service.Init()
	proto.RegisterUsersHandler(service.Server(), new()) //TODO missing handler

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}
