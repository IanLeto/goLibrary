package server

import (
	"context"
	"goLibrary/grpcDemo/proto/product"
)

type ProductServer struct {
}

func (p *ProductServer) GetProduct(ctx context.Context, req *product.ProductReq) (*product.ProductRes, error) {
	return &product.ProductRes{Name: "res service"}, nil
}
