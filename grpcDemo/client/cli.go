package client

import (
	"context"
	"github.com/sirupsen/logrus"
	pb "goLibrary/grpcDemo/pbs/product"
	"google.golang.org/grpc"
)

var productCli = NewProductClient("localhost:9991")

func NewProductClient(address string) pb.ProductInfoClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	defer func() { _ = conn.Close() }()
	if err != nil {
		logrus.Fatalf("client run failed %s", err)
	}
	return pb.NewProductInfoClient(conn)
}

func GetProductInfo(ctx context.Context, req *pb.ProductReq, ) (*pb.ProductRes, error) {
	return productCli.GetProduct(ctx, req)
}
