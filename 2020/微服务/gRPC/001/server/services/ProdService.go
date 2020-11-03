package services

import (
	context "context"
)

// ProdService ...
type ProdService struct {
}

// GetProdStock ...
func (p *ProdService) GetProdStock(ctx context.Context, in *ProductRequest) (*ProdResponse, error) {
	return &ProdResponse{
		ProdStock: 20,
	}, nil
}
// GetProductName ...
func (p *ProdService)GetProductName(ctx context.Context, in *ProductRequest) (*ProdResponse, error) {
	resp := &ProdResponse{
		ProductName: "this is iphone 12pro",
	}
	return resp, nil
}
