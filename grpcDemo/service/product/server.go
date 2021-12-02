package product

import (
	"context"
	"github.com/gofrs/uuid"
	pb "goLibrary/grpcDemo/pbs/product"
)

type Service struct {
}

// 单项流
func (s *Service) SearchProduct(ctx context.Context, req *pb.ProductSearchReq) (*pb.ProductRes, error) {
	panic("implement me")
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
