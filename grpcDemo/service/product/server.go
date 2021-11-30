package product

import (
	"context"
	pb "goLibrary/grpcDemo/pbs/product"
)

type Service struct {
}

func (s *Service) GetProduct(ctx context.Context, req *pb.ProductReq) (*pb.ProductRes, error) {
	panic("implement me")
}

func (s *Service) AddProduct(ctx context.Context, p *pb.Product) (*pb.ProductRes, error) {
	panic("implement me")
}
