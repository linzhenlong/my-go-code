package services

import "context"

// ProductServices ...
type ProductService struct {

}
// GetProductStockStatus ...
func (p *ProductService)GetProductStockStatus(context.Context, *ProductRequest) (*ProductResponse, error) {
	resp := &ProductResponse{
		StockStatus: 10,
	}
	return resp, nil
}

func (p *ProductService)GetProductName(context.Context, *ProductRequest) (*ProductResponse, error) {
	resp := &ProductResponse{
		ProductName: "这是华为mate 40 pro rs",
	}
	return resp, nil
	
}
