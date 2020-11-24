package services

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"time"
)
// OrderService ...
type OrderService struct {
}

func (o *OrderService) GetOrderMain(ctx context.Context, request *OrderRequest) (*OrderResponse, error) {
	fmt.Println(request)
	order := OrderMain{
		OrderId: 1,
		OrderNo: "202011111",
		OrderTime: &timestamp.Timestamp{
			Seconds: time.Now().Unix(),
		},
	}
	resp := &OrderResponse{
		ErrCode: 0,
		ErrMsg:  "SUCCESS",
		Data: []*OrderMain{
			&order,
		},
	}
	return resp, nil
}

func (o *OrderService) AddOrder(ctx context.Context, request *OrderRequest) (*OrderResponse, error) {
	// 引人字段校验
	err := request.OrderMain.Validate()
	if err != nil {
		resp := &OrderResponse{
			ErrCode: 1,
			ErrMsg:  err.Error(),
		}
		return resp,nil
	}
	order := OrderMain{
		//ProId:   2,
		OrderId: 2,
		OrderNo: "202011111",
		OrderTime: &timestamp.Timestamp{
			Seconds: time.Now().Unix(),
		},
	}
	resp := &OrderResponse{
		ErrCode: 0,
		ErrMsg:  "SUCCESS",
		Data: []*OrderMain{
			&order,
		},
	}
	return resp, nil
}
