package services

import (
	"context"
	"fmt"
)

type OrderService struct {
}

func (o *OrderService) NewOrder(ctx context.Context, request *OrderRequest) (*OrderResponse, error) {
	fmt.Println(request.OrderMain)
	return &OrderResponse{
		Message: "success",
		Status:  "ok",
	}, nil
}
