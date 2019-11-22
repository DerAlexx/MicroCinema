package main

import (
	micro "github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(micro.Name("users"))
	service.Init()

}
