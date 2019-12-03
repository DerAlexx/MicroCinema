package main

import (
	"fmt"

	micro "github.com/micro/go-micro"
	proto "github.com/ob-vss-ws19/blatt-4-pwn2own/users/proto"
	"github.com/ob-vss-ws19/blatt-4-pwn2own/users/users"
)

func main() {
	service := micro.NewService(micro.Name("users"))
	service.Init()
	proto.RegisterUsersHandler(service.Server(), new(users.UserHandlerService))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
