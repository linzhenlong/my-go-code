package services

import (
	"context"
	"strconv"
)

type ProductService struct {

}


func(p *ProductService)GetProductStockStatus(context.Context, *ProductRequest) (*ProductResponse, error) {
	return &ProductResponse{
		StockStatus: 200,
	},nil
}
func (p *ProductService)GetProductName(context.Context, *ProductRequest) (*ProductResponse, error) {
	return &ProductResponse{
		StockStatus: 200,
		ProductName: "THIS 好的手机",
	},nil
}
func (p *ProductService)GetProductProductList(ctx context.Context, size *QuerySize) (*ProductResponseList, error) {
	if size.Size <= 0 {
		size.Size = 10
	}
	list := make([]*ProductResponse,0)
	for  i :=0;i< int(size.Size);i++ {
		product := &ProductResponse{
			StockStatus: int32(i),
			ProductName: "this is 手机"+strconv.Itoa(i),
		}
		list = append(list, product)
	}
	resp := &ProductResponseList{
		List: list,
	}
	return resp,nil
}