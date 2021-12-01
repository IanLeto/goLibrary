package main

import (
	"goLibrary/grpcDemo/pbs/product"
	productServer "goLibrary/grpcDemo/service/product"
	"goLibrary/utils"
	"google.golang.org/grpc"
	"net"
)

const port = "6666"

func main() {
	listen, err := net.Listen("tcp", "localhost:9991")
	if err != nil {
		panic(err)
	}
	service := grpc.NewServer()
	product.RegisterProductInfoServer(service, &productServer.Service{})
	utils.NoErr(service.Serve(listen))
}
