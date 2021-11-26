package service

import (
	"context"
	"goLibrary/grpcDemo/proto/product"
)

// ctrl + i 继承 rpc service 方法
type ProductServer struct {
}

func (p *ProductServer) GetProduct(ctx context.Context, req *product.ProductReq) (*product.ProductRes, error) {
	return &product.ProductRes{Name: "res service"}, nil
}
