package main

import (
	"fmt"
	"proto/crm"
	"./src"
	"github.com/micro/go-micro"
)

func main() {
	crmService := new(crm.CRMService)
	crmService.Init("localhost:8500", "localhost")

	service := micro.NewService(
		micro.Name("crmService"),
	)

	service.Init()
	crm_api.RegisterCRMServiceHandler(service.Server(), crmService)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
