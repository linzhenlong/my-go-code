package services

import (
	"github.com/linzhenlong/my-go-code/ms/product/datamodels"
	"github.com/linzhenlong/my-go-code/ms/product/repositories"
)

// IOrderService 订单服务接口.
type IOrderService interface {
	GetOrderByID(int64) (*datamodels.Order, error)
	InsertOrder(*datamodels.Order) (int64, error)
	UpdateOrder(*datamodels.Order) error
	DeleteOrder(int64) bool
	GetOrderList(map[string]interface{}) ([]*datamodels.Order, error)
	GetOrderTotal(map[string]interface{}) int64
}

// OrderService 订单服务.
type OrderService struct {
	orderRepository repositories.IOrderRepository
}

// NewOrderService 实例OrderService.
func NewOrderService(repository repositories.IOrderRepository) IOrderService {
	return &OrderService{
		orderRepository: repository,
	}
}

// GetOrderByID 获取订单详情
func (o *OrderService) GetOrderByID(orderID int64) (order *datamodels.Order, err error) {
	return o.orderRepository.SelectByKey(orderID)
}

// InsertOrder 插入订单
func (o *OrderService) InsertOrder(order *datamodels.Order) (orderID int64, err error) {
	return o.orderRepository.Insert(order)
}

// UpdateOrder 更新订单.
func (o *OrderService) UpdateOrder(order *datamodels.Order) (err error) {
	return o.orderRepository.Update(order)
}

// DeleteOrder 删除订单.
func (o *OrderService) DeleteOrder(orderID int64) bool {
	return o.orderRepository.Delete(orderID)
}

// GetOrderList 获取订单列表.
func (o *OrderService) GetOrderList(params map[string]interface{}) (orders []*datamodels.Order, err error) {
	return o.orderRepository.SelectAll(params)
}

// GetOrderTotal 获取订单总数.
func (o *OrderService) GetOrderTotal(params map[string]interface{}) int64 {
	return o.orderRepository.GetTotal(params)
}
