package datamodels

// Order 订单结构体.
type Order struct {
	ID          int64 `json:"id" gorm:"id"`
	UserID      int64 `json:"user_id" gorm:"user_id" form:"user_id"`
	ProductID   int64 `json:"product_id"  gorm:"product_id" form:"product_id"`
	OrderStatus int64 `json:"order_status" gorm:"order_status" form:"order_status"` // 订单状态
}

const (
	// OrderWait 订单开始状态:0
	OrderWait = iota
	// OrderSuccess 订单等待:1
	OrderSuccess
	// OrderFailed 订单失败:2
	OrderFailed
)

// TableName gorm 重命名.
func (o Order) TableName() string {
	return "order"
}
