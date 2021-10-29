package main

import (
	"context"
	"fmt"
	"goLibrary/grpcDemo/proto/product"
	"goLibrary/grpcDemo/server"
	"google.golang.org/grpc"
	"net"
)

func main() {
	// 启动 rpc service
	go func() {
		rpcServer := grpc.NewServer()
		// 注册 product server gen 自动生成的
		product.RegisterProductServerServer(rpcServer, new(server.ProductServer))
		//
		listener, err := net.Listen("tcp", ":9001")
		if err != nil {
			panic(err)
		}
		_ = rpcServer.Serve(listener)
	}()
	conn, err := grpc.Dial(":9001", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	cli := product.NewProductServerClient(conn)
	res, err := cli.GetProduct(context.Background(), &product.ProductReq{Id: 1})
	if err != nil {
		panic(err)
	}
	fmt.Println("hello", res.Name)

}
