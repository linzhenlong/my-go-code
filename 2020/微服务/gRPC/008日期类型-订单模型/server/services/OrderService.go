package services

import (
	"context"
	"fmt"
)

type OrderService struct {
}

func (o *OrderService) NewOrder(ctx context.Context, request *OrderMain) (*OrderResponse, error) {
	fmt.Println(request)
	return &OrderResponse{
		Message: "success",
		Status:  "ok",
	}, nil
}
