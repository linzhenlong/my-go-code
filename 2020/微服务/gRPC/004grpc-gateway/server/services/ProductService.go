package services

import "context"

type ProductService struct {

}

func(p *ProductService)GetProductStockStatus(ctx context.Context, request *ProductRequest) (*ProductResponse, error) {
	resp := &ProductResponse{
		StockStatus: 200,
	}
	return resp, nil
}
func (p *ProductService)GetProductName(ctx context.Context, request *ProductRequest) (*ProductResponse, error) {
	productName := "this is mate40 pro";
	if request.ProId == 100 {
		productName = "iphone 12 pro"
	}
	resp := &ProductResponse{
		ProductName: productName,
	}
	return resp, nil
}