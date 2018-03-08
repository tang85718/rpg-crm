package main

import (
	"fmt"
	"proto/crm"
	"./src"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	crmService := new(crm.CRMService)
	crmService.Init("localhost:8500", "localhost")

	crm_api.RegisterCRMServiceServer(s, crmService)

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		fmt.Println(err)
	}

	//service := micro.NewService(
	//	micro.Name("crmService"),
	//)
	//service.Init()
	//crm_api.RegisterCRMServiceHandler(service.Server(), crmService)
	//
	//if err := service.Run(); err != nil {
	//	fmt.Println(err)
	//}
}
