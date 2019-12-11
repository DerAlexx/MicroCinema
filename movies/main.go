package main

import (
	"fmt"

	"github.com/micro/go-micro"
	"github.com/ob-vss-ws19/blatt-4-pwn2own/movies/movies"
	proto "github.com/ob-vss-ws19/blatt-4-pwn2own/movies/proto"
)

func main() {
	service := micro.NewService(micro.Name("movies"))
	service.Init()
	err1 := proto.RegisterMoviesHandler(service.Server(), movies.CreateNewMoviesHandlerInstance())
	if err1 == nil {
		if err := service.Run(); err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err1)
	}

}
