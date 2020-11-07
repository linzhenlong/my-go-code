package services

import (
	"context"
	"strconv"
)

// ProductService ...
type ProductService struct {
}

func (p *ProductService)GetProductInfo(ctx context.Context, request *ProductRequest) (*ProductResponse, error) {

	info := &ProductInfo{}
	if request.ProArea == ProductAreas_BEI_JING {
		info = &ProductInfo{
			ProName:  "this is iphone 12 pro for 来自北京仓库",
			ProId:    1,
			ProTag:   []string{"神价格", "绝对值", "bug价"},
			ProPrice: 1999.80,
			ProArea: ProductAreas_BEI_JING,
		}
	} else {
		info = &ProductInfo{
			ProName:  "this is iphone 12 pro for 来其他仓库",
			ProId:    1,
			ProTag:   []string{"神价格", "绝对值", "bug价"},
			ProPrice: 1999.80,
		}
	}
	data := make([]*ProductInfo,0)
	data = append(data, info)
	resp := &ProductResponse{
		ErrCode: 0,
		ErrMsg: "success",
		Data: data,
	}
	return resp, nil
}

func(p  *ProductService)GetProductList(ctx context.Context, request *ProductRequest) (*ProductResponse, error)  {
	if request.Size <= 0 && request.Size >=20 {
		request.Size = 10
	}
	data := make([]*ProductInfo,0)
	var i int32
	for i=0;i<request.Size;i++ {
		productInfo := &ProductInfo{
			ProId: i,
			ProName: "商品名称"+strconv.Itoa(int(i)),
			ProTag: []string{"神价格","绝对值","bug价"},
			ProPrice: float32(i)*66.66,
		}
		data = append(data, productInfo)
	}
	resp := &ProductResponse{
		ErrCode: 0,
		ErrMsg: "success",
		Data: data,
	}
	return resp, nil
}
