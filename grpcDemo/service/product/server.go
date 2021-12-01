package product

import (
	"context"
	"github.com/gofrs/uuid"
	pb "goLibrary/grpcDemo/pbs/product"
)

type Service struct {
}

func (s *Service) GetProduct(ctx context.Context, req *pb.ProductReq) (*pb.ProductRes, error) {
	out, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	return &pb.ProductRes{Name: "demo1", Id: out.String()}, nil
}

func (s *Service) AddProduct(ctx context.Context, p *pb.Product) (*pb.ProductRes, error) {
	panic("implement me")
}
