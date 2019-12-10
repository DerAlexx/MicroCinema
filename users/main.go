package main

import (
	"fmt"

	micro "github.com/micro/go-micro"
	us "github.com/ob-vss-ws19/blatt-4-pwn2own/users/users"
	res "github.com/ob-vss-ws19/blatt-4-pwn2own/reservation/reservation"
)

/*
Main Function to start a new users service.
*/
func main() {
	service := micro.NewService(micro.Name("users"))
	service.Init()
	newUserService := us.CreateNewUserHandleInstance()
	newUserService.AddDependency(us.Dependencies{
		ResService: func() res.RegisterUsersHandler {
			return res.CreateNewReservationHandlerInstance("reservation", service.Client())
		}
	})
	proto.RegisterUsersHandler(service.Server())

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
