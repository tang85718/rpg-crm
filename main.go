package main

import (
	"github.com/micro/go-micro"
	"fmt"
	"proto"
	"./src"
)

func main() {
	service := micro.NewService(
		micro.Name("crm"),
	)

	crm := new(crm.CRMService)
	crm.Init("localhost:8500", "localhost")

	service.Init()
	crm_api.RegisterCRMServiceHandler(service.Server(), crm)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
