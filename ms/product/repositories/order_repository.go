package repositories

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/linzhenlong/my-go-code/ms/product/common"
	"github.com/linzhenlong/my-go-code/ms/product/datamodels"
)

// IOrderRepository 订单接口.
type IOrderRepository interface {
	Conn() error
	Insert(*datamodels.Order) (int64, error)
	Delete(int64) bool
	Update(*datamodels.Order) error
	SelectByKey(int64) (*datamodels.Order, error)
	SelectAll(map[string]interface{}) ([]*datamodels.Order, error)
	GetTotal(map[string]interface{}) int64
}

// OrderManager 订单管理
type OrderManager struct {
	table    string
	gormCoon *gorm.DB
}

// NewOrderManager 实例化.
func NewOrderManager(table string, conn *gorm.DB) IOrderRepository {
	return &OrderManager{
		table:    table,
		gormCoon: conn,
	}
}

// Conn 数据库连接.
func (o *OrderManager) Conn() error {
	if o.gormCoon == nil {
		myGorm, err := common.NewGorm()
		if err != nil {
			return err
		}
		o.gormCoon = myGorm
	}
	if o.table == "" {
		o.table = "order"
	}
	return nil
}

// Insert 插入订单.
func (o *OrderManager) Insert(order *datamodels.Order) (orderID int64, err error) {
	if err = o.Conn(); err != nil {
		return
	}
	err = o.gormCoon.Debug().Create(&order).Error
	if err != nil {
		return orderID, err
	}
	return order.ID, nil
}

// Delete 删除订单.
func (o *OrderManager) Delete(orderID int64) bool {

	if err := o.Conn(); err != nil {
		return false
	}
	order := datamodels.Order{
		ID: orderID,
	}
	err := o.gormCoon.Debug().Delete(&order).Error
	if err != nil {
		return false
	}
	return true
}

// Update 更新.
func (o *OrderManager) Update(order *datamodels.Order) (err error) {
	if err = o.Conn(); err != nil {
		return err
	}

	err = o.gormCoon.Debug().Model(order).Updates(order).Error
	return
}

// SelectByKey 查询.
func (o *OrderManager) SelectByKey(orderID int64) (order *datamodels.Order, err error) {

	if err = o.Conn(); err != nil {
		return nil, err
	}
	if orderID == 0 {
		return nil, fmt.Errorf("%s", "orderId empty")
	}
	order = &datamodels.Order{}
	o.gormCoon.Debug().First(&order, orderID)
	return
}

// SelectAll 获取全部.
func (o *OrderManager) SelectAll(params map[string]interface{}) (orders []*datamodels.Order, err error) {
	if err = o.Conn(); err != nil {
		return nil, err
	}

	orderStatus, ok := params["orderStatus"].(int64)
	db := o.gormCoon.Debug()
	if ok {
		db = db.Where("order_status = ? ", orderStatus)
	}

	// 筛选商品.
	if productID, ok := params["product_id"].(int64); ok {
		db = db.Where("product_id = ? ", productID)
	}
	// 筛选用户.
	if userID, ok := params["user_id"].(int64); ok {
		db = db.Where("user_id = ? ", userID)
	}

	limit, ok := params["limit"].(int)
	if ok {
		db = db.Limit(limit)
	}
	offset, ok := params["offset"].(int)
	if ok {
		db = db.Offset(offset)
	}

	res := []*datamodels.Order{}
	err = db.Find(&res).Error
	return res, err
}

// GetTotal 获取总数
func (o *OrderManager) GetTotal(params map[string]interface{}) (total int64) {
	if err := o.Conn(); err != nil {
		return 0
	}
	orderStatus, ok := params["orderStatus"].(int64)
	db := o.gormCoon.Debug().Model(&datamodels.Order{})
	if ok {
		db = db.Where("order_status = ? ", orderStatus)
	}

	// 筛选商品.
	if productID, ok := params["product_id"].(int64); ok {
		db = db.Where("product_id = ? ", productID)
	}
	// 筛选用户.
	if userID, ok := params["user_id"].(int64); ok {
		db = db.Where("user_id = ? ", userID)
	}
	db.Count(&total)
	return
}
