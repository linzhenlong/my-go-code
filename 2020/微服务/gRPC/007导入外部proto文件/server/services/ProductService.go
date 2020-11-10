package services

import (
	"context"
	"strconv"
)

type ProductService struct {
}

func (p *ProductService) GetProductInfo(ctx context.Context, request *ProductRequest) (*ProductModel, error) {
	resp := &ProductModel{
		ProName:  "this is iphone 12 pro",
		ProPrice: 198.00,
		ProArea:  ProductAreas_BEI_JING,
		ProId:    1,
	}
	return resp, nil
}

func (p *ProductService) GetProductList(ctx context.Context, request *ProductRequest) (*ProductResponse, error) {
	if request.Size <= 0 || request.Size > 20 {
		request.Size = 20
	}
	var i int32
	rows := make([]*ProductModel, 0)
	for i = 0; i < request.Size; i++ {
		model := ProductModel{
			ProId:    i,
			ProPrice: float32(i) * 38.88,
			ProName:  "商品名称:" + strconv.Itoa(int(i)),
		}
		rows = append(rows, &model)
	}
	resp := &ProductResponse{
		ErrMsg:  "success",
		ErrCode: 0,
		Data:    rows,
	}
	return resp, nil
}
