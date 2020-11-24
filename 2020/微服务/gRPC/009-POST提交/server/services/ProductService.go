package services

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"strconv"
	"time"
)

type ProductService struct {
}

/**
GetProductList(context.Context, *ProductRequest) (*ProductResponse, error)
	GetProductInfo(context.Context, *ProductRequest) (*ProductResponse, error)
	AddProduct(context.Context, *ProductRequest) (*ProductResponse, error)
*/

func (p *ProductService) GetProductList(ctx context.Context, request *ProductRequest) (*ProductResponse, error) {
	fmt.Println(request)
	data := make([]*ProductInfo, 0)
	for i := 0; i < 10; i++ {
		productInfo := &ProductInfo{
			ProArea:        ProductAreas_BEIJING,
			ProId:          int32(i),
			ProName:        "商品名称:" + strconv.Itoa(i),
			ProStockStatus: int32(i * 20),
			ProTag:         []string{"手机", "手机2"},
			ProPrice:       1999.88,
			ProTime: &timestamp.Timestamp{
				Seconds: time.Now().Unix(),
			},
		}
		data = append(data, productInfo)
	}
	resp := &ProductResponse{
		Data:    data,
		ErrCode: 0,
		ErrMsg:  "SUCCESS",
	}
	return resp, nil
}

func (p *ProductService) GetProductInfo(ctx context.Context, request *ProductRequest) (*ProductResponse, error) {
	fmt.Println(request)
	info := &ProductInfo{
		ProArea:        ProductAreas_BEIJING,
		ProId:          int32(10),
		ProName:        "商品名称:" + strconv.Itoa(10),
		ProStockStatus: int32(200),
		ProTag:         []string{"手机", "手机2"},
		ProPrice:       1999.88,
		ProTime: &timestamp.Timestamp{
			Seconds: time.Now().Unix(),
		},
	}
	var data []*ProductInfo
	data = append(data, info)

	resp := &ProductResponse{
		ErrMsg:  "SUCCESS",
		ErrCode: 0,
		Data:    data,
	}
	return resp, nil
}


func(p *ProductService)AddProduct(ctx context.Context, request *ProductRequest) (*ProductResponse, error) {
	fmt.Println(request)
	info := &ProductInfo{
		ProArea:        ProductAreas_BEIJING,
		ProId:          int32(10),
		ProName:        "商品名称:" + strconv.Itoa(10)+"这是添加的",
		ProStockStatus: int32(200),
		ProTag:         []string{"手机", "手机2"},
		ProPrice:       1999.88,
		ProTime: &timestamp.Timestamp{
			Seconds: time.Now().Unix(),
		},
	}
	var data []*ProductInfo
	data = append(data, info)
	resp := &ProductResponse{
		ErrMsg:  "SUCCESS",
		ErrCode: 0,
		Data:    data,
	}
	return resp, nil
}